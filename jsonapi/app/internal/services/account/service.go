package account

import (
	"net/http"

	"logger/app/logging"
	"pkg/middleware"
	"pkg/server"

	"core/app/proto/pbAccount"

	"github.com/gorilla/mux"
)

var part = "/account"

type service struct {
	client pbAccount.AccountClient
}

func NewService(client pbAccount.AccountClient) http.Handler {

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

	r.Methods("POST").Path(part).Handler(server.NewChain(s.createAccount, options...))
	r.Methods("GET").Path(part).Handler(server.NewChain(s.get, options...))
	r.Methods("PATCH").Path(part).Handler(server.NewChain(s.updateAccount, options...))
	r.Methods("DELETE").Path(part).Handler(server.NewChain(s.deleteAccount, options...))
	r.Methods("PATCH").Path(part + "/switch").Handler(server.NewChain(s.switchAccounts, options...))
	r.Methods("GET").Path(part + "/accountGroups").Handler(server.NewChain(s.getAccountGroups, options...))

	return r
}
