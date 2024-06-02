package endpoint

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"server/app/pkg/server"
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

	options := []server.Option{
		server.Before(middleware.DefaultAuthorization),
	}

	r := chi.NewRouter()

	r.Method("POST", "/updateCurrencies", server.NewChain(e.updateCurrencies, options...))
	r.Method("POST", "/sendNotification", server.NewChain(e.sendNotification, options...))
	r.Method("GET", "/currencies", server.NewChain(e.getCurrencies, options...))
	r.Method("GET", "/icons", server.NewChain(e.getIcons, options...))

	// Without authorization
	r.Method("GET", "/version/{application}", server.NewChain(e.getVersion))

	return r
}
