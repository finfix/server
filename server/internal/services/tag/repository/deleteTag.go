package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/errors"
)

// DeleteTag удаляет подкатегорию
func (r *TagRepository) DeleteTag(ctx context.Context, id, userID uint32) error {

	// Удаляем подкатегорию
	rows, err := r.db.ExecWithRowsAffected(ctx, sq.Delete("coin.tags").
		Where(sq.Eq{"id": id}),
	)
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
