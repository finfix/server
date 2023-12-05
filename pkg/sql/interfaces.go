package sql

import "context"

type general interface {
	Get(ctx context.Context, dest any, query string, args ...any) error
	Select(ctx context.Context, dest any, query string, args ...any) error
	Query(ctx context.Context, query string, args ...any) (*Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) *Row
	Exec(ctx context.Context, query string, args ...any) error
	ExecWithLastInsertID(ctx context.Context, query string, args ...any) (uint32, error)
	ExecWithRowsAffected(ctx context.Context, query string, args ...any) (uint32, error)
	Prepare(ctx context.Context, query string) (*Stmt, error)
}

type scanner interface {
	SliceScan() ([]any, error)
	MapScan(dest map[string]any) error
	StructScan(dest any) error
}

type closer interface {
	Close() error
}
