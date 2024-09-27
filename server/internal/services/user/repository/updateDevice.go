package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/errors"

	userRepoModel "server/internal/services/user/repository/model"
)

// UpdateDevice редактирует девайс
func (r *UserRepository) UpdateDevice(ctx context.Context, fields userRepoModel.UpdateDeviceReq) error {

	updates := make(map[string]any)

	// Добавляем в запрос поля, которые нужно изменить
	if fields.RefreshToken != nil {
		updates["refresh_token"] = *fields.RefreshToken
	}
	if fields.NotificationToken != nil {
		updates["notification_token"] = *fields.NotificationToken
	}
	if fields.ApplicationInformation.BundleID != nil {
		updates["application_bundle_id"] = *fields.ApplicationInformation.BundleID
	}
	if fields.ApplicationInformation.Build != nil {
		updates["application_build"] = *fields.ApplicationInformation.Build
	}
	if fields.ApplicationInformation.Version != nil {
		updates["application_version"] = *fields.ApplicationInformation.Version
	}
	if fields.DeviceInformation.VersionOS != nil {
		updates["device_os_version"] = *fields.DeviceInformation.VersionOS
	}
	if fields.DeviceInformation.IPAddress != nil {
		updates["device_ip_address"] = *fields.DeviceInformation.IPAddress
	}
	if fields.DeviceInformation.UserAgent != nil {
		updates["device_user_agent"] = *fields.DeviceInformation.UserAgent
	}

	if len(updates) == 0 {
		return errors.BadRequest.New("No fields to update")
	}

	// Обновляем девайс
	return r.db.Exec(ctx, sq.Update("coin.devices").
		SetMap(updates).
		Where(sq.Eq{
			"user_id":   fields.UserID,
			"device_id": fields.DeviceID,
		}),
	)
}
