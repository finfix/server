package repository

import (
	"server/app/pkg/sql"
)

type Repository struct {
	db sql.SQL
}

func New(db sql.SQL, ) *Repository {
	return &Repository{
		db: db,
	}
}
