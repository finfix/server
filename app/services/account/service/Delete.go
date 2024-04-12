package service

import (
	"context"

	"server/app/services/account/model"
	"server/app/services/generalRepository/checker"
)

// Delete удаляет счет
func (s *Service) Delete(ctx context.Context, id model.DeleteReq) error {

	// Проверяем доступ пользователя к счету
	if err := s.general.CheckAccess(ctx, checker.Accounts, id.Necessary.UserID, []uint32{id.ID}); err != nil {
		return err
	}

	// Удаляем счет
	return s.accountRepository.Delete(ctx, id.ID)
}
