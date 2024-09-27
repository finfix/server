package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	tagRepoModel "server/internal/services/tag/repository/model"
)

// CreateTag создает новую подкатегорию
func (r *TagRepository) CreateTag(ctx context.Context, req tagRepoModel.CreateTagReq) (id uint32, err error) {

	// Создаем подкатегорию
	return r.db.ExecWithLastInsertID(ctx, sq.Insert(`coin.tags`).
		SetMap(map[string]any{
			"name":               req.Name,
			"account_group_id":   req.AccountGroupID,
			"created_by_user_id": req.CreatedByUserID,
			"datetime_create":    req.DatetimeCreate,
		}),
	)
}
