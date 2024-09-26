package service

import (
	"context"

	"pkg/slices"

	"server/internal/services/account/model"
	accountRepoModel "server/internal/services/account/repository/model"
	"server/internal/services/generalRepository/checker"
)

// CreateAccount создает новый счет
func (s *AccountService) CreateAccount(ctx context.Context, accountToCreate model.CreateAccountReq) (res model.CreateAccountRes, err error) {

	// Проверяем доступ пользователя к группе счетов
	if err = s.general.CheckUserAccessToObjects(ctx, checker.AccountGroups, accountToCreate.Necessary.UserID, []uint32{accountToCreate.AccountGroupID}); err != nil {
		return res, err
	}

	// Проверяем, можно ли привязать счет к родительскому счету
	if accountToCreate.ParentAccountID != nil {

		// Представляем, что счет уже создан
		account := accountToCreate.ConvertToAccount()

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
		if !accountToCreate.Remainder.IsZero() {

			// Получаем счет
			account, err := slices.FirstWithError(s.accountRepository.GetAccounts(ctx,
				accountRepoModel.GetAccountsReq{ //nolint:exhaustruct
					IDs: []uint32{res.ID},
				},
			))
			if err != nil {
				return err
			}

			// Меняем остаток счета созданием транзакции
			updateRes, err := s.ChangeAccountRemainder(ctxTx, account, accountToCreate.Remainder, accountToCreate.Necessary.UserID)
			if err != nil {
				return err
			}
			res.BalancingTransactionID = updateRes.BalancingTransactionID
			res.BalancingAccountID = updateRes.BalancingAccountID
			res.BalancingAccountSerialNumber = updateRes.BalancingAccountSerialNumber
		}

		return nil
	})
	if err != nil {
		return res, err
	}

	return res, nil
}
