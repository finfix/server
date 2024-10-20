package service

import (
	"context"

	"server/internal/services/transaction/model"
)

func (s *TransactionService) GetTransactions(ctx context.Context, filters model.GetTransactionsReq) (transactions []model.Transaction, err error) {

	// Проверяем доступ пользователя к группам счетов
	filters.AccountGroupIDs, err = s.userService.GetAccessedAccountGroups(ctx, filters.Necessary.UserID)
	if err != nil {
		return nil, err
	}

	// Если передан фильтр по счету
	if filters.AccountID != nil {

		// Проверяем доступ к этому счету
		if err = s.accountService.CheckAccess(ctx, filters.Necessary.UserID, []uint32{*filters.AccountID}); err != nil {
			return nil, err
		}
	}

	// Получаем все транзакции
	if transactions, err = s.transactionRepository.GetTransactions(ctx, filters); err != nil {
		return nil, err
	}

	// Заполняем массив ID транзакций
	transactionIDs := make([]uint32, len(transactions))
	for i, transaction := range transactions {
		transactionIDs[i] = transaction.ID
	}

	return transactions, nil
}
