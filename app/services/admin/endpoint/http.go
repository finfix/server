package endpoint

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"server/app/config"
	"server/app/pkg/errors"
	"server/app/pkg/middleware"
	"server/app/pkg/server"
	adminService "server/app/services/admin/service"
)

var part = "/admin"

type endpoint struct {
	service *adminService.Service
}

func authorization(ctx context.Context, r *http.Request) (context.Context, error) {

	if r.Header.Get("AdminSecretKey") != config.GetConfig().AdminSecretKey {
		return nil, errors.Forbidden.New("MySecretKey is incorrect", errors.Options{
			Params: map[string]any{
				"IP address": r.Header.Get("X-Real-IP"),
			},
		})
	}

	return ctx, nil
}

func NewEndpoint(service *adminService.Service) http.Handler {

	e := &endpoint{
		service: service,
	}

	options := []server.Option{
		server.Before(authorization),
		server.ResponseEncoder(middleware.DefaultResponseEncoder),
		server.ErrorEncoder(middleware.DefaultErrorEncoder),
	}

	r := mux.NewRouter()

	r.Methods("POST").Path(part + "/updCurrencies").Handler(server.NewChain(e.updateCurrencies, options...))
	return r
}
