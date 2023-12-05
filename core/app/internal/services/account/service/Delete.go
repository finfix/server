package service

import (
	"context"

	"core/app/internal/services/account/model"
	"core/app/internal/services/generalRepository/checker"
)

// Delete удаляет счет
func (s *Service) Delete(ctx context.Context, id model.DeleteReq) error {

	// Проверяем доступ пользователя к счету
	if err := s.general.CheckAccess(ctx, checker.Accounts, id.UserID, []uint32{id.ID}); err != nil {
		return err
	}

	// Удаляем счет
	return s.account.Delete(ctx, id.ID)
}
