package repository

import (
	"server/pkg/sql"
)

type Repository struct {
	db sql.SQL
}

func New(db sql.SQL, ) *Repository {
	return &Repository{
		db: db,
	}
}
