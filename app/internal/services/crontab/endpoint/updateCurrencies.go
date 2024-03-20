package endpoint

import (
	"context"
	"net/http"
)

func (s *endpoint) updateCurrencies(ctx context.Context, _ *http.Request) (any, error) {

	// Вызываем метод сервиса
	return s.service.UpdateCurrencies(ctx)
}
