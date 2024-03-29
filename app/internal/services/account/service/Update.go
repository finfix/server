package service

import (
	"context"
	"math"
	"server/app/enum/transactionType"
	"server/app/internal/services/account/model"
	"server/app/internal/services/generalRepository/checker"
	transactionModel "server/app/internal/services/transaction/model"
	"server/pkg/datetime/date"
	"server/pkg/errors"
	"server/pkg/pointer"
)

// Update обновляет счета по конкретным полям
func (s *Service) Update(ctx context.Context, account model.UpdateReq) error {

	// Проверяем доступ пользователя к счету
	if err := s.general.CheckAccess(ctx, checker.Accounts, account.UserID, []uint32{account.ID}); err != nil {
		return err
	}

	// Получаем разрешения счета
	permissions, err := s.GetPermissions(ctx, account.ID)
	if err != nil {
		return err
	}

	// Проверяем, что входные данные не противоречат разрешениям
	if err = s.CheckPermissions(account, permissions); err != nil {
		return err
	}

	return s.general.WithinTransaction(ctx, func(ctxTx context.Context) error {

		// Если есть остаток
		if account.Remainder != nil {

			// Получаем остаток счета
			remainder, err := s.account.GetRemainder(ctx, account.ID)
			if err != nil {
				if err != nil {
					return err
				}
				return err
			}

			// Проверяем, что остаток счета не равен написанному
			if *account.Remainder-remainder == 0 {
				return errors.BadRequest.New("Остаток счета равен написанному")
			}

			// Создаем транзакцию балансировки
			if _, err = s.transaction.Create(ctx, transactionModel.CreateReq{
				Type:            transactionType.Balancing,
				AmountTo:        math.Round((*account.Remainder-remainder)*10000000) / 10000000,
				AccountToID:     account.ID,
				DateTransaction: date.Now(),
				IsExecuted:      pointer.Pointer(true),
			}); err != nil {
				return err
			}
		}

		// Редактируем счет
		return s.account.Update(ctx, account)
	})
}
