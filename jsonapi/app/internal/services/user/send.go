package user

import (
	"net/http"

	"logger/app/logging"
	"pkg/middleware"
	"pkg/server"

	"github.com/gorilla/mux"
	"core/app/proto/pbUser"
)

var part = "/user"

type service struct {
	client pbUser.UserClient
}

func NewService(client pbUser.UserClient) http.Handler {

	s := &service{
		client: client,
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
