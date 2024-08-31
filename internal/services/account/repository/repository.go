package repository

import (
	"server/pkg/sql"
)

type AccountRepository struct {
	db sql.SQL
}

func NewAccountRepository(db sql.SQL, ) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}
