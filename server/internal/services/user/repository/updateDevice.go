package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/errors"
	"server/internal/services/user/repository/deviceDDL"

	userRepoModel "server/internal/services/user/repository/model"
)

// UpdateDevice редактирует девайс
func (r *UserRepository) UpdateDevice(ctx context.Context, fields userRepoModel.UpdateDeviceReq) error {

	updates := make(map[string]any)

	// Добавляем в запрос поля, которые нужно изменить
	if fields.RefreshToken != nil {
		updates[deviceDDL.ColumnRefreshToken] = *fields.RefreshToken
	}
	if fields.NotificationToken != nil {
		updates[deviceDDL.ColumnNotificationToken] = *fields.NotificationToken
	}
	if fields.ApplicationInformation.BundleID != nil {
		updates[deviceDDL.ColumnApplicationBundleID] = *fields.ApplicationInformation.BundleID
	}
	if fields.ApplicationInformation.Build != nil {
		updates[deviceDDL.ColumnApplicationBuild] = *fields.ApplicationInformation.Build
	}
	if fields.ApplicationInformation.Version != nil {
		updates[deviceDDL.ColumnApplicationVersion] = *fields.ApplicationInformation.Version
	}
	if fields.DeviceInformation.VersionOS != nil {
		updates[deviceDDL.ColumnDeviceOSVersion] = *fields.DeviceInformation.VersionOS
	}
	if fields.DeviceInformation.IPAddress != nil {
		updates[deviceDDL.ColumnDeviceIPAddress] = *fields.DeviceInformation.IPAddress
	}
	if fields.DeviceInformation.UserAgent != nil {
		updates[deviceDDL.ColumnDeviceUserAgent] = *fields.DeviceInformation.UserAgent
	}

	if len(updates) == 0 {
		return errors.BadRequest.New("No fields to update")
	}

	// Обновляем девайс
	return r.db.Exec(ctx, sq.Update(deviceDDL.Table).
		SetMap(updates).
		Where(sq.Eq{
			deviceDDL.ColumnUserID:   fields.UserID,
			deviceDDL.ColumnDeviceID: fields.DeviceID,
		}),
	)
}
