package auth

import (
	pb "auth/app/proto/pbAuth"
	"context"
	"encoding/json"
	"net/http"
	"pkg/converter"

	"jsonapi/app/internal/services/auth/model"
	"pkg/errors"
	"pkg/validation"
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
func (s *service) signIn(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := decodeSignInReq(ctx, r)
	if err != nil {
		return nil, err
	}

	in, err := converter.Convert(pb.SignInReq{}, req)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	out, err := s.client.SignIn(ctx, &in)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	res, err := converter.Convert(model.AuthRes{}, out)
	if err != nil {
		return nil, err
	}

	// Конвертируем ответ во внутреннюю структуру
	return res, nil
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
