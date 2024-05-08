package endpoint

import (
	"net/http"

	"github.com/gorilla/mux"

	"server/app/pkg/server"
	"server/app/pkg/server/middleware"
	tagService "server/app/services/tag/service"
)

var part = "/tag"

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

	router := mux.NewRouter()

	router.Methods("POST").Path(part).Handler(server.NewChain(s.createTag, options...))
	router.Methods("PATCH").Path(part).Handler(server.NewChain(s.updateTag, options...))
	router.Methods("DELETE").Path(part).Handler(server.NewChain(s.deleteTag, options...))
	router.Methods("GET").Path(part).Handler(server.NewChain(s.getTags, options...))

	router.Methods("GET").Path(part + "/to_transactions").Handler(server.NewChain(s.getTagsToTransaction, options...))
	return router
}
