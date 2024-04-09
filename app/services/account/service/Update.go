package service

import (
	"context"

	"server/app/services/account/model"
	"server/app/services/generalRepository/checker"
	"server/pkg/errors"
)

// Update обновляет счета по конкретным полям
func (s *Service) Update(ctx context.Context, updateReq model.UpdateReq) error {

	// Проверяем доступ пользователя к счету
	if err := s.general.CheckAccess(ctx, checker.Accounts, updateReq.UserID, []uint32{updateReq.ID}); err != nil {
		return err
	}

	return s.general.WithinTransaction(ctx, func(ctxTx context.Context) error {
		return s.update(ctxTx, updateReq)
	})
}

func (s *Service) update(ctx context.Context, updateReq model.UpdateReq) error {

	// Получаем счет
	accounts, err := s.accountRepository.Get(ctx, model.GetReq{IDs: []uint32{updateReq.ID}})
	if err != nil {
		return err
	}
	if len(accounts) == 0 {
		return errors.NotFound.New("Счет не найден", errors.Options{
			Params: map[string]any{
				"accountID": updateReq.ID,
			},
		})
	}
	account := accounts[0]

	// Проверяем, что входные данные не противоречат разрешениям
	if err = s.permissionsService.CheckPermissions(updateReq, s.permissionsService.GetPermissions(account)); err != nil {
		return err
	}

	// Если редактируется остаток
	if updateReq.Remainder != nil {
		if err = s.accountService.ChangeRemainder(ctx, account, *updateReq.Remainder); err != nil {
			return err
		}
	}

	if updateReq.ParentAccountID != nil && *updateReq.ParentAccountID != 0 {
		if err = s.accountService.ValidateUpdateParentAccountID(ctx, account, *updateReq.ParentAccountID, updateReq.UserID); err != nil {
			return err
		}
	}

	// Редактируем счет
	return s.accountRepository.Update(ctx, updateReq)
}
