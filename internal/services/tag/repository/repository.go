package repository

import (
	"server/pkg/sql"
)

type TagRepository struct {
	db sql.SQL
}

func New(db sql.SQL, ) *TagRepository {
	return &TagRepository{
		db: db,
	}
}
