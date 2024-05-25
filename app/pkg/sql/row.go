package sql

import (
	"github.com/jmoiron/sqlx"
)

type RowInterface interface {
	Scan(dest ...any) error
	scanner
}

type Row struct {
	Row *sqlx.Row
}

func (s *Row) Scan(dest ...any) error {
	if err := s.Row.Scan(dest...); err != nil {
		return wrapSQLError(err)
	}
	return nil
}

func (s *Row) StructScan(dest any) error {
	if err := s.Row.StructScan(dest); err != nil {
		return wrapSQLError(err)
	}
	return nil
}

func (s *Row) SliceScan() ([]any, error) {
	res, err := s.Row.SliceScan()
	if err != nil {
		return nil, wrapSQLError(err)
	}
	return res, nil

}

func (s *Row) MapScan(dest map[string]any) error {
	if err := s.Row.MapScan(dest); err != nil {
		return wrapSQLError(err)
	}
	return nil
}
