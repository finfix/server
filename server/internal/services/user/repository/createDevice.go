package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	userModel "server/internal/services/user/model"
)

// CreateDevice Создает новый девайс для пользователя
func (r *UserRepository) CreateDevice(ctx context.Context, req userModel.Device) (id uint32, err error) {
	return r.db.ExecWithLastInsertID(ctx, sq.
		Insert(`coin.devices`).
		SetMap(map[string]any{
			"refresh_token":         req.RefreshToken,
			"device_id":             req.DeviceID,
			"user_id":               req.UserID,
			"device_os_name":        req.DeviceInformation.NameOS,
			"device_os_version":     req.DeviceInformation.VersionOS,
			"device_name":           req.DeviceInformation.DeviceName,
			"device_model_name":     req.DeviceInformation.ModelName,
			"device_ip_address":     req.DeviceInformation.IPAddress,
			"device_user_agent":     req.DeviceInformation.UserAgent,
			"application_bundle_id": req.ApplicationInformation.BundleID,
			"application_version":   req.ApplicationInformation.Version,
			"application_build":     req.ApplicationInformation.Build,
		}),
	)
}
