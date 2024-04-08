package service

import (
	"context"
	"math"

	"server/app/enum/accountType"
	"server/app/enum/transactionType"
	"server/app/internal/services/account/model"
	"server/app/internal/services/generalRepository/checker"
	transactionModel "server/app/internal/services/transaction/model"
	"server/pkg/datetime/date"
	"server/pkg/errors"
	"server/pkg/pointer"
)

// Update обновляет счета по конкретным полям
func (s *Service) Update(ctx context.Context, accountFields model.UpdateReq) error {

	// Проверяем доступ пользователя к счету
	if err := s.general.CheckAccess(ctx, checker.Accounts, accountFields.UserID, []uint32{accountFields.ID}); err != nil {
		return err
	}

	return s.general.WithinTransaction(ctx, func(ctxTx context.Context) error {

		// Получаем счет
		accounts, err := s.account.Get(ctx, model.GetReq{IDs: []uint32{accountFields.ID}})
		if err != nil {
			return err
		}
		if len(accounts) == 0 {
			return errors.NotFound.New("Счет не найден")
		}
		account := accounts[0]

		// Проверяем, что входные данные не противоречат разрешениям
		if err = s.permissionsService.CheckPermissions(accountFields, s.permissionsService.GetPermissions(account)); err != nil {
			return err
		}

		// Если редактируется остаток
		if accountFields.Remainder != nil {
			if err = s.changeRemainder(ctxTx, account, *accountFields.Remainder); err != nil {
				return err
			}
		}

		// Редактируем счет
		return s.account.Update(ctx, accountFields)
	})
}

func (s *Service) changeRemainder(ctx context.Context, account model.Account, remainderToUpdate float64) error {
	// Получаем остаток счета
	currentRemainder, err := s.account.GetRemainder(ctx, account.ID)
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
	balancingAccounts, err := s.account.Get(ctx, model.GetReq{
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
