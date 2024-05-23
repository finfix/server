package endpoint

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"server/app/pkg/server"
	"server/app/pkg/server/middleware"
	accountGroupService "server/app/services/accountGroup/service"
)

type endpoint struct {
	service *accountGroupService.Service
}

func NewEndpoint(service *accountGroupService.Service) http.Handler {

	e := &endpoint{
		service: service,
	}

	options := []server.Option{
		server.Before(middleware.DefaultAuthorization),
	}

	r := chi.NewRouter()

	r.Method("POST", "/", server.NewChain(e.createAccountGroup, options...))
	r.Method("PATCH", "/", server.NewChain(e.updateAccountGroup, options...))
	r.Method("DELETE", "/", server.NewChain(e.deleteAccountGroup, options...))
	r.Method("GET", "/", server.NewChain(e.getAccountGroups, options...))

	return r
}
