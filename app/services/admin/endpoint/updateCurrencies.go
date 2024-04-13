package endpoint

import (
	"context"
	"net/http"
)

// @Summary Обновление курсов валют
// @Tags admin
// @Security SecretKey
// @Success 200 "При успешном выполнении возвращается пустой ответ"
// @Failure 400,401,403,500 {object} errors.CustomError
// @Router /admin/updateCurrencies [post]
func (s *endpoint) updateCurrencies(ctx context.Context, _ *http.Request) (any, error) {
	return nil, s.service.UpdateCurrencies(ctx)
}
