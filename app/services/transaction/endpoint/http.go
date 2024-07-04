package endpoint

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"server/app/pkg/server"
	"server/app/pkg/server/middleware"
	transactionService "server/app/services/transaction/service"
)

type endpoint struct {
	service *transactionService.Service
}

func NewEndpoint(service *transactionService.Service) http.Handler {

	s := &endpoint{
		service: service,
	}

	options := []server.Option{
		server.Before(middleware.DefaultAuthorization),
	}

	router := chi.NewRouter()

	router.Method("POST", "/", server.NewChain(s.createTransaction, options...))
	router.Method("PATCH", "/", server.NewChain(s.updateTransaction, options...))
	router.Method("DELETE", "/", server.NewChain(s.deleteTransaction, options...))
	router.Method("GET", "/", server.NewChain(s.getTransactions, options...))

	router.Method("POST", "/file", server.NewChain(s.createFile, options...))
	return router
}
