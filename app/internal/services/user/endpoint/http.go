package endpoint

import (
	"net/http"

	"github.com/gorilla/mux"

	userService "server/app/internal/services/user/service"
	"server/pkg/logging"
	"server/pkg/middleware"
	"server/pkg/server"
)

var part = "/user"

type endpoint struct {
	service *userService.Service
}

func NewEndpoint(service *userService.Service) http.Handler {

	e := &endpoint{
		service: service,
	}

	options := []server.Option{
		server.LoggingRequest(logging.DefaultRequestLoggerFunc),
		server.Before(middleware.DefaultAuthorization),
		server.ResponseEncoder(middleware.DefaultResponseEncoder),
		server.ErrorEncoder(middleware.DefaultErrorEncoder),
		server.ErrorLoggingFunc(logging.DefaultErrorLoggerFunc),
	}

	r := mux.NewRouter()

	r.Methods("GET").Path(part + "/currencies").Handler(server.NewChain(e.getCurrencies, options...))
	r.Methods("GET").Path(part + "/").Handler(server.NewChain(e.getUser, options...))
	return r
}
