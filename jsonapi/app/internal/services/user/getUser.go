package user

import (
	"context"
	pb "core/app/proto/pbUser"
	"jsonapi/app/internal/services/user/model"
	"net/http"
	"pkg/converter"
	"pkg/errors"
	"pkg/validation"
)

// @Summary Получение данных пользователя
// @Tags user
// @Security AuthJWT
// @Produce json
// @Param Authorization header string true "Бла бла бла"
// @Success 200 {object} model.User
// @Failure 401,404,500 {object} errors.CustomError
// @Router /user/ [get]
func (s *service) getUser(ctx context.Context, r *http.Request) (any, error) {

	req, err := decodeGetUserReq(ctx, r)
	if err != nil {
		return nil, err
	}

	in, err := converter.Convert(pb.GetReq{}, req)
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	out, err := s.client.Get(ctx, &in)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	users, err := converter.Convert([]model.User{}, out)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.NotFound.NewCtx("Пользователь не найден", "UserID: %v", req.ID)
	}

	// Конвертируем ответ во внутреннюю структуру
	return users[0], nil
}

func decodeGetUserReq(ctx context.Context, _ *http.Request) (req model.GetReq, err error) {

	// Заполняем поля из контекста
	req.ID, _ = ctx.Value("UserID").(uint32)

	// Проверяем обязательные поля на zero value
	return req, validation.ZeroValue(req)
}
