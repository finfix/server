package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/server/middleware"
	"server/app/services/user/model"
)

// @Summary Редактирование пользователя
// @Description Изменение данных пользователя
// @Tags user
// @Security AuthJWT
// @Accept json
// @Param Body body model.UpdateUserReq true "model.UpdateUserReq"
// @Produce json
// @Success 200 "Если редактирование пользователя прошло успешно, возвращается пустой ответ"
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /user [patch]
func (s *endpoint) updateUser(ctx context.Context, r *http.Request) (any, error) {

	var req model.UpdateUserReq

	// Декодируем запрос
	if err := middleware.DefaultDecoder(ctx, r, middleware.DecodeJSON, &req); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return nil, s.service.UpdateUser(ctx, req)
}
