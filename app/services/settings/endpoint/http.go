package endpoint

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"server/app/config"
	"server/app/pkg/errors"
	"server/app/pkg/middleware"
	"server/app/pkg/server"
	settingsService "server/app/services/settings/service"
)

var part = "/settings"

type endpoint struct {
	service *settingsService.Service
}

func authorizationWithAdminKey(ctx context.Context, r *http.Request) (context.Context, error) {

	if r.Header.Get("AdminSecretKey") != config.GetConfig().AdminSecretKey {
		return ctx, errors.Forbidden.New("MySecretKey is incorrect", errors.Options{
			Params: map[string]any{
				"IP address": r.Header.Get("X-Real-IP"),
			},
		})
	}

	return ctx, nil
}

func NewEndpoint(service *settingsService.Service) http.Handler {

	e := &endpoint{
		service: service,
	}

	adminMethodsOptions := []server.Option{
		server.Before(authorizationWithAdminKey),
		server.ResponseEncoder(middleware.DefaultResponseEncoder),
		server.ErrorEncoder(middleware.DefaultErrorEncoder),
	}

	userMethodsOptions := []server.Option{
		server.Before(middleware.DefaultAuthorization),
		server.ResponseEncoder(middleware.DefaultResponseEncoder),
		server.ErrorEncoder(middleware.DefaultErrorEncoder),
	}

	r := mux.NewRouter()

	r.Methods("POST").Path(part + "/updateCurrencies").Handler(server.NewChain(e.updateCurrencies, adminMethodsOptions...))
	r.Methods("GET").Path(part + "/currencies").Handler(server.NewChain(e.getCurrencies, userMethodsOptions...))
	r.Methods("GET").Path(part + "/icons").Handler(server.NewChain(e.getIcons, userMethodsOptions...))
	r.Methods("GET").Path(part + "/version").Handler(server.NewChain(e.getVersion, []server.Option{
		server.ResponseEncoder(middleware.DefaultResponseEncoder),
		server.ErrorEncoder(middleware.DefaultErrorEncoder),
	}...))

	return r
}
