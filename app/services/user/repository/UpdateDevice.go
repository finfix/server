package repository

import (
	"context"
	"fmt"
	"strings"

	"server/app/pkg/errors"
	userRepoModel "server/app/services/user/repository/model"
)

// UpdateDevice редактирует девайс
func (repo *Repository) UpdateDevice(ctx context.Context, fields userRepoModel.UpdateDeviceReq) error {

	// Изменяем поля девайса
	var (
		args        []any
		queryFields []string
		query       string
	)

	// Добавляем в запрос поля, которые нужно изменить
	if fields.RefreshToken != nil {
		queryFields = append(queryFields, `refresh_token = ?`)
		args = append(args, fields.RefreshToken)
	}
	if fields.NotificationToken != nil {
		queryFields = append(queryFields, `notification_token = ?`)
		args = append(args, fields.NotificationToken)
	}

	if len(queryFields) == 0 {
		return errors.BadRequest.New("No fields to update")
	}

	// Конструируем запрос
	query = fmt.Sprintf(`
 			   UPDATE coin.devices 
               SET %v
			   WHERE user_id = ?
			     AND device_id = ?`,
		strings.Join(queryFields, ", "),
	)
	args = append(args, fields.UserID, fields.DeviceID)

	// Обновляем девайс
	return repo.db.Exec(ctx, query, args...)
}
