package endpoint

import (
	"context"
	"encoding/json"
	"net/http"

	"server/app/pkg/contextKeys"
	"server/app/pkg/errors"
	"server/app/pkg/validation"
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
	req, err := decodeSignUpReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return s.service.SignUp(ctx, req)
}

func decodeSignUpReq(ctx context.Context, r *http.Request) (req model.SignUpReq, err error) {

	// Декодируем тело запроса в структуру
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, errors.BadRequest.Wrap(err)
	}

	// Заполняем поля из контекста
	if deviceID := contextKeys.GetDeviceID(ctx); deviceID != nil {
		req.DeviceID = *deviceID
	} else {
		return req, errors.BadRequest.New("device id not found or not string")
	}

	// Валидируем поля
	if err = validation.Mail(req.Email); err != nil {
		return req, err
	}

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
