package endpoint

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"server/app/pkg/server/chain"
	"server/app/pkg/server/middleware"
	accountService "server/app/services/account/service"
)

type endpoint struct {
	service *accountService.Service
}

func NewEndpoint(service *accountService.Service) http.Handler {

	e := &endpoint{
		service: service,
	}

	options := []chain.Option{
		chain.Before(middleware.DefaultAuthorization),
	}

	r := chi.NewRouter()

	r.Method("POST", "/", chain.NewChain(e.createAccount, options...))
	r.Method("GET", "/", chain.NewChain(e.get, options...))
	r.Method("PATCH", "/", chain.NewChain(e.updateAccount, options...))
	r.Method("DELETE", "/", chain.NewChain(e.deleteAccount, options...))

	return r
}
