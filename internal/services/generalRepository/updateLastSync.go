package generalRepository

import (
	"context"
	"time"
)

// UpdateLastSync Обновляет время последней синхронизации для устройства
func (repo *Repository) UpdateLastSync(ctx context.Context, deviceID string) error {
	return repo.db.Exec(ctx, `
			UPDATE coin.devices 
			SET last_sync = ? 
			WHERE device_id = ?`,
		time.Now(),
		deviceID,
	)
}
