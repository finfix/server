package endpoint

import (
	"context"
	"net/http"

	"server/app/pkg/server/middleware"
	"server/app/services/settings/model"
)

// @Summary Обновление курсов валют
// @Tags settings
// @Security AuthJWT
// @Success 200 "При успешном выполнении возвращается пустой ответ"
// @Failure 400,401,403,500 {object} errors.CustomError
// @Router /settings/updateCurrencies [post]
func (s *endpoint) updateCurrencies(ctx context.Context, r *http.Request) (any, error) {

	var req model.UpdateCurrenciesReq

	// Декодируем запрос
	if err := middleware.DefaultDecoder(ctx, r, middleware.DecodeJSON, &req); err != nil {
		return nil, err
	}

	return nil, s.service.UpdateCurrencies(ctx, req)
}
