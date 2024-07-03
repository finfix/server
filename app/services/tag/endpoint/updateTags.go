package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/server/middleware"
	"server/app/services/tag/model"
)

// @Summary Редактирование транзакции
// @Description Изменение данных транзакции
// @Tags tag
// @Security AuthJWT
// @Accept json
// @Param Body body model.UpdateTagReq true "model.UpdateTagReq"
// @Success 200 "При успешном выполнении запроса возвращает пустой ответ"
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /tag [patch]
func (s *endpoint) updateTag(ctx context.Context, r *http.Request) (any, error) {

	var req model.UpdateTagReq

	// Декодируем запрос
	if err := middleware.DefaultDecoder(ctx, r, middleware.DecodeJSON, &req); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return nil, s.service.UpdateTag(ctx, req)
}
