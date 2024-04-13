package service

import (
	"context"

	"server/app/pkg/errors"
	"server/app/services/account/model"
	accountRepoModel "server/app/services/account/repository/model"
	"server/app/services/generalRepository/checker"
)

// CreateAccount создает новый счет
func (s *Service) CreateAccount(ctx context.Context, accountToCreate model.CreateAccountReq) (res model.CreateAccountRes, err error) {

	// Проверяем доступ пользователя к группе счетов
	if err = s.general.CheckUserAccessToObjects(ctx, checker.AccountGroups, accountToCreate.Necessary.UserID, []uint32{accountToCreate.AccountGroupID}); err != nil {
		return res, err
	}

	// Проверяем, можно ли привязать счет к родительскому счету
	if accountToCreate.ParentAccountID != nil {

		// Представляем, что счет уже создан
		account := accountToCreate.ContertToAccount()

		// Проверяем возможность привязки
		if err = s.ValidateUpdateParentAccountID(ctx, account, *accountToCreate.ParentAccountID, accountToCreate.Necessary.UserID); err != nil {
			return res, err
		}
	}

	// Создаем SQL-транзакцию
	err = s.general.WithinTransaction(ctx, func(ctxTx context.Context) error {

		// Создаем счет
		if res.ID, res.SerialNumber, err = s.accountRepository.CreateAccount(ctx, accountToCreate.ConvertToRepoReq()); err != nil {
			return err
		}

		// Если счет создался с остатком
		if accountToCreate.Remainder != 0 {

			// Получаем счет
			accounts, err := s.accountRepository.GetAccounts(ctx, accountRepoModel.GetAccountsReq{IDs: []uint32{res.ID}})
			if err != nil {
				return err
			}
			if len(accounts) == 0 {
				return errors.NotFound.New("Счет не найден")
			}
			account := accounts[0]

			// Меняем остаток счета созданием транзакции
			updateRes, err := s.accountService.ChangeAccountRemainder(ctxTx, account, accountToCreate.Remainder, accountToCreate.Necessary.UserID)
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
