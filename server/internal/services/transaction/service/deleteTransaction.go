package service

import (
	"context"

	"server/internal/services/transaction/model"
)

// DeleteTransaction удаляет транзакцию
func (s *TransactionService) DeleteTransaction(ctx context.Context, id model.DeleteTransactionReq) error {

	// Проверяем доступ пользователя к транзакции
	if err := s.CheckAccess(ctx, id.Necessary.UserID, []uint32{id.ID}); err != nil {
		return err
	}

	// Удаляем транзакцию
	return s.transactionRepository.DeleteTransaction(ctx, id.ID, id.Necessary.UserID)
}
