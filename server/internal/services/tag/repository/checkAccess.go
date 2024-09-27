package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"pkg/errors"
)

// CheckAccess проверяет, имеет ли набор групп подкатегорий пользователя доступ к указанным идентификаторам подкатегорий
func (r *TagRepository) CheckAccess(ctx context.Context, accountGroupIDs, tagIDs []uint32) error {

	// Получаем все доступные счета по группам подкатегорий и перечисленным подкатегориям
	rows, err := r.db.Query(ctx, sq.
		Select("t.id").
		From("coin.tags t").
		Where(sq.Eq{
			"t.account_group_id": accountGroupIDs,
			"t.id":               tagIDs,
		}),
	)
	if err != nil {
		return err
	}

	// Формируем мапу доступных подкатегорий
	accessedTagIDs := make(map[uint32]struct{})

	// Проходимся по каждой доступной подкатегории
	for rows.Next() {

		// Считываем ID подкатегории
		var tagID uint32
		if err = rows.Scan(&tagID); err != nil {
			return err
		}

		// Добавляем ID подкатегории в мапу
		accessedTagIDs[tagID] = struct{}{}
	}

	if len(accessedTagIDs) == 0 {
		return errors.Forbidden.New("You don't have access to any of the requested tags",
			errors.ParamsOption("AccountGroupIDs", accountGroupIDs, "TagIDs", tagIDs),
		)
	}

	// Проходимся по каждому запрашиваемой подкатегории
	for _, tagID := range tagIDs {

		// Если подкатегории нет в мапе доступных подкатегорий, возвращаем ошибку
		if _, ok := accessedTagIDs[tagID]; !ok {
			return errors.Forbidden.New(fmt.Sprintf("You don't have access to tag with ID %v", tagID),
				errors.ParamsOption("AccountGroupIDs", accountGroupIDs, "TagID", tagID),
				errors.SkipPreviousCallerOption(),
			)
		}
	}

	return nil
}
