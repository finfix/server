package sql

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"pkg/errors"
)

var _ SQL = &DB{DB: nil}

type SQL interface {
	Unsafe() *DB
	Begin(context.Context) (*Tx, error)
	Ping() error
	Get(ctx context.Context, dest any, q sq.Sqlizer) error
	Select(ctx context.Context, dest any, q sq.Sqlizer) error
	Query(ctx context.Context, q sq.Sqlizer) (*Rows, error)
	QueryRow(ctx context.Context, q sq.Sqlizer) (*Row, error)
	Exec(ctx context.Context, q sq.Sqlizer) error
	ExecWithLastInsertID(ctx context.Context, q sq.Sqlizer) (uint32, error)
	ExecWithRowsAffected(ctx context.Context, q sq.Sqlizer) (uint32, error)
	Prepare(ctx context.Context, q sq.Sqlizer) (*Stmt, error)
	closer
}

type DB struct {
	DB *sqlx.DB
}

func Open(driverName string, url string) (*DB, error) {
	db, err := sqlx.Open(driverName, url)
	if err != nil {
		return nil, wrapSQLError(err)
	}
	return &DB{db}, nil
}

func (s *DB) Close() error {
	if err := s.DB.Close(); err != nil {
		return wrapSQLError(err)
	}
	return nil
}

func (s *DB) Begin(ctx context.Context) (*Tx, error) {
	tx, err := s.DB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, wrapSQLError(err)
	}
	return &Tx{tx}, nil
}

func (s *DB) Ping() error {
	if err := s.DB.Ping(); err != nil {
		return wrapSQLError(err)
	}
	return nil
}

func (s *DB) Unsafe() *DB {
	return &DB{s.DB.Unsafe()}
}

func (s *DB) Select(ctx context.Context, dest any, q sq.Sqlizer) (err error) {

	// Формируем запрос из билдера
	query, args, err := q.ToSql()
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}

	// Заменяем все ? на $1, $2 и т.д.
	query, err = sq.Dollar.ReplacePlaceholders(query)

	if err != nil {
		return errors.InternalServer.Wrap(err)
	}

	// Извлекаем транзакцию из контекста
	if tx := extractTx(ctx); tx != nil {

		// Выполняем запрос в рамках транзакции
		err = tx.Tx.SelectContext(ctx, dest, query, args...)
	} else {

		// Выполняем запрос
		err = s.DB.SelectContext(ctx, dest, query, args...)
	}

	// Обрабатываем ошибки
	if err != nil {
		return wrapSQLError(err)
	}

	return nil
}

func (s *DB) Get(ctx context.Context, dest any, q sq.Sqlizer) (err error) {

	// Формируем запрос из билдера
	query, args, err := q.ToSql()
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}

	// Заменяем все ? на $1, $2 и т.д.
	query, err = sq.Dollar.ReplacePlaceholders(query)

	if err != nil {
		return errors.InternalServer.Wrap(err)
	}

	// Извлекаем транзакцию из контекста
	if tx := extractTx(ctx); tx != nil {

		// Выполняем запрос в рамках транзакции
		err = tx.Tx.GetContext(ctx, dest, query, args...)
	} else {

		// Выполняем запрос
		err = s.DB.GetContext(ctx, dest, query, args...)
	}

	// Обрабатываем ошибки
	if err != nil {
		return wrapSQLError(err)
	}

	return nil
}

func (s *DB) Query(ctx context.Context, q sq.Sqlizer) (_ *Rows, err error) {

	// Формируем запрос из билдера
	query, args, err := q.ToSql()
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	// Заменяем все ? на $1, $2 и т.д.
	query, err = sq.Dollar.ReplacePlaceholders(query)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	rows := &Rows{Rows: nil}

	// Извлекаем транзакцию из контекста
	if tx := extractTx(ctx); tx != nil {

		// Выполняем запрос в рамках транзакции
		rows.Rows, err = tx.Tx.QueryxContext(ctx, query, args...)
	} else {

		// Выполняем запрос
		rows.Rows, err = s.DB.QueryxContext(ctx, query, args...)
	}

	// Обрабатываем ошибки
	if err != nil {
		return nil, wrapSQLError(err)
	}

	return rows, nil
}

func (s *DB) QueryRow(ctx context.Context, q sq.Sqlizer) (*Row, error) {

	// Формируем запрос из билдера
	query, args, err := q.ToSql()
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	// Заменяем все ? на $1, $2 и т.д.
	query, err = sq.Dollar.ReplacePlaceholders(query)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	row := &Row{Row: nil}

	// Извлекаем транзакцию из контекста
	if tx := extractTx(ctx); tx != nil {

		// Выполняем запрос в рамках транзакции
		row.Row = tx.Tx.QueryRowxContext(ctx, query, args...)
	} else {

		// Выполняем запрос
		row.Row = s.DB.QueryRowxContext(ctx, query, args...)
	}

	return row, nil
}

func (s *DB) Prepare(ctx context.Context, q sq.Sqlizer) (_ *Stmt, err error) {

	// Формируем запрос из билдера
	query, _, err := q.ToSql()
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	// Заменяем все ? на $1, $2 и т.д.
	query, err = sq.Dollar.ReplacePlaceholders(query)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	var stmt = &Stmt{Stmt: nil}

	// Извлекаем транзакцию из контекста
	if tx := extractTx(ctx); tx != nil {

		// Подготавливаем запрос в рамках транзакции
		stmt.Stmt, err = tx.Tx.PreparexContext(ctx, query)
	} else {

		// Подготавливаем запрос
		stmt.Stmt, err = s.DB.PreparexContext(ctx, query)
	}

	// Обрабатываем ошибки
	if err != nil {
		return nil, wrapSQLError(err)
	}

	return stmt, nil
}

func (s *DB) Exec(ctx context.Context, q sq.Sqlizer) (err error) {

	// Формируем запрос из билдера
	query, args, err := q.ToSql()
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}

	// Заменяем все ? на $1, $2 и т.д.
	query, err = sq.Dollar.ReplacePlaceholders(query)

	if err != nil {
		return errors.InternalServer.Wrap(err)
	}

	// Извлекаем транзакцию из контекста
	if tx := extractTx(ctx); tx != nil {

		// Исполняем запрос в рамках транзакции
		_, err = tx.Tx.ExecContext(ctx, query, args...)
	} else {

		// Исполняем запрос
		_, err = s.DB.ExecContext(ctx, query, args...)
	}

	// Обрабатываем ошибки
	if err != nil {
		return wrapSQLError(err)
	}

	return nil
}

func (s *DB) ExecWithLastInsertID(ctx context.Context, q sq.Sqlizer) (id uint32, err error) {

	// Формируем запрос из билдера
	query, args, err := q.ToSql()
	if err != nil {
		return 0, errors.InternalServer.Wrap(err)
	}

	query += " RETURNING id"

	// Заменяем все ? на $1, $2 и т.д.
	query, err = sq.Dollar.ReplacePlaceholders(query)

	if err != nil {
		return 0, errors.InternalServer.Wrap(err)
	}

	// Извлекаем транзакцию из контекста
	if tx := extractTx(ctx); tx != nil {

		// Исполняем запрос в рамках транзакции
		err = tx.Tx.GetContext(ctx, &id, query, args...)
	} else {

		// Исполняем запрос
		err = s.DB.GetContext(ctx, &id, query, args...)
	}

	// Обрабатываем ошибки
	if err != nil {
		return 0, wrapSQLError(err)
	}

	return id, nil
}

func (s *DB) ExecWithRowsAffected(ctx context.Context, q sq.Sqlizer) (_ uint32, err error) {

	// Формируем запрос из билдера
	query, args, err := q.ToSql()
	if err != nil {
		return 0, errors.InternalServer.Wrap(err)
	}

	// Заменяем все ? на $1, $2 и т.д.
	query, err = sq.Dollar.ReplacePlaceholders(query)

	if err != nil {
		return 0, errors.InternalServer.Wrap(err)
	}

	var result sql.Result

	// Извлекаем транзакцию из контекста
	if tx := extractTx(ctx); tx != nil {

		// Исполняем запрос в рамках транзакции
		result, err = tx.Tx.ExecContext(ctx, query, args...)
	} else {

		// Исполняем запрос
		result, err = s.DB.ExecContext(ctx, query, args...)
	}

	// Обрабатываем ошибки
	if err != nil {
		return 0, wrapSQLError(err)
	}

	// Получаем количество затронутых строк
	affected, err := result.RowsAffected()
	if err != nil {
		return 0, wrapSQLError(err)
	}

	return uint32(affected), nil
}

func wrapSQLError(err error) error {

	thirdPathDepthOption := errors.SkipPreviousCallerOption()

	switch {
	case errors.Is(err, context.Canceled):
		return errors.ClientReject.Wrap(err, thirdPathDepthOption)
	case errors.Is(err, sql.ErrNoRows):
		return errors.NotFound.Wrap(err, thirdPathDepthOption)
	default:
		return errors.InternalServer.Wrap(err, thirdPathDepthOption)
	}
}
