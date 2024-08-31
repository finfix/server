package generalRepository

import (
	"context"
	"time"

	"server/internal/services/action/model/enum"
)

// CreateAction создает новый лог действия пользователя
func (repo *Repository) CreateAction(ctx context.Context, actionType enum.ActionType, deviceID string, userID, objectID uint32) error {

	// Добавляем лог
	if err := repo.db.Exec(ctx, `
			INSERT INTO coin.action_history(
			  action_type_signatura, 
		      user_id, 
		      object_id, 
		      action_time
		    ) VALUES (?, ?, ?, ?)`,
		actionType,
		userID,
		objectID,
		time.Now(),
	); err != nil {
		return err
	}

	// Обновляем время последней синхронизации
	return repo.UpdateLastSync(ctx, deviceID)
}
