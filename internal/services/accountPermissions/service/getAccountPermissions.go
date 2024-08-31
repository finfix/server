package service

import (
	"context"

	accountModel "server/internal/services/account/model"
	"server/internal/services/accountPermissions/model"
	"server/internal/services/accountPermissions/service/utils"
)

func (s *AccountPermissionsService) GetAccountsPermissions(ctx context.Context, accounts ...accountModel.Account) (permissions []model.AccountPermissions, err error) {
	permissionsSet, err := s.accountPermissionsRepository.GetAccountPermissions(ctx)
	if err != nil {
		return permissions, err
	}
	permissionsArr := make([]model.AccountPermissions, 0, len(accounts))
	for _, account := range accounts {
		permissionsArr = append(permissionsArr, utils.JoinAccountPermissions(
			permissionsSet.TypeToPermissions[account.Type],
			permissionsSet.IsParentToPermissions[account.IsParent],
		))
	}
	return permissionsArr, nil
}
