package endpoint

import (
	"net/http"

	"github.com/gorilla/mux"

	transactionService "server/app/services/transaction/service"
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

	options := []server.Option{
		server.LoggingRequest(logging.DefaultRequestLoggerFunc),
		server.Before(middleware.DefaultAuthorization),
		server.ResponseEncoder(middleware.DefaultResponseEncoder),
		server.ErrorEncoder(middleware.DefaultErrorEncoder),
		server.ErrorLoggingFunc(logging.DefaultErrorLoggerFunc),
	}

	router := mux.NewRouter()

	router.Methods("POST").Path(part).Handler(server.NewChain(s.createTransaction, options...))
	router.Methods("PATCH").Path(part).Handler(server.NewChain(s.updateTransaction, options...))
	router.Methods("DELETE").Path(part).Handler(server.NewChain(s.deleteTransaction, options...))
	router.Methods("GET").Path(part).Handler(server.NewChain(s.getTransactions, options...))
	return router
}
