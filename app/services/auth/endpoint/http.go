package endpoint

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"server/app/pkg/errors"
	"server/app/pkg/http/chain"
	"server/app/pkg/jwtManager"
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
		chain.Before(chain.DefaultDeviceIDValidator),
	}

	r := chi.NewRouter()

	r.Method("POST", "/signIn", chain.NewChain(s.signIn, options...))
	r.Method("POST", "/signUp", chain.NewChain(s.signUp, options...))
	r.Method("POST", "/signOut", chain.NewChain(s.signOut, options...))
	r.Method("POST", "/refreshTokens", chain.NewChain(s.refreshTokens, append(options, chain.Before(extractDataFromToken))...))
	return r
}

func extractDataFromToken(ctx context.Context, r *http.Request) (context.Context, error) {

	// Проводим авторизацию
	ctx, err := chain.DefaultAuthorization(ctx, r)
	if err != nil {

		// Если ошибка истекшего токена, то это ок, так как мы смогли его распарсить и получить оттуда данные
		if !errors.Is(err, jwtManager.ErrUserUnauthorized) {
			return ctx, err
		}
	}

	return ctx, nil
}
