package endpoint

import (
	"context"
	"net/http"

	"pkg/errors"
	"pkg/http/decoder"

	"server/internal/services/user/model"
)

// @Summary Получение данных пользователя
// @Tags user
// @Security AuthJWT
// @Produce json
// @Success 200 {object} model.User
// @Failure 401,404,500 {object} errors.Error
// @Router /user [get]
func (s *endpoint) getUser(ctx context.Context, r *http.Request) (any, error) {

	var req model.GetUsersReq

	// Декодируем запрос
	if err := decoder.Decoder(ctx, r, &req, decoder.DecodeSchema); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	users, err := s.service.GetUsers(ctx, req)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	if len(users) == 0 {
		return nil, errors.InternalServer.New("Пользователь не найден",
			errors.ParamsOption("UserID", req.Necessary.UserID),
		)
	}

	// Конвертируем ответ во внутреннюю структуру
	return users[0], nil
}
