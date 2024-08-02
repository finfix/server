package endpoint

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"server/app/pkg/server/chain"
	"server/app/pkg/server/middleware"
	userService "server/app/services/user/service"
)

type endpoint struct {
	service *userService.Service
}

func NewEndpoint(service *userService.Service) http.Handler {

	e := &endpoint{
		service: service,
	}

	options := []chain.Option{
		chain.Before(middleware.DefaultAuthorization),
	}

	r := chi.NewRouter()

	r.Method("GET", "/", chain.NewChain(e.getUser, options...))
	r.Method("PATCH", "/", chain.NewChain(e.updateUser, options...))
	return r
}
