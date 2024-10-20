package service

import (
	"context"
	"time"

	"github.com/shopspring/decimal"

	"pkg/datetime"
	"pkg/errors"

	"server/internal/services/account/model"
	accountRepoModel "server/internal/services/account/repository/model"
	"server/internal/services/transaction/model/transactionType"
	transactionRepoModel "server/internal/services/transaction/repository/model"
)

func (s *AccountService) ChangeAccountRemainder(ctx context.Context, account model.Account, remainderToUpdate decimal.Decimal, userID uint32) (res model.UpdateAccountRes, err error) {

	// Получаем остаток счета
	remainders, err := s.accountRepository.GetSumAllTransactionsToAccount(ctx, accountRepoModel.CalculateRemaindersAccountsReq{
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
	balancingTransactionID, err := s.transactionRepository.CreateTransaction(ctx, transactionRepoModel.CreateTransactionReq{
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
		AccountGroupID:     account.AccountGroupID,
	})
	if err != nil {
		return res, err
	}
	res.BalancingTransactionID = &balancingTransactionID

	return res, nil
}
