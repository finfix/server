package repository

import (
	"server/pkg/sql"
)

type AccountGroupRepository struct {
	db sql.SQL
}

func NewAccountGroupRepository(db sql.SQL, ) *AccountGroupRepository {
	return &AccountGroupRepository{
		db: db,
	}
}
