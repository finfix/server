package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/errors"

	"server/internal/services/tag/model"
)

// GetTags возвращает все подкатегории по фильтрам
func (r *TagRepository) GetTags(ctx context.Context, req model.GetTagsReq) (tags []model.Tag, err error) {

	filtersEq := make(sq.Eq)

	if len(req.AccountGroupIDs) > 0 {
		filtersEq["account_group_id"] = req.AccountGroupIDs
	}

	// Проверяем, что есть фильтры
	if len(filtersEq) == 0 {
		return nil, errors.BadRequest.New("No filters specified")
	}

	// Получаем подкатегории
	return tags, r.db.Select(ctx, &tags, sq.
		Select("*").
		From("coin.tags").
		Where(filtersEq),
	)
}
