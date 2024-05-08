package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/server/middleware"
	"server/app/services/account/model"
)

// @Summary Редактирование счета
// @Description Изменение данных счета
// @Tags account
// @Security AuthJWT
// @Accept json
// @Param Body body model.UpdateAccountReq true "model.UpdateAccountReq"
// @Produce json
// @Success 200 "Если редактирование счета прошло успешно, возвращается пустой ответ"
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /account [patch]
func (s *endpoint) updateAccount(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := middleware.DefaultDecoder(ctx, r, middleware.DecodeJSON, model.UpdateAccountReq{}) //nolint:exhaustruct
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return s.service.Update(ctx, req)
}
