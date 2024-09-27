package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/errors"
)

// DeleteTransaction удаляет транзакцию
func (r *TransactionRepository) DeleteTransaction(ctx context.Context, id, userID uint32) error {

	// Удаляем транзакцию
	rows, err := r.db.ExecWithRowsAffected(ctx, sq.
		Delete("coin.transactions").
		Where(sq.Eq{"id": id}),
	)
	if err != nil {
		return err
	}

	// Проверяем, что в базе данных что-то изменилось
	if rows == 0 {
		return errors.NotFound.New("No such model found for model",
			errors.ParamsOption("UserID", userID, "TransactionID", id),
		)
	}

	return nil
}
