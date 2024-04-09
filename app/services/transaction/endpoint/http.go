package endpoint

import (
	"net/http"

	"github.com/gorilla/mux"

	"server/app/pkg/logging"
	middleware2 "server/app/pkg/middleware"
	"server/app/pkg/server"
	transactionService "server/app/services/transaction/service"
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
		server.Before(middleware2.DefaultAuthorization),
		server.ResponseEncoder(middleware2.DefaultResponseEncoder),
		server.ErrorEncoder(middleware2.DefaultErrorEncoder),
		server.ErrorLoggingFunc(logging.DefaultErrorLoggerFunc),
	}

	router := mux.NewRouter()

	router.Methods("POST").Path(part).Handler(server.NewChain(s.createTransaction, options...))
	router.Methods("PATCH").Path(part).Handler(server.NewChain(s.updateTransaction, options...))
	router.Methods("DELETE").Path(part).Handler(server.NewChain(s.deleteTransaction, options...))
	router.Methods("GET").Path(part).Handler(server.NewChain(s.getTransactions, options...))
	return router
}
