package service

import (
	"context"

	"server/internal/services/generalRepository/checker"
	"server/internal/services/transaction/model"
)

// DeleteTransaction удаляет транзакцию
func (s *Service) DeleteTransaction(ctx context.Context, id model.DeleteTransactionReq) error {

	// Проверяем доступ пользователя к транзакции
	if err := s.generalRepository.CheckUserAccessToObjects(ctx, checker.Transactions, id.Necessary.UserID, []uint32{id.ID}); err != nil {
		return err
	}

	// Удаляем транзакцию
	return s.transactionRepository.DeleteTransaction(ctx, id.ID, id.Necessary.UserID)
}
