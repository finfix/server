package endpoint

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"server/app/pkg/server"
	"server/app/pkg/server/middleware"
	tagService "server/app/services/tag/service"
)

type endpoint struct {
	service *tagService.Service
}

func NewEndpoint(service *tagService.Service) http.Handler {

	s := &endpoint{
		service: service,
	}

	options := []server.Option{
		server.Before(middleware.DefaultAuthorization),
	}

	router := chi.NewRouter()

	router.Method("POST", "/", server.NewChain(s.createTag, options...))
	router.Method("PATCH", "/", server.NewChain(s.updateTag, options...))
	router.Method("DELETE", "/", server.NewChain(s.deleteTag, options...))
	router.Method("GET", "/", server.NewChain(s.getTags, options...))

	router.Method("GET", "/to_transactions", server.NewChain(s.getTagsToTransaction, options...))
	return router
}
