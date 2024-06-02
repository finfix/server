package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/server/middleware"
	settingsModel "server/app/services/settings/model"
)

// @Summary Отправка уведомления пользователю
// @Tags settings
// @Security AuthJWT
// @Param Body body settingsModel.SendNotificationReq true "settingsModel.SendNotificationReq"
// @Success 200 "При успешном выполнении возвращается пустой ответ"
// @Failure 400,401,403,500 {object} errors.CustomError
// @Router /settings/sendNotification [post]
func (s *endpoint) sendNotification(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := middleware.DefaultDecoder(ctx, r, middleware.DecodeJSON, settingsModel.SendNotificationReq{}) //nolint:exhaustruct
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return s.service.SendNotification(ctx, req)
}
