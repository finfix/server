package endpoint

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"server/app/pkg/http/chain"
	tagService "server/app/services/tag/service"
)

type endpoint struct {
	service *tagService.Service
}

func NewEndpoint(service *tagService.Service) http.Handler {

	s := &endpoint{
		service: service,
	}

	options := []chain.Option{
		chain.Before(chain.DefaultAuthorization),
	}

	router := chi.NewRouter()

	router.Method("POST", "/", chain.NewChain(s.createTag, options...))
	router.Method("PATCH", "/", chain.NewChain(s.updateTag, options...))
	router.Method("DELETE", "/", chain.NewChain(s.deleteTag, options...))
	router.Method("GET", "/", chain.NewChain(s.getTags, options...))

	router.Method("GET", "/to_transactions", chain.NewChain(s.getTagsToTransaction, options...))
	return router
}
