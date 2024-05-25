package service

import (
	"context"

	"server/app/services/accountGroup/model"
	"server/app/services/generalRepository/checker"
)

// Update обновляет счета по конкретным полям
func (s *Service) Update(ctx context.Context, updateReq model.UpdateAccountGroupReq) error {

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
