package repository

import (
	"context"

	tagRepoModel "server/internal/services/tag/repository/model"
)

// CreateTag создает новую подкатегорию
func (repo *TagRepository) CreateTag(ctx context.Context, req tagRepoModel.CreateTagReq) (id uint32, err error) {

	// Создаем подкатегорию
	if id, err = repo.db.ExecWithLastInsertID(ctx, `
			INSERT INTO coin.tags (
    		  name,
			  account_group_id,
			  created_by_user_id,
			  datetime_create
            ) VALUES (?, ?, ?, ?)`,
		req.Name,
		req.AccountGroupID,
		req.CreatedByUserID,
		req.DatetimeCreate,
	); err != nil {
		return id, err
	}
	return id, nil
}
