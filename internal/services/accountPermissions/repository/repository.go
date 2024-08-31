package repository

import (
	"time"

	"server/internal/services/accountPermissions/model"
	"server/pkg/cache"
	"server/pkg/sql"
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
