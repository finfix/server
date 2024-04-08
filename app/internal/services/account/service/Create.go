package service

import (
	"context"

	"server/app/internal/services/account/model"
	"server/app/internal/services/generalRepository/checker"
	"server/pkg/errors"
)

// Create создает новый счет
func (s *Service) Create(ctx context.Context, accountToCreate model.CreateReq) (res model.CreateRes, err error) {

	// Проверяем доступ пользователя к группе счетов
	if err = s.general.CheckAccess(ctx, checker.AccountGroups, accountToCreate.UserID, []uint32{accountToCreate.AccountGroupID}); err != nil {
		return res, err
	}

	// Создаем SQL-транзакцию
	err = s.general.WithinTransaction(ctx, func(ctxTx context.Context) error {

		// Создаем счет
		if res.ID, res.SerialNumber, err = s.account.Create(ctx, accountToCreate); err != nil {
			return err
		}

		// Если счет создался с остатком
		if accountToCreate.Remainder != 0 {

			// Получаем счет
			accounts, err := s.account.Get(ctx, model.GetReq{IDs: []uint32{res.ID}})
			if err != nil {
				return err
			}
			if len(accounts) == 0 {
				return errors.NotFound.New("Счет не найден")
			}
			account := accounts[0]

			// Создаем меняем остаток счета созданием транзакции
			if err := s.changeRemainder(ctxTx, account, account.Remainder); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return res, err
	}

	return res, nil
}
