package repository

import (
	"context"

	userRepoModel "server/app/services/user/repository/model"
)

// CreateDevice Создает новый девайс для пользователя
func (repo *Repository) CreateDevice(ctx context.Context, req userRepoModel.CreateDeviceReq) (id uint32, err error) {
	return repo.db.ExecWithLastInsertID(ctx, `
			INSERT INTO coin.devices (
			  refresh_token, 
			  device_id, 
			  user_id,
			  os,
			  bundle_id
        	) VALUES (?, ?, ?, ?, ?)`,
		req.RefreshToken,
		req.DeviceID,
		req.UserID,
		req.OS,
		req.BundleID,
	)
}
