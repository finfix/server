package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/server/middleware"
	"server/app/services/auth/model"
)

// @Summary Авторизация пользователя по логину и паролю
// @Tags auth
// @Accept json
// @Param Body body model.SignInReq true "model.SignInReq"
// @Param DeviceID header string true "Нужен для идентификации устройства"
// @Produce json
// @Success 200 {object} model.AuthRes
// @Failure 400,404,500 {object} errors.CustomError
// @Router /auth/signIn [post]
func (s *endpoint) signIn(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := middleware.DefaultDecoder(ctx, r, middleware.DecodeJSON, model.SignInReq{}) //nolint:exhaustruct
	if err != nil {
		return nil, err
	}

	req.Device.IPAddress = r.Header.Get("X-Real-IP")
	req.Device.UserAgent = r.Header.Get("User-Agent")

	// Вызываем метод сервиса
	return s.service.SignIn(ctx, req)
}
