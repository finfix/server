package repository

import (
	"server/pkg/sql"
)

type TransactionRepository struct {
	db sql.SQL
}

func New(db sql.SQL, ) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}
