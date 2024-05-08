package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/server/middleware"
	"server/app/services/account/model"
)

// @Summary Получение счетов по фильтрам
// @Tags account
// @Security AuthJWT
// @Param Query query model.GetAccountsReq false "model.GetAccountsReq"
// @Produce json
// @Success 200 {object} []model.Account
// @Failure 400,401,403,404,500 {object} errors.CustomError
// @Router /account [get]
func (s *endpoint) get(ctx context.Context, r *http.Request) (any, error) {

	// Декодируем запрос
	req, err := middleware.DefaultDecoder(ctx, r, middleware.DecodeSchema, model.GetAccountsReq{}) //nolint:exhaustruct
	if err != nil {
		return nil, err
	}

	// Вызываем метод сервиса
	return s.service.GetAccounts(ctx, req)
}
