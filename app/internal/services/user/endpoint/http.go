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

	s := &endpoint{
		service: service,
	}

	options := []server.ServerOption{
		server.ServerLoggingRequest(logging.DefaultRequestLoggerFunc),
		server.ServerBefore(middleware.DefaultAuthorization),
		server.ServerResponseEncoder(middleware.DefaultResponseEncoder),
		server.ServerErrorEncoder(middleware.DefaultErrorEncoder),
		server.ServerErrorLoggingFunc(logging.DefaultErrorLoggerFunc),
	}

	r := mux.NewRouter()

	r.Methods("GET").Path(part + "/currencies").Handler(server.NewChain(s.getCurrencies, options...))
	r.Methods("GET").Path(part + "/").Handler(server.NewChain(s.getUser, options...))
	return r
}
