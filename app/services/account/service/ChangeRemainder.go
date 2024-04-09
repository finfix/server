package service

import (
	"context"
	"math"

	"server/app/pkg/datetime/date"
	"server/app/pkg/errors"
	"server/app/pkg/pointer"
	"server/app/services/account/model"
	"server/app/services/account/model/accountType"
	transactionModel "server/app/services/transaction/model"
	"server/app/services/transaction/model/transactionType"
)

func (s *Service) ChangeRemainder(ctx context.Context, account model.Account, remainderToUpdate float64) error {

	// Получаем остаток счета
	currentRemainder, err := s.accountRepository.GetRemainder(ctx, account.ID)
	if err != nil {
		if err != nil {
			return err
		}
		return err
	}

	// Проверяем, что остаток счета не равен написанному
	if remainderToUpdate == currentRemainder {
		return errors.BadRequest.New("Остаток счета равен написанному")
	}

	// Получаем балансировочный счет группы, чтобы создать для нее транзакцию
	balancingAccounts, err := s.accountRepository.Get(ctx, model.GetReq{
		Type:            pointer.Pointer(accountType.Balancing),
		AccountGroupIDs: []uint32{account.AccountGroupID},
	})
	if err != nil {
		return err
	}
	if len(balancingAccounts) == 0 {
		return errors.NotFound.New("Не найден счет для балансировки для счета", errors.Options{
			Params: map[string]any{"accountID": account.ID},
		})
	}
	balancingAccount := balancingAccounts[0]

	const rounding = 0.0000001

	// Создаем транзакцию балансировки
	if _, err = s.transaction.Create(ctx, transactionModel.CreateReq{
		Type:            transactionType.Balancing,
		AmountTo:        math.Round((remainderToUpdate-currentRemainder)/rounding) * rounding,
		AccountToID:     account.ID,
		AccountFromID:   balancingAccount.ID,
		DateTransaction: date.Now(),
		IsExecuted:      pointer.Pointer(true),
	}); err != nil {
		return err
	}

	return nil
}
