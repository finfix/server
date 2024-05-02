package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/errors"
	"server/app/pkg/validation"
	"server/app/services"
	userModel "server/app/services/user/model"
)

// @Summary Получение данных пользователя
// @Tags user
// @Security AuthJWT
// @Produce json
// @Success 200 {object} model.User
// @Failure 401,404,500 {object} errors.CustomError
// @Router /user/ [get]
func (s *endpoint) getUser(ctx context.Context, r *http.Request) (any, error) {

	req, err := decodeGetUserReq(ctx, r)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	users, err := s.service.GetUsers(ctx, req)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	if len(users) == 0 {
		return nil, errors.InternalServer.New("Пользователь не найден", errors.Options{Params: map[string]any{
			"UserID": req.Necessary.UserID,
		}})
	}

	// Конвертируем ответ во внутреннюю структуру
	return users[0], nil
}

func decodeGetUserReq(ctx context.Context, _ *http.Request) (req userModel.GetReq, err error) {

	// Заполняем поля из контекста
	req.Necessary, err = services.ExtractNecessaryFromCtx(ctx)
	if err != nil {
		return req, err
	}

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
