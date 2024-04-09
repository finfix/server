package endpoint

import (
	"net/http"

	"github.com/gorilla/mux"

	"server/app/pkg/logging"
	middleware2 "server/app/pkg/middleware"
	"server/app/pkg/server"
	authService "server/app/services/auth/service"
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
		server.Before(middleware2.DefaultDeviceIDValidator),
		server.ResponseEncoder(middleware2.DefaultResponseEncoder),
		server.ErrorEncoder(middleware2.DefaultErrorEncoder),
		server.ErrorLoggingFunc(logging.DefaultErrorLoggerFunc),
	}

	r := mux.NewRouter()

	r.Methods("POST").Path(part + "/signIn").Handler(server.NewChain(s.signIn, options...))
	r.Methods("POST").Path(part + "/signUp").Handler(server.NewChain(s.signUp, options...))
	r.Methods("POST").Path(part + "/refreshTokens").Handler(server.NewChain(s.refreshTokens, options...))
	return r
}
