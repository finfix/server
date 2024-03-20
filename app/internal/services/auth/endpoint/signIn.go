package endpoint

import (
	"context"
	"encoding/json"
	"net/http"

	"server/app/internal/services/auth/model"
	"server/pkg/errors"
	"server/pkg/validation"
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
	req, err := decodeSignInReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return s.service.SignIn(ctx, req)
}

func decodeSignInReq(ctx context.Context, r *http.Request) (req model.SignInReq, err error) {

	// Декодируем тело запроса в структуру
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, errors.BadRequest.Wrap(err)
	}

	// Заполняем поля из контекста
	req.DeviceID, _ = ctx.Value("DeviceID").(string)

	// Валидируем поля
	if err = validation.Mail(req.Email); err != nil {
		return req, err
	}

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
