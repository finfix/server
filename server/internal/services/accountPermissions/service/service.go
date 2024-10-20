package service

import (
	"context"

	"server/internal/services/accountPermissions/model"
)

type AccountPermissionsService struct {
	accountPermissionsRepository AccountPermissionsRepository
}

type AccountPermissionsRepository interface {
	GetAccountPermissions(context.Context) (model.PermissionSet, error)
}

func NewAccountPermissionsService(accountPermissionsRepository AccountPermissionsRepository) *AccountPermissionsService {
	return &AccountPermissionsService{
		accountPermissionsRepository: accountPermissionsRepository,
	}
}
