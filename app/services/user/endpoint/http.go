package endpoint

import (
	"net/http"

	"github.com/gorilla/mux"

	"server/app/pkg/logging"
	middleware2 "server/app/pkg/middleware"
	"server/app/pkg/server"
	userService "server/app/services/user/service"
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
		server.Before(middleware2.DefaultAuthorization),
		server.ResponseEncoder(middleware2.DefaultResponseEncoder),
		server.ErrorEncoder(middleware2.DefaultErrorEncoder),
		server.ErrorLoggingFunc(logging.DefaultErrorLoggerFunc),
	}

	r := mux.NewRouter()

	r.Methods("GET").Path(part + "/currencies").Handler(server.NewChain(e.getCurrencies, options...))
	r.Methods("GET").Path(part + "/").Handler(server.NewChain(e.getUser, options...))
	return r
}
