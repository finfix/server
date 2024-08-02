package endpoint

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"server/app/pkg/server/chain"
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

	options := []chain.Option{
		chain.Before(middleware.DefaultDeviceIDValidator),
	}

	r := chi.NewRouter()

	r.Method("POST", "/signIn", chain.NewChain(s.signIn, options...))
	r.Method("POST", "/signUp", chain.NewChain(s.signUp, options...))
	r.Method("POST", "/signOut", chain.NewChain(s.signOut, options...))
	r.Method("POST", "/refreshTokens", chain.NewChain(s.refreshTokens,
		chain.Before(middleware.DefaultDeviceIDValidator),
		chain.Before(middleware.ExtractDataFromToken),
	))
	return r
}
