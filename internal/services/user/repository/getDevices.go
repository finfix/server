package repository

import (
	"context"
	"fmt"
	"strings"

	userModel "server/internal/services/user/model"
	userRepoModel "server/internal/services/user/repository/model"
)

// GetDevices Возвращает девайсы пользователя
func (repo *Repository) GetDevices(ctx context.Context, filters userRepoModel.GetDevicesReq) (devices []userModel.Device, err error) {

	query := `
			SELECT *
			FROM coin.devices`

	var (
		queryArgs []string
		args      []any
	)

	if len(filters.IDs) > 0 {
		_query, _args, err := repo.db.In("id IN (?)", filters.IDs)
		if err != nil {
			return devices, err
		}
		queryArgs = append(queryArgs, _query)
		args = append(args, _args...)
	}

	if len(filters.DeviceIDs) > 0 {
		_query, _args, err := repo.db.In("device_id IN (?)", filters.DeviceIDs)
		if err != nil {
			return devices, err
		}
		queryArgs = append(queryArgs, _query)
		args = append(args, _args...)
	}

	if len(filters.UserIDs) > 0 {
		_query, _args, err := repo.db.In("user_id IN (?)", filters.UserIDs)
		if err != nil {
			return devices, err
		}
		queryArgs = append(queryArgs, _query)
		args = append(args, _args...)
	}

	if len(queryArgs) > 0 {
		query = fmt.Sprintf("%s WHERE %s", query, strings.Join(queryArgs, " AND "))
	}

	return devices, repo.db.Select(ctx, &devices, query, args...)
}
