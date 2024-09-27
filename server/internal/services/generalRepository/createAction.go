package generalRepository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"

	"server/internal/services/action/model/enum"
)

// CreateAction создает новый лог действия пользователя
func (repo *GeneralRepository) CreateAction(ctx context.Context, actionType enum.ActionType, deviceID string, userID, objectID uint32) error {

	// Добавляем лог
	if err := repo.db.Exec(ctx, sq.
		Insert(`coin.action_history`).
		SetMap(map[string]any{
			"action_type_signatura": actionType,
			"user_id":               userID,
			"object_id":             objectID,
			"action_time":           time.Now(),
		}),
	); err != nil {
		return err
	}

	// Обновляем время последней синхронизации
	return repo.UpdateLastSync(ctx, deviceID)
}
