package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/server/middleware"
	"server/app/services/auth/model"
)

// @Summary Обновление токенов
// @Tags auth
// @Accept json
// @Security AuthJWT
// @Param Dody body model.RefreshTokensReq true "model.RefreshTokensReq"
// @Produce json
// @Success 200 {object} model.RefreshTokensRes
// @Failure 400,401,500 {object} errors.CustomError
// @Router /auth/refreshTokens [post]
func (s *endpoint) refreshTokens(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := middleware.DefaultDecoder(ctx, r, middleware.DecodeJSON, model.RefreshTokensReq{}) //nolint:exhaustruct
	if err != nil {
		return nil, err
	}

	req.Device.IPAddress = r.Header.Get("X-Real-IP")
	req.Device.UserAgent = r.Header.Get("User-Agent")

	// Вызываем метод сервиса
	return s.service.RefreshTokens(ctx, req)
}
