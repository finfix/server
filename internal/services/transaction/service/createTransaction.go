package service

import (
	"context"

	accountModel "server/internal/services/account/model"
	accountRepoModel "server/internal/services/account/repository/model"
	"server/internal/services/generalRepository/checker"
	transactionModel "server/internal/services/transaction/model"
	"server/internal/services/transaction/service/utils"
	"server/pkg/errors"
	"server/pkg/slices"
)

// CreateTransaction создает новую транзакцию
func (s *TransactionService) CreateTransaction(ctx context.Context, transaction transactionModel.CreateTransactionReq) (id uint32, err error) {

	// Проверяем доступ пользователя к счетам
	if err = s.generalRepository.CheckUserAccessToObjects(ctx, checker.Accounts, transaction.Necessary.UserID, []uint32{transaction.AccountFromID, transaction.AccountToID}); err != nil {
		return id, err
	}

	// Получаем счета
	_accounts, err := s.accountRepository.GetAccounts(ctx, accountRepoModel.GetAccountsReq{ //nolint:exhaustruct
		IDs: []uint32{transaction.AccountFromID, transaction.AccountToID},
	})
	if err != nil {
		return id, err
	}
	accountsMap := slices.ToMap(_accounts, func(account accountModel.Account) uint32 { return account.ID })

	// Проверяем, может ли пользователь использовать счета
	if err = utils.TransactionAndAccountTypesValidation(
		accountsMap[transaction.AccountFromID],
		accountsMap[transaction.AccountToID],
		transaction.Type,
	); err != nil {
		return id, err
	}

	// Получаем разрешения счетов
	permissionsAccountFrom := s.permissionsService.GetAccountPermissions(accountsMap[transaction.AccountFromID])
	permissionsAccountTo := s.permissionsService.GetAccountPermissions(accountsMap[transaction.AccountToID])

	// Проверяем, что счета можно использовать для создания транзакции
	if !permissionsAccountFrom.CreateTransaction || !permissionsAccountTo.CreateTransaction {
		return id, errors.BadRequest.New("Нельзя создать транзакцию для этих счетов",
			errors.ParamsOption(
				"AccountFromID", transaction.AccountFromID,
				"AccountGroupFromID", accountsMap[transaction.AccountFromID].AccountGroupID,
				"AccountToID", transaction.AccountToID,
				"AccountGroupToID", accountsMap[transaction.AccountToID].AccountGroupID,
			),
		)
	}

	// Проверяем, что счета находятся в одной группе
	if accountsMap[transaction.AccountFromID].AccountGroupID != accountsMap[transaction.AccountToID].AccountGroupID {
		return id, errors.BadRequest.New("Счета находятся в разных группах",
			errors.ParamsOption(
				"AccountFromID", transaction.AccountFromID,
				"AccountToID", transaction.AccountToID,
			))
	}

	return id, s.generalRepository.WithinTransaction(ctx, func(ctxTx context.Context) error {

		// Создаем транзакцию
		id, err = s.transactionRepository.CreateTransaction(ctx, transaction.ConvertToRepoReq())
		if err != nil {
			return err
		}

		// Если переданы теги
		if len(transaction.TagIDs) != 0 {
			if err = s.updateTransactionTags(ctx, transaction.Necessary.UserID, id, transaction.TagIDs); err != nil {
				return err
			}
		}
		return nil
	})
}
