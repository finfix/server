package endpoint

import (
	"context"
	"net/http"
)

func (s *endpoint) updateCurrencies(ctx context.Context, _ *http.Request) (any, error) {
	return nil, s.service.UpdateCurrencies(ctx)
}
