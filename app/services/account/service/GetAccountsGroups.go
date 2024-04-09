package service

import (
	"context"

	"server/app/services/account/model"
)

// GetAccountGroups Возвращает все группы счетов пользователя
func (s *Service) GetAccountGroups(ctx context.Context, req model.GetAccountGroupsReq) ([]model.AccountGroup, error) {
	return s.accountRepository.GetAccountGroups(ctx, model.GetAccountGroupsReq{
		UserID:          req.UserID,
		AccountGroupIDs: req.AccountGroupIDs,
	})
}
