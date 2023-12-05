package service

import (
	"context"

	"core/app/enum/transactionType"
	"core/app/internal/services/account/model"
	"core/app/internal/services/generalRepository/checker"
	transactionModel "core/app/internal/services/transaction/model"
	"pkg/datetime/date"
	"pkg/pointer"
)

// Create создает новый счет
func (s *Service) Create(ctx context.Context, account model.CreateReq) (id uint32, err error) {

	// Проверяем доступ пользователя к группе счетов
	if err = s.general.CheckAccess(ctx, checker.AccountGroups, account.UserID, []uint32{account.AccountGroupID}); err != nil {
		return id, err
	}

	// Создаем SQL-транзакцию
	err = s.general.WithinTransaction(ctx, func(ctxTx context.Context) error {

		// Создаем счет
		if id, err = s.account.Create(ctx, account); err != nil {
			return err
		}

		// Если на нем есть остаток, создаем транзакцию
		if account.Remainder != 0 {
			if _, err = s.transaction.Create(ctx, transactionModel.CreateReq{
				Type:            transactionType.Balancing,
				AmountTo:        account.Remainder,
				AccountToID:     id,
				DateTransaction: date.Now(),
				IsExecuted:      pointer.Pointer(true),
			}); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return id, err
	}

	return id, nil
}
