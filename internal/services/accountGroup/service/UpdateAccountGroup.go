package service

import (
	"context"

	"server/internal/services/accountGroup/model"
	"server/internal/services/generalRepository/checker"
)

// UpdateAccountGroup обновляет группу счетов по конкретным полям
func (s *AccountGroupService) UpdateAccountGroup(ctx context.Context, updateReq model.UpdateAccountGroupReq) error {

	// Проверяем доступ пользователя к группе счетов
	if err := s.general.CheckUserAccessToObjects(ctx, checker.AccountGroups, updateReq.Necessary.UserID, []uint32{updateReq.ID}); err != nil {
		return err
	}

	return s.general.WithinTransaction(ctx, func(ctxTx context.Context) error {
		err := s.accountGroupRepository.UpdateAccountGroup(ctxTx, updateReq)
		if err != nil {
			return err
		}
		return nil
	})
}
