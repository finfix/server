package endpoint

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"server/app/internal/config"
	adminService "server/app/internal/services/admin/service"
	"server/pkg/errors"
	"server/pkg/logging"
	"server/pkg/middleware"
	"server/pkg/server"
)

var part = "/admin"

type endpoint struct {
	service *adminService.Service
}

func authorization(ctx context.Context, r *http.Request) (context.Context, error) {

	if r.Header.Get("MySecretKey") != config.GetConfig().SecretKey {
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
		server.LoggingRequest(logging.DefaultRequestLoggerFunc),
		server.Before(authorization),
		server.ResponseEncoder(middleware.DefaultResponseEncoder),
		server.ErrorEncoder(middleware.DefaultErrorEncoder),
		server.ErrorLoggingFunc(logging.DefaultErrorLoggerFunc),
	}

	r := mux.NewRouter()

	r.Methods("POST").Path(part + "/updCurrencies").Handler(server.NewChain(e.updateCurrencies, options...))
	return r
}
