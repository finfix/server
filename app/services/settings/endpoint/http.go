package endpoint

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"server/app/pkg/errors"
	"server/app/pkg/server"
	"server/app/pkg/server/middleware"
	settingsService "server/app/services/settings/service"
)

type endpoint struct {
	service  *settingsService.Service
	adminKey string
}

func authorizationWithAdminKey(adminKey string) func(ctx context.Context, r *http.Request) (context.Context, error) {
	return func(ctx context.Context, r *http.Request) (context.Context, error) {
		if r.Header.Get("AdminSecretKey") != adminKey {
			return ctx, errors.Forbidden.New("MySecretKey is incorrect", []errors.Option{
				errors.ParamsOption("IP address", r.Header.Get("X-Real-IP")),
			}...)
		}
		return ctx, nil
	}
}

func NewEndpoint(
	service *settingsService.Service,
	adminKey string,
) http.Handler {

	e := &endpoint{
		service:  service,
		adminKey: adminKey,
	}

	adminMethodsOptions := []server.Option{
		server.Before(authorizationWithAdminKey(adminKey)),
	}

	userMethodsOptions := []server.Option{
		server.Before(middleware.DefaultAuthorization),
	}

	r := chi.NewRouter()

	r.Method("POST", "/updateCurrencies", server.NewChain(e.updateCurrencies, adminMethodsOptions...))
	r.Method("GET", "/currencies", server.NewChain(e.getCurrencies, userMethodsOptions...))
	r.Method("GET", "/icons", server.NewChain(e.getIcons, userMethodsOptions...))
	r.Method("GET", "/version/{application}", server.NewChain(e.getVersion))

	return r
}
