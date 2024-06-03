package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/server/middleware"
	"server/app/services/auth/model"
)

// @Summary Регистрация пользователя
// @Tags auth
// @Accept json
// @Param Body body model.SignUpReq true "model.SignUpReq"
// @Param DeviceID header string true "Нужен для идентификации устройства"
// @Produce json
// @Success 200 {object} model.AuthRes
// @Failure 400,403,500 {object} errors.CustomError
// @Router /auth/signUp [post]
func (s *endpoint) signUp(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := middleware.DefaultDecoder(ctx, r, middleware.DecodeJSON, model.SignUpReq{}) //nolint:exhaustruct
	if err != nil {
		return nil, err
	}

	req.Device.IPAddress = r.Header.Get("X-Real-IP")
	req.Device.UserAgent = r.Header.Get("User-Agent")

	// Вызываем метод сервиса
	return s.service.SignUp(ctx, req)
}
