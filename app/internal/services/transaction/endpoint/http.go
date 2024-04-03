package endpoint

import (
	"net/http"

	"github.com/gorilla/mux"

	transactionService "server/app/internal/services/transaction/service"
	"server/pkg/logging"
	"server/pkg/middleware"
	"server/pkg/server"
)

var part = "/transaction"

type endpoint struct {
	service *transactionService.Service
}

func NewEndpoint(service *transactionService.Service) http.Handler {

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

	r.Methods("POST").Path(part).Handler(server.NewChain(s.createTransaction, options...))
	r.Methods("PATCH").Path(part).Handler(server.NewChain(s.updateTransaction, options...))
	r.Methods("DELETE").Path(part).Handler(server.NewChain(s.deleteTransaction, options...))
	r.Methods("GET").Path(part).Handler(server.NewChain(s.getTransactions, options...))
	return r
}
