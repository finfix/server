package repository

import (
	"context"
	defErr "errors"
	"fmt"
	"strings"

	"pkg/errors"

	"server/internal/services/tag/model"
)

// GetTags возвращает все подкатегории по фильтрам
func (repo *TagRepository) GetTags(ctx context.Context, req model.GetTagsReq) (tags []model.Tag, err error) {

	var (
		args        []any
		queryFields []string
	)

	_query, _args, err := repo.db.In(`account_group_id IN (?)`, req.AccountGroupIDs)
	if err != nil {
		return nil, err
	}
	queryFields = append(queryFields, _query)
	args = append(args, _args...)

	// Конструируем запрос
	query := fmt.Sprintf(`SELECT *
		   FROM coin.tags 
		   WHERE %v`,
		strings.Join(queryFields, " AND "),
	)

	// Получаем подкатегории
	if err = repo.db.Select(ctx, &tags, query, args...); err != nil {
		if defErr.Is(err, context.Canceled) {
			return nil, errors.ClientReject.New("HTTP connection terminated")
		}
		return nil, err
	}

	return tags, nil
}
