package repository

import (
	"server/pkg/sql"
)

type TransactionRepository struct {
	db sql.SQL
}

func NewTransactionRepository(db sql.SQL, ) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}
