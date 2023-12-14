package service

import (
	"context"
	"pkg/errors"

	"core/app/internal/services/account/model"
	"core/app/internal/services/generalRepository/checker"
)

// Delete удаляет счет
func (s *Service) Delete(ctx context.Context, id model.DeleteReq) error {

	// Проверяем доступ пользователя к счету
	if err := s.general.CheckAccess(ctx, checker.Accounts, id.UserID, []uint32{id.ID}); err != nil {
		return err
	}

	// Получаем разрешения счета
	permissions, err := s.GetPermissions(ctx, id.ID)
	if err != nil {
		return err
	}

	// Проверяем, что счет можно удалять
	if !permissions.DeleteAccount {
		return errors.BadRequest.New("Нельзя удалять счет")
	}

	// Удаляем счет
	return s.account.Delete(ctx, id.ID)
}
