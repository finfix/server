package endpoint

import (
	"context"
	"net/http"

	"pkg/contextKeys"
	"pkg/errors"
	"pkg/http/decoder"

	"server/internal/services/auth/model"
)

// @Summary Авторизация пользователя по логину и паролю
// @Tags auth
// @Accept json
// @Param Body body model.SignInReq true "model.SignInReq"
// @Param DeviceID header string true "Нужен для идентификации устройства"
// @Produce json
// @Success 200 {object} model.AuthRes
// @Failure 400,404,500 {object} errors.Error
// @Router /auth/signIn [post]
func (s *endpoint) signIn(ctx context.Context, r *http.Request) (any, error) {

	var req model.SignInReq

	deviceID := contextKeys.GetDeviceID(ctx)
	if deviceID == nil {
		return nil, errors.BadRequest.New("DeviceID не задан")
	}
	req.DeviceID = *deviceID

	// Декодируем запрос
	if err := decoder.Decoder(ctx, r, &req, decoder.DecodeJSON); err != nil {
		return nil, err
	}

	req.Device.IPAddress = r.Header.Get("X-Real-IP")
	req.Device.UserAgent = r.Header.Get("User-Agent")

	// Вызываем метод сервиса
	return s.service.SignIn(ctx, req)
}
