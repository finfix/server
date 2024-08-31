package service

import (
	"context"

	"server/internal/services/accountGroup/model"
)

// GetAccountGroups Возвращает все группы счетов пользователя
func (s *Service) GetAccountGroups(ctx context.Context, req model.GetAccountGroupsReq) ([]model.AccountGroup, error) {
	return s.accountGroupRepository.GetAccountGroups(ctx, req)
}
