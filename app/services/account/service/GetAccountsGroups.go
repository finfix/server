package service

import (
	"context"

	model2 "server/app/services/account/model"
)

// GetAccountGroups Возвращает все группы счетов пользователя
func (s *Service) GetAccountGroups(ctx context.Context, req model2.GetAccountGroupsReq) ([]model2.AccountGroup, error) {
	filters := model2.GetAccountGroupsReq{
		UserID:          req.UserID,
		AccountGroupIDs: req.AccountGroupIDs,
	}
	return s.account.GetAccountGroups(ctx, filters)
}
