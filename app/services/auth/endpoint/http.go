package endpoint

import (
	"net/http"

	"github.com/gorilla/mux"

	"server/app/pkg/server"
	middleware2 "server/app/pkg/server/middleware"
	authService "server/app/services/auth/service"
)

var part = "/auth"

type endpoint struct {
	service *authService.Service
}

func NewEndpoint(service *authService.Service) http.Handler {

	s := &endpoint{
		service: service,
	}

	options := []server.Option{
		server.Before(middleware2.DefaultDeviceIDValidator),
	}

	r := mux.NewRouter()

	r.Methods("POST").Path(part + "/signIn").Handler(server.NewChain(s.signIn, options...))
	r.Methods("POST").Path(part + "/signUp").Handler(server.NewChain(s.signUp, options...))
	r.Methods("POST").Path(part + "/refreshTokens").Handler(server.NewChain(s.refreshTokens, []server.Option{
		server.Before(middleware2.DefaultDeviceIDValidator),
		server.Before(middleware2.ExtractDataFromToken),
	}...))
	return r
}
