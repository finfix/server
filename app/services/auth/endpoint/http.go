package endpoint

import (
	"net/http"

	"github.com/gorilla/mux"

	"server/app/pkg/middleware"
	"server/app/pkg/server"
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
		server.Before(middleware.DefaultDeviceIDValidator),
		server.ResponseEncoder(middleware.DefaultResponseEncoder),
		server.ErrorEncoder(middleware.DefaultErrorEncoder),
	}

	r := mux.NewRouter()

	r.Methods("POST").Path(part + "/signIn").Handler(server.NewChain(s.signIn, options...))
	r.Methods("POST").Path(part + "/signUp").Handler(server.NewChain(s.signUp, options...))
	r.Methods("POST").Path(part + "/refreshTokens").Handler(server.NewChain(s.refreshTokens, []server.Option{
		server.Before(middleware.DefaultDeviceIDValidator),
		server.Before(middleware.ExtractDataFromToken),
		server.ResponseEncoder(middleware.DefaultResponseEncoder),
		server.ErrorEncoder(middleware.DefaultErrorEncoder),
	}...))
	return r
}
