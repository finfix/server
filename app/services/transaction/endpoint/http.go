package endpoint

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"server/app/pkg/http/chain"
	transactionService "server/app/services/transaction/service"
)

type endpoint struct {
	service *transactionService.Service
}

func NewEndpoint(service *transactionService.Service) http.Handler {

	s := &endpoint{
		service: service,
	}

	options := []chain.Option{
		chain.Before(chain.DefaultAuthorization),
	}

	router := chi.NewRouter()

	router.Method("POST", "/", chain.NewChain(s.createTransaction, options...))
	router.Method("PATCH", "/", chain.NewChain(s.updateTransaction, options...))
	router.Method("DELETE", "/", chain.NewChain(s.deleteTransaction, options...))
	router.Method("GET", "/", chain.NewChain(s.getTransactions, options...))
	return router
}
