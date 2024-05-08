package endpoint

import (
	"net/http"

	"github.com/gorilla/mux"

	"server/app/pkg/server"
	"server/app/pkg/server/middleware"
	userService "server/app/services/user/service"
)

var part = "/user"

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

	r := mux.NewRouter()

	r.Methods("GET").Path(part + "/").Handler(server.NewChain(e.getUser, options...))
	return r
}
