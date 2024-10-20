package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/ddlHelper"
	userModel "server/internal/services/user/model"
	"server/internal/services/user/repository/deviceDDL"
	userRepoModel "server/internal/services/user/repository/model"
)

// GetDevices Возвращает девайсы пользователя
func (r *UserRepository) GetDevices(ctx context.Context, filters userRepoModel.GetDevicesReq) (devices []userModel.Device, err error) {

	filtersEq := make(sq.Eq)

	if len(filters.IDs) > 0 {
		filtersEq[deviceDDL.ColumnID] = filters.IDs
	}
	if len(filters.DeviceIDs) > 0 {
		filtersEq[deviceDDL.ColumnDeviceID] = filters.DeviceIDs
	}
	if len(filters.UserIDs) > 0 {
		filtersEq[deviceDDL.ColumnUserID] = filters.UserIDs
	}

	// Получаем устройства пользователей
	return devices, r.db.Select(ctx, &devices, sq.
		Select(ddlHelper.SelectAll).
		From(deviceDDL.Table).
		Where(filtersEq),
	)
}
