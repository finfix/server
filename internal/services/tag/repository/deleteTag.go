package repository

import (
	"context"

	"server/pkg/errors"
)

// DeleteTag удаляет подкатегорию
func (repo *TagRepository) DeleteTag(ctx context.Context, id, userID uint32) error {

	// Удаляем подкатегорию
	rows, err := repo.db.ExecWithRowsAffected(ctx, ` DELETE FROM coin.tags WHERE id = ?`, id)
	if err != nil {
		return err
	}

	// Проверяем, что в базе данных что-то изменилось
	if rows == 0 {
		return errors.NotFound.New("No such model found for model",
			errors.ParamsOption("UserID", userID, "TagID", id),
		)
	}

	return nil
}
