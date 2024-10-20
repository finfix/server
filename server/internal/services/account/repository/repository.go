package repository

import (
	"pkg/sql"
)

type AccountRepository struct {
	db sql.SQL
}

func NewAccountRepository(db sql.SQL, ) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}
