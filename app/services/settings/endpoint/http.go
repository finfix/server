package endpoint

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"server/app/pkg/server/chain"
	"server/app/pkg/server/middleware"
	settingsService "server/app/services/settings/service"
)

type endpoint struct {
	service *settingsService.Service
}

func NewEndpoint(service *settingsService.Service) http.Handler {

	e := &endpoint{
		service: service,
	}

	options := []chain.Option{
		chain.Before(middleware.DefaultAuthorization),
	}

	r := chi.NewRouter()

	r.Method("POST", "/updateCurrencies", chain.NewChain(e.updateCurrencies, options...))
	r.Method("POST", "/sendNotification", chain.NewChain(e.sendNotification, options...))
	r.Method("GET", "/currencies", chain.NewChain(e.getCurrencies, options...))
	r.Method("GET", "/icons", chain.NewChain(e.getIcons, options...))

	// Without authorization
	r.Method("GET", "/version/{application}", chain.NewChain(e.getVersion))

	return r
}
