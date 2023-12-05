package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"jsonapi/app/internal/services/auth/converter"
	"jsonapi/app/internal/services/auth/model"
	"pkg/errors"
	"pkg/validation"
)

// @Summary Обновление токенов
// @Tags auth
// @Accept jsonapi
// @Param Dody body pbAuth.RefreshTokensReq true "pbAuth.RefreshTokensReq"
// @Produce jsonapi
// @Success 200 {object} model.RefreshTokensRes
// @Failure 400,401,500 {object} errors.CustomError
// @Router /auth/refreshTokens [post]
func (s *service) refreshTokens(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := decodeRefreshTokensReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	proto, err := s.client.RefreshTokens(ctx, converter.RefreshTokensReq{&req}.ConvertToProto())
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	// Конвертируем ответ во внутреннюю структуру
	return converter.PbRefreshTokensRes{proto}.ConvertToStruct(), nil
}

func decodeRefreshTokensReq(_ context.Context, r *http.Request) (req model.RefreshTokensReq, err error) {

	// Декодируем тело запроса в структуру
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, errors.BadRequest.Wrap(err)
	}

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
