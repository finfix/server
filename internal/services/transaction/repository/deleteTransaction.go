package repository

import (
	"context"

	"server/pkg/errors"
)

// DeleteTransaction удаляет транзакцию
func (repo *TransactionRepository) DeleteTransaction(ctx context.Context, id, userID uint32) error {

	// Удаляем транзакцию
	rows, err := repo.db.ExecWithRowsAffected(ctx, `
			   DELETE FROM coin.transactions 
			   WHERE id = ?`,
		id,
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
