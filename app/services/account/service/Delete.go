package service

import (
	"context"

	"server/app/services/account/model"
	"server/app/services/generalRepository/checker"
)

// DeleteAccount удаляет счет
func (s *Service) DeleteAccount(ctx context.Context, id model.DeleteAccountReq) error {

	// Проверяем доступ пользователя к счету
	if err := s.general.CheckUserAccessToObjects(ctx, checker.Accounts, id.Necessary.UserID, []uint32{id.ID}); err != nil {
		return err
	}

	// Удаляем счет
	return s.accountRepository.DeleteAccount(ctx, id.ID)
}
