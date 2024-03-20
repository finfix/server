package service

import (
	"context"

	"server/app/internal/services/account/model"
)

// GetAccountGroups Возвращает все группы счетов пользователя
func (s *Service) GetAccountGroups(ctx context.Context, req model.GetAccountGroupsReq) ([]model.AccountGroup, error) {
	filters := model.GetAccountGroupsReq{
		UserID:          req.UserID,
		AccountGroupIDs: req.AccountGroupIDs,
	}
	return s.account.GetAccountGroups(ctx, filters)
}
