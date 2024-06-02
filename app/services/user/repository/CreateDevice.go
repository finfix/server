package repository

import (
	"context"

	userModel "server/app/services/user/model"
)

// CreateDevice Создает новый девайс для пользователя
func (repo *Repository) CreateDevice(ctx context.Context, req userModel.Device) (id uint32, err error) {
	return repo.db.ExecWithLastInsertID(ctx, `
			INSERT INTO coin.devices (
			  refresh_token, 
			  device_id, 
			  user_id,
			  device_os_name,
			  device_os_version,
			  device_name,
			  device_model_name,
			  device_ip_address,
			  device_user_agent,
			  application_bundle_id,
			  application_version,
			  application_build
        	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		req.RefreshToken,
		req.DeviceID,
		req.UserID,
		req.DeviceInformation.NameOS,
		req.DeviceInformation.VersionOS,
		req.DeviceInformation.DeviceName,
		req.DeviceInformation.ModelName,
		req.DeviceInformation.IPAddress,
		req.DeviceInformation.UserAgent,
		req.ApplicationInformation.BundleID,
		req.ApplicationInformation.Version,
		req.ApplicationInformation.Build,
	)
}
