package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/server/middleware"
	"server/app/services/accountGroup/model"
)

// @Summary Редактирование группы счетов
// @Tags accountGroup
// @Security AuthJWT
// @Accept json
// @Param Body body model.UpdateAccountGroupReq true "model.UpdateAccountGroupReq"
// @Produce json
// @Success 200 "Если редактирование группы счетов прошло успешно, возвращается пустой ответ"
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /accountGroup [patch]
func (s *endpoint) updateAccountGroup(ctx context.Context, r *http.Request) (any, error) {

	var req model.UpdateAccountGroupReq

	// Декодируем запрос
	if err := middleware.DefaultDecoder(ctx, r, middleware.DecodeJSON, &req); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return nil, s.service.Update(ctx, req)
}