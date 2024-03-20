package endpoint

import (
	"server/pkg/logging"
	"server/pkg/middleware"
	"server/pkg/server"

	"net/http"

	authService "server/app/internal/services/auth/service"

	"github.com/gorilla/mux"
)

var part = "/auth"

type endpoint struct {
	service *authService.Service
	logger  *logging.Logger
}

func NewEndpoint(service *authService.Service) http.Handler {

	s := &endpoint{
		service: service,
	}

	options := []server.ServerOption{
		server.ServerLoggingRequest(logging.DefaultRequestLoggerFunc),
		server.ServerBefore(middleware.DefaultDeviceIDValidator),
		server.ServerResponseEncoder(middleware.DefaultResponseEncoder),
		server.ServerErrorEncoder(middleware.DefaultErrorEncoder),
		server.ServerErrorLoggingFunc(logging.DefaultErrorLoggerFunc),
	}

	r := mux.NewRouter()

	r.Methods("POST").Path(part + "/signIn").Handler(server.NewChain(s.signIn, options...))
	r.Methods("POST").Path(part + "/signUp").Handler(server.NewChain(s.signUp, options...))
	r.Methods("POST").Path(part + "/refreshTokens").Handler(server.NewChain(s.refreshTokens, options...))
	return r
}
