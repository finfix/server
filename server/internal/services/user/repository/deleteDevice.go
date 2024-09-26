package repository

import "context"

// DeleteDevice Удаляет девайс пользователя
func (repo *UserRepository) DeleteDevice(ctx context.Context, userID uint32, deviceID string) error {
	return repo.db.Exec(ctx, `
			DELETE 
			FROM coin.devices 
			WHERE user_id = ? 
			  AND device_ID = ?`,
		userID,
		deviceID,
	)
}
