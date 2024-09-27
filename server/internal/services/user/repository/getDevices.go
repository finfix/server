package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	userModel "server/internal/services/user/model"
	userRepoModel "server/internal/services/user/repository/model"
)

// GetDevices Возвращает девайсы пользователя
func (r *UserRepository) GetDevices(ctx context.Context, filters userRepoModel.GetDevicesReq) (devices []userModel.Device, err error) {

	filtersEq := make(sq.Eq)

	if len(filters.IDs) > 0 {
		filtersEq["id"] = filters.IDs
	}
	if len(filters.DeviceIDs) > 0 {
		filtersEq["device_id"] = filters.DeviceIDs
	}
	if len(filters.UserIDs) > 0 {
		filtersEq["user_id"] = filters.UserIDs
	}

	// Получаем устройства пользователей
	return devices, r.db.Select(ctx, &devices, sq.
		Select("*").
		From("coin.devices").
		Where(filtersEq),
	)
}
