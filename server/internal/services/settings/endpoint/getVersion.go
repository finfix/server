package endpoint

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	_ "server/internal/services/settings/model" //nolint:golint
	"server/internal/services/settings/model/applicationType"
)

// @Summary Получение текущей версии сервера
// @Tags settings
// @Produce json
// @Success 200 {object} model.Version
// @Router /settings/version/ [get]
func (s *endpoint) getVersion(ctx context.Context, r *http.Request) (any, error) {
	appType := applicationType.Type(chi.URLParam(r, "application"))
	if err := appType.Validate(); err != nil {
		return nil, err
	}
	return s.service.GetVersion(ctx, appType)
}
