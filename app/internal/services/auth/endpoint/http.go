package endpoint

import (
	"net/http"

	"github.com/gorilla/mux"

	authService "server/app/internal/services/auth/service"
	"server/pkg/logging"
	"server/pkg/middleware"
	"server/pkg/server"
)

var part = "/auth"

type endpoint struct {
	service *authService.Service
}

func NewEndpoint(service *authService.Service) http.Handler {

	s := &endpoint{
		service: service,
	}

	options := []server.Option{
		server.LoggingRequest(logging.DefaultRequestLoggerFunc),
		server.Before(middleware.DefaultDeviceIDValidator),
		server.ResponseEncoder(middleware.DefaultResponseEncoder),
		server.ErrorEncoder(middleware.DefaultErrorEncoder),
		server.ErrorLoggingFunc(logging.DefaultErrorLoggerFunc),
	}

	r := mux.NewRouter()

	r.Methods("POST").Path(part + "/signIn").Handler(server.NewChain(s.signIn, options...))
	r.Methods("POST").Path(part + "/signUp").Handler(server.NewChain(s.signUp, options...))
	r.Methods("POST").Path(part + "/refreshTokens").Handler(server.NewChain(s.refreshTokens, options...))
	return r
}
