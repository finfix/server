package sql

import (
	"server/pkg/errors"

	"github.com/jmoiron/sqlx"
)

type Rows struct {
	*sqlx.Rows
}

type RowsInterface interface {
	scanner
	closer
}

func (s *Rows) SliceScan() ([]any, error) {
	res, err := s.Rows.SliceScan()
	if err != nil {
		return nil, errors.InternalServer.WrapPath(err, pathDepth)
	}
	return res, nil
}

func (s *Rows) MapScan(dest map[string]any) error {
	if err := s.Rows.MapScan(dest); err != nil {
		return errors.InternalServer.WrapPath(err, pathDepth)
	}
	return nil
}

func (s *Rows) StructScan(dest any) error {
	if err := s.Rows.StructScan(dest); err != nil {
		return errors.InternalServer.WrapPath(err, pathDepth)
	}
	return nil
}

func (s *Rows) Close() error {
	if err := s.Rows.Close(); err != nil {
		return errors.InternalServer.WrapPath(err, pathDepth)
	}
	return nil
}
