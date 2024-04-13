package endpoint

import (
	"net/http"

	"github.com/gorilla/mux"

	"server/app/pkg/middleware"
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
		server.Before(middleware.DefaultAuthorization),
		server.ResponseEncoder(middleware.DefaultResponseEncoder),
		server.ErrorEncoder(middleware.DefaultErrorEncoder),
	}

	router := mux.NewRouter()

	router.Methods("POST").Path(part).Handler(server.NewChain(s.createTransaction, options...))
	router.Methods("PATCH").Path(part).Handler(server.NewChain(s.updateTransaction, options...))
	router.Methods("DELETE").Path(part).Handler(server.NewChain(s.deleteTransaction, options...))
	router.Methods("GET").Path(part).Handler(server.NewChain(s.getTransactions, options...))
	return router
}