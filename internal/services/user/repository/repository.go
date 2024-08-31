package repository

import (
	"server/pkg/sql"
)

type UserRepository struct {
	db sql.SQL
}

func NewUserRepository(db sql.SQL) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
