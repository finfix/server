package service

import (
	"context"

	"server/app/internal/services/account/model"
	"server/app/internal/services/generalRepository/checker"
	"server/pkg/errors"
)

// Delete удаляет счет
func (s *Service) Delete(ctx context.Context, id model.DeleteReq) error {

	// Проверяем доступ пользователя к счету
	if err := s.general.CheckAccess(ctx, checker.Accounts, id.UserID, []uint32{id.ID}); err != nil {
		return err
	}

	// Получаем счет
	accounts, err := s.account.Get(ctx, model.GetReq{IDs: []uint32{id.ID}})
	if err != nil {
		return err
	}
	if len(accounts) == 0 {
		return errors.NotFound.New("Счет не найден")
	}
	account := accounts[0]

	// Проверяем, что счет можно удалять
	if !s.permissionsService.GetPermissions(account).DeleteAccount {
		return errors.BadRequest.New("Нельзя удалять счет")
	}

	// Удаляем счет
	return s.account.Delete(ctx, id.ID)
}
