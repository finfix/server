package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/server/middleware"
	"server/app/services/accountGroup/model"
)

// @Summary Создание группы счетов
// @Description Создается новая группа счетов
// @Tags accountGroup
// @Security AuthJWT
// @Accept json
// @Param Body body model.CreateAccountGroupReq true "model.CreateAccountGroupReq"
// @Produce json
// @Success 200 {object} model.CreateAccountGroupRes
// @Failure 400,401,403,500 {object} errors.CustomError
// @Router /accountGroup [post]
func (s *endpoint) createAccountGroup(ctx context.Context, r *http.Request) (any, error) {

	var req model.CreateAccountGroupReq

	// Декодируем запрос
	if err := middleware.DefaultDecoder(ctx, r, middleware.DecodeJSON, &req); err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return s.service.CreateAccountGroup(ctx, req)
}
