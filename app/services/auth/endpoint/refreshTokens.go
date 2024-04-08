package endpoint

import (
	"context"
	"encoding/json"
	"net/http"

	"server/app/services/auth/model"
	"server/pkg/errors"
	"server/pkg/validation"
)

// @Summary Обновление токенов
// @Tags auth
// @Accept json
// @Param Dody body model.RefreshTokensReq true "model.RefreshTokensReq"
// @Produce json
// @Success 200 {object} model.RefreshTokensRes
// @Failure 400,401,500 {object} errors.CustomError
// @Router /auth/refreshTokens [post]
func (s *endpoint) refreshTokens(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := decodeRefreshTokensReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return s.service.RefreshTokens(ctx, req.Token)
}

func decodeRefreshTokensReq(_ context.Context, r *http.Request) (req model.RefreshTokensReq, err error) {

	// Декодируем тело запроса в структуру
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, errors.BadRequest.Wrap(err)
	}

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
