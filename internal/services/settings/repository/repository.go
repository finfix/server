package repository

import (
	"server/pkg/sql"
)

type SettingsRepository struct {
	db sql.SQL
}

func NewSettingsRepository(db sql.SQL) *SettingsRepository {
	return &SettingsRepository{
		db: db,
	}
}
