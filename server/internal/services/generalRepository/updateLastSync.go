package generalRepository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
)

// UpdateLastSync Обновляет время последней синхронизации для устройства
func (repo *GeneralRepository) UpdateLastSync(ctx context.Context, deviceID string) error {
	return repo.db.Exec(ctx, sq.
		Update(`coin.devices`).
		Set(`last_sync`, time.Now()).
		Where(sq.Eq{"device_id": deviceID}),
	)
}
