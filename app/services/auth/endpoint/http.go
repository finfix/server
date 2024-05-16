package endpoint

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"server/app/pkg/server"
	"server/app/pkg/server/middleware"
	authService "server/app/services/auth/service"
)

type endpoint struct {
	service *authService.Service
}

func NewEndpoint(service *authService.Service) http.Handler {

	s := &endpoint{
		service: service,
	}

	options := []server.Option{
		server.Before(middleware.DefaultDeviceIDValidator),
	}

	r := chi.NewRouter()

	r.Method("POST", "/signIn", server.NewChain(s.signIn, options...))
	r.Method("POST", "/signUp", server.NewChain(s.signUp, options...))
	r.Method("POST", "/refreshTokens", server.NewChain(s.refreshTokens, []server.Option{
		server.Before(middleware.DefaultDeviceIDValidator),
		server.Before(middleware.ExtractDataFromToken),
	}...))
	return r
}
