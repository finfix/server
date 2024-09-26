package repository

import (
	"time"

	"pkg/cache"
	"pkg/sql"

	"server/internal/services/accountPermissions/model"
)

type AccountPermissionsRepository struct {
	db    sql.SQL
	cache *cache.Cache[struct{}, model.PermissionSet]
}

func NewAccountPermissionsRepository(db sql.SQL) *AccountPermissionsRepository {
	return &AccountPermissionsRepository{
		db:    db,
		cache: cache.NewCache[struct{}, model.PermissionSet](time.Minute),
	}
}
