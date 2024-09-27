package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

// DeleteDevice Удаляет девайс пользователя
func (r *UserRepository) DeleteDevice(ctx context.Context, userID uint32, deviceID string) error {
	return r.db.Exec(ctx, sq.
		Delete(`coin.devices`).
		From(`coin.devices`).
		Where(sq.Eq{
			"user_id":   userID,
			"device_ID": deviceID,
		}),
	)
}
