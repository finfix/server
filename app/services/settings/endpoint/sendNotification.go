package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/http/decoder"
	"server/app/services/settings/model"
)

// @Summary Отправка уведомления пользователю
// @Tags settings
// @Security AuthJWT
// @Param Body body model.SendNotificationReq true "model.SendNotificationReq"
// @Success 200 "При успешном выполнении возвращается пустой ответ"
// @Failure 400,401,403,500 {object} errors.Error
// @Router /settings/sendNotification [post]
func (s *endpoint) sendNotification(ctx context.Context, r *http.Request) (any, error) {

	var req model.SendNotificationReq

	// Декодируем запрос
	if err := decoder.Decoder(ctx, r, &req, decoder.DecodeJSON); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return s.service.SendNotification(ctx, req)
}
