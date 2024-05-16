package endpoint

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"server/app/pkg/server"
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

	options := []server.Option{
		server.Before(middleware.DefaultAuthorization),
	}

	r := chi.NewRouter()

	r.Method("GET", "/", server.NewChain(e.getUser, options...))
	return r
}
