package transaction

import (
	"net/http"

	"github.com/gorilla/mux"

	"logger/app/logging"
	"pkg/middleware"
	"pkg/server"

	"core/app/proto/pbTransaction"
)

var part = "/transaction"

type service struct {
	client pbTransaction.TransactionClient
}

func NewService(client pbTransaction.TransactionClient) http.Handler {

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

	r.Methods("POST").Path(part).Handler(server.NewChain(s.createTransaction, options...))
	r.Methods("PATCH").Path(part).Handler(server.NewChain(s.updateTransaction, options...))
	r.Methods("DELETE").Path(part).Handler(server.NewChain(s.deleteTransaction, options...))
	r.Methods("GET").Path(part).Handler(server.NewChain(s.getTransactions, options...))
	return r
}
