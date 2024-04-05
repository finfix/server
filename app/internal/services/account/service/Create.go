package service

import (
	"context"

	"server/app/enum/accountType"
	"server/app/enum/transactionType"
	"server/app/internal/services/account/model"
	"server/app/internal/services/generalRepository/checker"
	transactionModel "server/app/internal/services/transaction/model"
	"server/pkg/datetime/date"
	"server/pkg/errors"
	"server/pkg/pointer"
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
				Params: map[string]any{"accountID": id},
			})
		}
		balancingAccount := balancingAccounts[0]

		// Если на нем есть остаток, создаем транзакцию
		if account.Remainder != 0 {
			if _, err = s.transaction.Create(ctx, transactionModel.CreateReq{
				Type:            transactionType.Balancing,
				AmountTo:        account.Remainder,
				AccountFromID:   balancingAccount.ID,
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
