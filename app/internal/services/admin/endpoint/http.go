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
		return nil, errors.Forbidden.NewCtx("MySecretKey is incorrect", "IP address: %v", r.Header.Get("X-Real-IP"))
	}

	return ctx, nil
}

func NewEndpoint(service *adminService.Service) http.Handler {

	s := &endpoint{
		service: service,
	}

	options := []server.ServerOption{
		server.ServerLoggingRequest(logging.DefaultRequestLoggerFunc),
		server.ServerBefore(authorization),
		server.ServerResponseEncoder(middleware.DefaultResponseEncoder),
		server.ServerErrorEncoder(middleware.DefaultErrorEncoder),
		server.ServerErrorLoggingFunc(logging.DefaultErrorLoggerFunc),
	}

	r := mux.NewRouter()

	r.Methods("POST").Path(part + "/updCurrencies").Handler(server.NewChain(s.updateCurrencies, options...))
	return r
}
