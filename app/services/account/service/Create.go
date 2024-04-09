package service

import (
	"context"

	"server/app/pkg/errors"
	model2 "server/app/services/account/model"
	"server/app/services/generalRepository/checker"
)

// Create создает новый счет
func (s *Service) Create(ctx context.Context, accountToCreate model2.CreateReq) (res model2.CreateRes, err error) {

	// Проверяем доступ пользователя к группе счетов
	if err = s.general.CheckAccess(ctx, checker.AccountGroups, accountToCreate.UserID, []uint32{accountToCreate.AccountGroupID}); err != nil {
		return res, err
	}

	// Создаем SQL-транзакцию
	err = s.general.WithinTransaction(ctx, func(ctxTx context.Context) error {

		// Создаем счет
		if res.ID, res.SerialNumber, err = s.accountRepository.Create(ctx, accountToCreate); err != nil {
			return err
		}

		// Если счет создался с остатком
		if accountToCreate.Remainder != 0 {

			// Получаем счет
			accounts, err := s.accountRepository.Get(ctx, model2.GetReq{IDs: []uint32{res.ID}})
			if err != nil {
				return err
			}
			if len(accounts) == 0 {
				return errors.NotFound.New("Счет не найден")
			}
			account := accounts[0]

			// Создаем меняем остаток счета созданием транзакции
			if err := s.accountService.ChangeRemainder(ctxTx, account, account.Remainder); err != nil {
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
