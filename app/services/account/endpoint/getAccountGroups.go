package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/server/middleware"
	"server/app/services/account/model"
)

// @Summary Получение списка групп счетов
// @Tags account
// @Security AuthJWT
// @Param Query query model.GetAccountGroupsReq true "model.GetAccountGroupsReq"
// @Produce json
// @Success 200 {object} []model.AccountGroup "[]model.AccountGroup"
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /account/accountGroups [get]
func (s *endpoint) getAccountGroups(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := middleware.DefaultDecoder(ctx, r, middleware.DecodeSchema, model.GetAccountGroupsReq{}) //nolint:exhaustruct
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return s.service.GetAccountGroups(ctx, req)
}
