package endpoint

import (
	"net/http"

	"github.com/gorilla/mux"

	accountService "server/app/services/account/service"
	"server/pkg/logging"
	"server/pkg/middleware"
	"server/pkg/server"
)

var part = "/account"

type endpoint struct {
	service *accountService.Service
}

func NewEndpoint(service *accountService.Service) http.Handler {

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

	router := mux.NewRouter()

	router.Methods("POST").Path(part).Handler(server.NewChain(e.createAccount, options...))
	router.Methods("GET").Path(part).Handler(server.NewChain(e.get, options...))
	router.Methods("PATCH").Path(part).Handler(server.NewChain(e.updateAccount, options...))
	router.Methods("DELETE").Path(part).Handler(server.NewChain(e.deleteAccount, options...))
	router.Methods("PATCH").Path(part + "/switch").Handler(server.NewChain(e.switchAccounts, options...))
	router.Methods("GET").Path(part + "/accountGroups").Handler(server.NewChain(e.getAccountGroups, options...))

	return router
}
