package service

import (
	"context"
	"server/app/internal/services/generalRepository/checker"

	"server/app/internal/services/account/model"
)

// Switch меняет местами два счета
func (s *Service) Switch(ctx context.Context, req model.SwitchReq) error {

	// Проверяем доступ пользователя к счетам
	if err := s.general.CheckAccess(ctx, checker.Accounts, req.UserID, []uint32{req.ID1, req.ID2}); err != nil {
		return err
	}

	// Меняем местами счета
	return s.account.Switch(ctx, req.ID1, req.ID2)
}
