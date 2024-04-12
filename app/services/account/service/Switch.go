package service

import (
	"context"

	"server/app/services/account/model"
	"server/app/services/generalRepository/checker"
)

// SwitchAccountBetweenThemselves меняет местами два счета
func (s *Service) SwitchAccountBetweenThemselves(ctx context.Context, req model.SwitchAccountBetweenThemselvesReq) error {

	// Проверяем доступ пользователя к счетам
	if err := s.general.CheckUserAccessToObjects(ctx, checker.Accounts, req.Necessary.UserID, []uint32{req.ID1, req.ID2}); err != nil {
		return err
	}

	// Меняем местами счета
	return s.accountRepository.SwitchAccountsBetweenThemselves(ctx, req.ID1, req.ID2)
}
