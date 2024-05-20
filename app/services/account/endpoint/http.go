package endpoint

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"server/app/pkg/server"
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

	options := []server.Option{
		server.Before(middleware.DefaultAuthorization),
	}

	r := chi.NewRouter()

	r.Method("POST", "/", server.NewChain(e.createAccount, options...))
	r.Method("GET", "/", server.NewChain(e.get, options...))
	r.Method("PATCH", "/", server.NewChain(e.updateAccount, options...))
	r.Method("DELETE", "/", server.NewChain(e.deleteAccount, options...))
	r.Method("GET", "/accountGroups", server.NewChain(e.getAccountGroups, options...))

	return r
}
