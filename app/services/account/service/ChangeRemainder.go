package service

import (
	"context"
	"math"

	"server/app/pkg/datetime/date"
	"server/app/pkg/errors"
	"server/app/pkg/pointer"
	"server/app/services/account/model"
	"server/app/services/account/model/accountType"
	accountRepoModel "server/app/services/account/repository/model"
	transactionModel "server/app/services/transaction/model"
	"server/app/services/transaction/model/transactionType"
)

func (s *Service) ChangeRemainder(ctx context.Context, account model.Account, remainderToUpdate float64) (res model.UpdateRes, err error) {

	// Получаем остаток счета
	remainders, err := s.accountRepository.CalculateRemainderAccounts(ctx, accountRepoModel.CalculateRemaindersAccountsReq{
		IDs: []uint32{account.ID},
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

	const rounding = 0.0000001

	roundedAmount := math.Round((remainderToUpdate-remainders[account.ID])/rounding) * rounding

	// Создаем транзакцию балансировки
	balancingAccountID, err = s.transaction.Create(ctx, transactionModel.CreateReq{
		Type:            transactionType.Balancing,
		AmountFrom:      roundedAmount,
		AmountTo:        roundedAmount,
		AccountToID:     account.ID,
		AccountFromID:   balancingAccountID,
		DateTransaction: date.Now(),
		IsExecuted:      pointer.Pointer(true),
	})
	if err != nil {
		return res, err
	}
	res.BalancingAccountID = &balancingAccountID

	return res, nil
}

// GetBalancingAccountID получает ID балансировочного счета, подходящего для конкретного счета
func (s *Service) GetBalancingAccountID(ctx context.Context, account model.Account) (balancingAccountID uint32, serialNumber uint32, wasCreate bool, err error) {

	// Получаем балансировочный счет группы в нужной валюте, чтобы создать для нее транзакцию
	balancingAccounts, err := s.accountRepository.Get(ctx, accountRepoModel.GetReq{
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
	parentBalancingAccounts, err := s.accountRepository.Get(ctx, accountRepoModel.GetReq{
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
		return balancingAccountID, serialNumber, wasCreate, errors.InternalServer.New("Родительский балансировочный счет не найден", errors.Options{Params: map[string]any{
			"accountID":      account,
			"accountGroupID": account.AccountGroupID,
		}})
	}

	parentBalancingAccount = parentBalancingAccounts[0]

	// Создаем балансировочный счет
	balancingAccountID, serialNumber, err = s.accountRepository.Create(ctx, accountRepoModel.CreateReq{
		Name:            "Балансировочный",
		Visible:         parentBalancingAccount.Visible,
		IconID:          0,
		Type:            accountType.Balancing,
		Currency:        account.Currency,
		AccountGroupID:  parentBalancingAccount.AccountGroupID,
		Accounting:      parentBalancingAccount.Accounting,
		IsParent:        false,
		ParentAccountID: &parentBalancingAccount.ID,
	})
	if err != nil {
		return balancingAccountID, serialNumber, wasCreate, err
	}
	wasCreate = true
	return balancingAccountID, serialNumber, wasCreate, nil
}
