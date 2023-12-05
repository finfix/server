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

// @Summary Регистрация пользователя
// @Tags auth
// @Accept jsonapi
// @Param Body body model.SignUpReq true "model.SignUpReq"
// @Param DeviceID header string true "Нужен для идентификации устройства"
// @Produce jsonapi
// @Success 200 {object} model.AuthRes
// @Failure 400,403,500 {object} errors.CustomError
// @Router /auth/signUp [post]
func (s *service) signUp(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := decodeSignUpReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	proto, err := s.client.SignUp(ctx, converter.SignUpReq{&req}.ConvertToProto())
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	// Конвертируем ответ во внутреннюю структуру
	return converter.PbAuthRes{proto}.ConvertToStruct(), nil
}

func decodeSignUpReq(ctx context.Context, r *http.Request) (req model.SignUpReq, err error) {

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
