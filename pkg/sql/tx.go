package sql

import (
	"server/pkg/errors"

	"github.com/jmoiron/sqlx"
)

type TxInterface interface {
	Commit() error
	Rollback() error
}

type Tx struct {
	Tx *sqlx.Tx
}

func (s *Tx) Commit() error {
	if err := s.Tx.Commit(); err != nil {
		return errors.InternalServer.WrapPath(err, pathDepth)
	}
	return nil
}

func (s *Tx) Rollback() error {
	if err := s.Tx.Rollback(); err != nil {
		return errors.InternalServer.WrapPath(err, pathDepth)
	}
	return nil
}
