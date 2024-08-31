package repository

import (
	"server/pkg/sql"
)

type TagRepository struct {
	db sql.SQL
}

func NewTagRepository(db sql.SQL, ) *TagRepository {
	return &TagRepository{
		db: db,
	}
}
