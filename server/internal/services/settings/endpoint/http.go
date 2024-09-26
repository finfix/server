package endpoint

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"pkg/http/chain"

	"server/internal/services/settings/model"
	"server/internal/services/settings/model/applicationType"
)

type endpoint struct {
	service settingsService
}

type settingsService interface {
	UpdateCurrencies(context.Context, model.UpdateCurrenciesReq) error
	SendNotification(context.Context, model.SendNotificationReq) (model.SendNotificationRes, error)
	GetCurrencies(context.Context) ([]model.Currency, error)
	GetIcons(context.Context) ([]model.Icon, error)
	GetVersion(context.Context, applicationType.Type) (model.Version, error)
}

func MountSettingsEndpoints(mux *chi.Mux, service settingsService) {
	mux.Mount("/settings", newSettingsEndpoint(service))
}

func newSettingsEndpoint(service settingsService) http.Handler {

	e := &endpoint{
		service: service,
	}

	options := []chain.Option{
		chain.Before(chain.DefaultAuthorization),
	}

	r := chi.NewRouter()

	r.Method(http.MethodPost, "/updateCurrencies", chain.NewChain(e.updateCurrencies, options...))
	r.Method(http.MethodPost, "/sendNotification", chain.NewChain(e.sendNotification, options...))
	r.Method(http.MethodGet, "/currencies", chain.NewChain(e.getCurrencies, options...))
	r.Method(http.MethodGet, "/icons", chain.NewChain(e.getIcons, options...))

	// Without authorization
	r.Method(http.MethodGet, "/version/{application}", chain.NewChain(e.getVersion))

	return r
}
