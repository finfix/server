package endpoint

import (
	"net/http"
	accountService "server/app/internal/services/account/service"
	"server/pkg/logging"
	"server/pkg/middleware"
	"server/pkg/server"

	"github.com/gorilla/mux"
)

var part = "/account"

type endpoint struct {
	service *accountService.Service
}

func NewEndpoint(service *accountService.Service) http.Handler {

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

	r.Methods("POST").Path(part).Handler(server.NewChain(s.createAccount, options...))
	r.Methods("GET").Path(part).Handler(server.NewChain(s.get, options...))
	r.Methods("PATCH").Path(part).Handler(server.NewChain(s.updateAccount, options...))
	r.Methods("DELETE").Path(part).Handler(server.NewChain(s.deleteAccount, options...))
	r.Methods("PATCH").Path(part + "/switch").Handler(server.NewChain(s.switchAccounts, options...))
	r.Methods("GET").Path(part + "/accountGroups").Handler(server.NewChain(s.getAccountGroups, options...))

	return r
}
