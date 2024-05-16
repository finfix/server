package service

import (
	"context"
	"time"

	"github.com/shopspring/decimal"

	"server/app/pkg/datetime"
	"server/app/pkg/errors"
	"server/app/pkg/pointer"
	"server/app/services/account/model"
	"server/app/services/account/model/accountType"
	accountRepoModel "server/app/services/account/repository/model"
	"server/app/services/transaction/model/transactionType"
	transactionRepoModel "server/app/services/transaction/repository/model"
)

func (s *Service) ChangeAccountRemainder(ctx context.Context, account model.Account, remainderToUpdate decimal.Decimal, userID uint32) (res model.UpdateAccountRes, err error) {

	// Получаем остаток счета
	remainders, err := s.accountRepository.CalculateRemainderAccounts(ctx, accountRepoModel.CalculateRemaindersAccountsReq{
		IDs:             []uint32{account.ID},
		AccountGroupIDs: nil,
		Types:           nil,
		DateFrom:        nil,
		DateTo:          nil,
	})
	if err != nil {
		return res, err
	}

	// Проверяем, что остаток счета не равен написанному
	if remainderToUpdate == remainders[account.ID] {
		return res, errors.BadRequest.New("Остаток счета равен написанному")
	}

	// Получаем балансировочный счет
	balancingAccountID, serialNumber, wasCreate, err := s.GetBalancingAccountID(ctx, account)
	if err != nil {
		return res, err
	}
	if wasCreate {
		res.BalancingAccountID = &balancingAccountID
		res.BalancingTransactionID = &serialNumber
	}

	amount := remainderToUpdate.Sub(remainders[account.ID])

	// Создаем транзакцию балансировки
	balancingTransactionID, err := s.transaction.CreateTransaction(ctx, transactionRepoModel.CreateTransactionReq{
		Type:               transactionType.Balancing,
		AmountFrom:         amount,
		AmountTo:           amount,
		Note:               "",
		AccountToID:        account.ID,
		AccountFromID:      balancingAccountID,
		DateTransaction:    datetime.Date{Time: time.Now()},
		IsExecuted:         true,
		CreatedByUserID:    userID,
		DatetimeCreate:     time.Now(),
		AccountingInCharts: true,
	})
	if err != nil {
		return res, err
	}
	res.BalancingTransactionID = &balancingTransactionID

	return res, nil
}

// GetBalancingAccountID получает ID балансировочного счета, подходящего для конкретного счета
func (s *Service) GetBalancingAccountID(ctx context.Context, account model.Account) (balancingAccountID uint32, serialNumber uint32, wasCreate bool, err error) {

	// Получаем балансировочный счет группы в нужной валюте, чтобы создать для нее транзакцию
	balancingAccounts, err := s.accountRepository.GetAccounts(ctx, accountRepoModel.GetAccountsReq{ //nolint:exhaustruct
		Types:           []accountType.Type{accountType.Balancing},
		AccountGroupIDs: []uint32{account.AccountGroupID},
		Currencies:      []string{account.Currency},
		IsParent:        pointer.Pointer(false),
	})
	if err != nil {
		return balancingAccountID, serialNumber, wasCreate, err
	}
	// Если счет найден
	if len(balancingAccounts) != 0 {
		// Возвращаем его ID
		return balancingAccounts[0].ID, serialNumber, wasCreate, nil
	}

	// Получаем общий балансировочный счет
	parentBalancingAccounts, err := s.accountRepository.GetAccounts(ctx, accountRepoModel.GetAccountsReq{ //nolint:exhaustruct
		Types:           []accountType.Type{accountType.Balancing},
		AccountGroupIDs: []uint32{account.AccountGroupID},
		IsParent:        pointer.Pointer(true),
	})
	if err != nil {
		return balancingAccountID, serialNumber, wasCreate, err
	}

	var parentBalancingAccount model.Account

	// Если общий балансировочный счет не найден
	if len(parentBalancingAccounts) == 0 {
		return balancingAccountID, serialNumber, wasCreate, errors.InternalServer.New("Родительский балансировочный счет не найден", []errors.Option{errors.ParamsOption(
			"accountID", account,
			"accountGroupID", account.AccountGroupID,
		)}...)
	}

	parentBalancingAccount = parentBalancingAccounts[0]

	// Создаем балансировочный счет
	balancingAccountID, serialNumber, err = s.accountRepository.CreateAccount(ctx, accountRepoModel.CreateAccountReq{
		Budget: accountRepoModel.CreateReqBudget{
			Amount:         decimal.Zero,
			GradualFilling: false,
			FixedSum:       decimal.Zero,
			DaysOffset:     0,
		},
		Name:               "Балансировочный",
		Visible:            parentBalancingAccount.Visible,
		IconID:             parentBalancingAccount.IconID,
		Type:               accountType.Balancing,
		Currency:           account.Currency,
		AccountGroupID:     parentBalancingAccount.AccountGroupID,
		AccountingInHeader: parentBalancingAccount.AccountingInHeader,
		AccountingInCharts: true,
		IsParent:           false,
		ParentAccountID:    &parentBalancingAccount.ID,
		UserID:             account.CreatedByUserID,
		DatetimeCreate:     time.Now(),
	})
	if err != nil {
		return balancingAccountID, serialNumber, wasCreate, err
	}
	wasCreate = true
	return balancingAccountID, serialNumber, wasCreate, nil
}
