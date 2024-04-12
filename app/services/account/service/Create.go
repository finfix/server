package service

import (
	"context"

	"server/app/pkg/errors"
	"server/app/services/account/model"
	accountRepoModel "server/app/services/account/repository/model"
	"server/app/services/generalRepository/checker"
)

// Create создает новый счет
func (s *Service) Create(ctx context.Context, accountToCreate model.CreateReq) (res model.CreateRes, err error) {

	// Проверяем доступ пользователя к группе счетов
	if err = s.general.CheckAccess(ctx, checker.AccountGroups, accountToCreate.Necessary.UserID, []uint32{accountToCreate.AccountGroupID}); err != nil {
		return res, err
	}

	// Создаем SQL-транзакцию
	err = s.general.WithinTransaction(ctx, func(ctxTx context.Context) error {

		// Создаем счет
		if res.ID, res.SerialNumber, err = s.accountRepository.Create(ctx, accountToCreate.ConvertToRepoReq()); err != nil {
			return err
		}

		// Если счет создался с остатком
		if accountToCreate.Remainder != 0 {

			// Получаем счет
			accounts, err := s.accountRepository.Get(ctx, accountRepoModel.GetReq{IDs: []uint32{res.ID}})
			if err != nil {
				return err
			}
			if len(accounts) == 0 {
				return errors.NotFound.New("Счет не найден")
			}
			account := accounts[0]

			// Меняем остаток счета созданием транзакции
			updateRes, err := s.accountService.ChangeRemainder(ctxTx, account, account.Remainder)
			if err != nil {
				return err
			}
			res.BalancingTransactionID = updateRes.BalancingTransactionID
		}

		return nil
	})
	if err != nil {
		return res, err
	}

	return res, nil
}
