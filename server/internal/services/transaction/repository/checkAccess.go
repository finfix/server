package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"pkg/errors"
	"server/internal/services/transaction/repository/transactionDDL"
)

// CheckAccess проверяет, имеет ли набор групп счетов пользователя доступ к указанным идентификаторам транзакций
func (r *TransactionRepository) CheckAccess(ctx context.Context, accountGroupIDs, transactionIDs []uint32) error {

	// Получаем все доступные транзакции по группам счетов и перечисленным транзакциям
	rows, err := r.db.Query(ctx, sq.
		Select(transactionDDL.ColumnID).
		From(transactionDDL.TableWithAlias).
		Where(sq.Eq{
			transactionDDL.ColumnAccountGroupID: accountGroupIDs,
			transactionDDL.ColumnID:             transactionIDs,
		}))
	if err != nil {
		return err
	}

	// Формируем мапу доступных транзакций
	accessedTransactionIDs := make(map[uint32]struct{})

	// Проходимся по каждой доступной транзакции
	for rows.Next() {

		// Считываем ID транзакции
		var transactionID uint32
		if err = rows.Scan(&transactionID); err != nil {
			return err
		}

		// Добавляем ID транзакции в мапу
		accessedTransactionIDs[transactionID] = struct{}{}
	}

	if len(accessedTransactionIDs) == 0 {
		return errors.Forbidden.New("You don't have access to any of the requested transactions",
			errors.ParamsOption("AccountGroupIDs", accountGroupIDs, "TransactionIDs", transactionIDs),
		)
	}

	// Проходимся по каждой запрашиваемой транзакции
	for _, transactionID := range transactionIDs {

		// Если счета нет в мапе доступных транзакций, возвращаем ошибку
		if _, ok := accessedTransactionIDs[transactionID]; !ok {
			return errors.Forbidden.New(fmt.Sprintf("You don't have access to transaction with ID %v", transactionID),
				errors.ParamsOption("AccountGroupIDs", accountGroupIDs, "TransactionID", transactionID),
				errors.SkipPreviousCallerOption(),
			)
		}
	}

	return nil
}
