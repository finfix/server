package endpoint

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"server/app/pkg/server/chain"
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

	options := []chain.Option{
		chain.Before(middleware.DefaultAuthorization),
	}

	r := chi.NewRouter()

	r.Method("POST", "/", chain.NewChain(e.createAccountGroup, options...))
	r.Method("PATCH", "/", chain.NewChain(e.updateAccountGroup, options...))
	r.Method("DELETE", "/", chain.NewChain(e.deleteAccountGroup, options...))
	r.Method("GET", "/", chain.NewChain(e.getAccountGroups, options...))

	return r
}
