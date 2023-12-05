package sql

import (
	"context"
	"github.com/jmoiron/sqlx"
	"pkg/errors"
)

type StmtInterface interface {
	Select(ctx context.Context, dest any, args ...any) error
	Get(ctx context.Context, dest any, args ...any) error
	ExecWithLastInsertId(ctx context.Context, args ...any) (uint32, error)
	ExecWithAffectedRows(ctx context.Context, args ...any) (uint32, error)
	QueryRow(ctx context.Context, args ...any) RowInterface
	Query(ctx context.Context, args ...any) (RowsInterface, error)
	closer
}

type Stmt struct {
	Stmt *sqlx.Stmt
}

func (s *Stmt) Select(ctx context.Context, dest any, args ...any) error {
	if err := s.Stmt.SelectContext(ctx, dest, args...); err != nil {
		return errors.InternalServer.WrapPath(err, pathDepth)
	}
	return nil
}

func (s *Stmt) Get(ctx context.Context, dest any, args ...any) error {
	if err := s.Stmt.GetContext(ctx, dest, args...); err != nil {
		return errors.InternalServer.WrapPath(err, pathDepth)
	}
	return nil
}

func (s *Stmt) Exec(ctx context.Context, args ...any) error {
	_, err := s.Stmt.ExecContext(ctx, args...)
	if err != nil {
		return errors.InternalServer.WrapPath(err, pathDepth)
	}
	return nil
}

func (s *Stmt) ExecWithLastInsertId(ctx context.Context, args ...any) (uint32, error) {
	res, err := s.Stmt.ExecContext(ctx, args...)
	if err != nil {
		return 0, errors.InternalServer.WrapPath(err, pathDepth)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.InternalServer.WrapPath(err, pathDepth)
	}
	return uint32(id), nil
}

func (s *Stmt) ExecWithAffectedRows(ctx context.Context, args ...any) (uint32, error) {
	res, err := s.Stmt.ExecContext(ctx, args...)
	if err != nil {
		return 0, errors.InternalServer.WrapPath(err, pathDepth)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return 0, errors.InternalServer.WrapPath(err, pathDepth)
	}
	return uint32(rows), nil
}

func (s *Stmt) QueryRow(ctx context.Context, args ...any) *Row {
	return &Row{s.Stmt.QueryRowxContext(ctx, args...)}
}

func (s *Stmt) Query(ctx context.Context, args ...any) (*Rows, error) {
	rows, err := s.Stmt.QueryxContext(ctx, args...)
	if err != nil {
		return nil, errors.InternalServer.WrapPath(err, pathDepth)
	}
	return &Rows{rows}, nil
}

func (s *Stmt) Close() error {
	if err := s.Stmt.Close(); err != nil {
		return errors.InternalServer.WrapPath(err, pathDepth)
	}
	return nil
}
