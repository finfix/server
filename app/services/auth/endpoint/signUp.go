package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/contextKeys"
	"server/app/pkg/errors"
	"server/app/pkg/http/decoder"
	"server/app/services/auth/model"
)

// @Summary Регистрация пользователя
// @Tags auth
// @Accept json
// @Param Body body model.SignUpReq true "model.SignUpReq"
// @Param DeviceID header string true "Нужен для идентификации устройства"
// @Produce json
// @Success 200 {object} model.AuthRes
// @Failure 400,403,500 {object} errors.Error
// @Router /auth/signUp [post]
func (s *endpoint) signUp(ctx context.Context, r *http.Request) (any, error) {

	var req model.SignUpReq

	deviceID := contextKeys.GetDeviceID(ctx)
	if deviceID != nil {
		req.DeviceID = *deviceID
	} else {
		return nil, errors.BadRequest.New("Не передан DeviceID в заголовке запроса")
	}

	// Декодируем запрос
	if err := decoder.Decoder(ctx, r, &req, decoder.DecodeJSON); err != nil {
		return nil, err
	}

	req.Device.IPAddress = r.Header.Get("X-Real-IP")
	req.Device.UserAgent = r.Header.Get("User-Agent")

	// Вызываем метод сервиса
	return s.service.SignUp(ctx, req)
}
