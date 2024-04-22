package endpoint

import (
	"context"
	"net/http"

	_ "server/app/services/settings/model" //nolint:golint
)

// @Summary Получение текущей версии сервера
// @Tags settings
// @Produce json
// @Success 200 {object} model.Version
// @Router /settings/version [get]
func (s *endpoint) getVersion(_ context.Context, _ *http.Request) (any, error) {
	return s.service.GetVersion(), nil
}
