package endpoint

import (
	"context"
	"net/http"

	_ "server/app/services/user/model" //nolint:golint
)

// @Summary Получение списка валют
// @Tags settings
// @Security AuthJWT
// @Produce json
// @Success 200 {object} []model.Currency
// @Failure 401,500 {object} errors.CustomError
// @Router /settings/currencies [get]
func (s *endpoint) getCurrencies(ctx context.Context, _ *http.Request) (any, error) {

	// Вызываем метод сервиса
	return s.service.GetCurrencies(ctx)
}
