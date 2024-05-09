package sql

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	"server/app/pkg/errors"
)

var _ SQL = &DB{DB: nil}

type SQL interface {
	Unsafe() *DB
	Begin(context.Context) (*Tx, error)
	Ping() error
	In(query string, args ...any) (string, []any, error)
	Get(ctx context.Context, dest any, query string, args ...any) error
	Select(ctx context.Context, dest any, query string, args ...any) error
	Query(ctx context.Context, query string, args ...any) (*Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) (*Row, error)
	Exec(ctx context.Context, query string, args ...any) error
	ExecWithLastInsertID(ctx context.Context, query string, args ...any) (uint32, error)
	ExecWithRowsAffected(ctx context.Context, query string, args ...any) (uint32, error)
	Prepare(ctx context.Context, query string) (*Stmt, error)
	closer
}

type DB struct {
	DB *sqlx.DB
}

func Open(driverName string, url string) (*DB, error) {
	db, err := sqlx.Open(driverName, url)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err, secondPathDepthOption...)
	}
	return &DB{db}, nil
}

func (s *DB) Close() error {
	if err := s.DB.Close(); err != nil {
		return errors.InternalServer.Wrap(err, secondPathDepthOption...)
	}
	return nil
}

func (s *DB) Begin(ctx context.Context) (*Tx, error) {
	tx, err := s.DB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err, secondPathDepthOption...)
	}
	return &Tx{tx}, nil
}

func (s *DB) Ping() error {
	if err := s.DB.Ping(); err != nil {
		return errors.InternalServer.Wrap(err, secondPathDepthOption...)
	}
	return nil
}

func (s *DB) In(query string, args ...any) (_ string, _ []any, err error) {
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return "", nil, errors.InternalServer.Wrap(err, secondPathDepthOption...)
	}
	return query, args, nil
}

func (s *DB) Unsafe() *DB {
	return &DB{s.DB.Unsafe()}
}

func (s *DB) Select(ctx context.Context, dest any, query string, args ...any) (err error) {

	query, err = replacePlaceholders(query)
	if err != nil {
		return err
	}

	if tx := extractTx(ctx); tx != nil {
		err = tx.Tx.SelectContext(ctx, dest, query, args...)
	} else {
		err = s.DB.SelectContext(ctx, dest, query, args...)
	}

	if err != nil {
		return errors.InternalServer.Wrap(err, secondPathDepthOption...)
	}

	return nil
}

func (s *DB) Get(ctx context.Context, dest any, query string, args ...any) (err error) {

	query, err = replacePlaceholders(query)
	if err != nil {
		return err
	}

	if tx := extractTx(ctx); tx != nil {
		err = tx.Tx.GetContext(ctx, dest, query, args...)
	} else {
		err = s.DB.GetContext(ctx, dest, query, args...)
	}

	if err != nil {
		return errors.InternalServer.Wrap(err, secondPathDepthOption...)
	}

	return nil
}

func (s *DB) Query(ctx context.Context, query string, args ...any) (_ *Rows, err error) {

	query, err = replacePlaceholders(query)
	if err != nil {
		return nil, err
	}

	rows := &Rows{Rows: nil}
	if tx := extractTx(ctx); tx != nil {
		rows.Rows, err = tx.Tx.QueryxContext(ctx, query, args...)
	} else {
		rows.Rows, err = s.DB.QueryxContext(ctx, query, args...)
	}

	if err != nil {
		return nil, errors.InternalServer.Wrap(err, secondPathDepthOption...)
	}

	return rows, nil
}

func (s *DB) QueryRow(ctx context.Context, query string, args ...any) (*Row, error) {

	query, err := replacePlaceholders(query)
	if err != nil {
		return nil, err
	}

	row := &Row{Row: nil}
	if tx := extractTx(ctx); tx != nil {
		row.Row = tx.Tx.QueryRowxContext(ctx, query, args...)
	} else {
		row.Row = s.DB.QueryRowxContext(ctx, query, args...)
	}

	return row, nil
}

func (s *DB) Prepare(ctx context.Context, query string) (_ *Stmt, err error) {

	query, err = replacePlaceholders(query)
	if err != nil {
		return nil, err
	}

	var stmt = &Stmt{Stmt: nil}
	if tx := extractTx(ctx); tx != nil {
		stmt.Stmt, err = tx.Tx.PreparexContext(ctx, query)
	} else {
		stmt.Stmt, err = s.DB.PreparexContext(ctx, query)
	}

	if err != nil {
		return nil, errors.InternalServer.Wrap(err, secondPathDepthOption...)
	}

	return stmt, nil
}

func (s *DB) Exec(ctx context.Context, query string, args ...any) (err error) {

	query, err = replacePlaceholders(query)
	if err != nil {
		return err
	}

	if tx := extractTx(ctx); tx != nil {
		_, err = tx.Tx.ExecContext(ctx, query, args...)
	} else {
		_, err = s.DB.ExecContext(ctx, query, args...)
	}

	if err != nil {
		return errors.InternalServer.Wrap(err, secondPathDepthOption...)
	}

	return nil
}

func (s *DB) ExecWithLastInsertID(ctx context.Context, query string, args ...any) (id uint32, err error) {

	query += " RETURNING id"
	query, err = replacePlaceholders(query)
	if err != nil {
		return 0, err
	}

	if tx := extractTx(ctx); tx != nil {
		err = tx.Tx.GetContext(ctx, &id, query, args...)
	} else {
		err = s.DB.GetContext(ctx, &id, query, args...)
	}

	if err != nil {
		return 0, errors.InternalServer.Wrap(err, secondPathDepthOption...)
	}

	return id, nil
}

func (s *DB) ExecWithRowsAffected(ctx context.Context, query string, args ...any) (_ uint32, err error) {

	query, err = replacePlaceholders(query)
	if err != nil {
		return 0, err
	}

	var result sql.Result

	if tx := extractTx(ctx); tx != nil {
		result, err = tx.Tx.ExecContext(ctx, query, args...)
	} else {
		result, err = s.DB.ExecContext(ctx, query, args...)
	}

	if err != nil {
		return 0, errors.InternalServer.Wrap(err, secondPathDepthOption...)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return 0, errors.InternalServer.Wrap(err, secondPathDepthOption...)
	}

	return uint32(affected), nil
}
