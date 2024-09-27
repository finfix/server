package service

import (
	"context"

	"server/internal/services/accountGroup/model"
)

// DeleteAccountGroup удаляет группу счетов
func (s *AccountGroupService) DeleteAccountGroup(ctx context.Context, id model.DeleteAccountGroupReq) error {

	// Проверяем доступ пользователя к счету
	if err := s.CheckAccess(ctx, id.Necessary.UserID, []uint32{id.ID}); err != nil {
		return err
	}

	// Отвязываем пользователя от группы счетов
	if err := s.accountGroupRepository.UnlinkUserFromAccountGroup(ctx, id.Necessary.UserID, id.ID); err != nil {
		return err
	}

	// Удаляем счет
	return s.accountGroupRepository.DeleteAccountGroup(ctx, id.ID)
}
