package endpoint

import (
	"net/http"

	"github.com/gorilla/mux"

	"server/app/pkg/logging"
	"server/app/pkg/middleware"
	"server/app/pkg/server"
	tagService "server/app/services/tag/service"
)

var part = "/tag"

type endpoint struct {
	service *tagService.Service
}

func NewEndpoint(logger *logging.Logger, service *tagService.Service) http.Handler {

	s := &endpoint{
		service: service,
	}

	options := []server.Option{
		server.Before(middleware.DefaultAuthorization),
		server.ResponseEncoder(middleware.DefaultResponseEncoder),
		server.ErrorEncoder(middleware.DefaultErrorEncoder),
	}

	router := mux.NewRouter()

	router.Methods("POST").Path(part).Handler(server.NewChain(logger, s.createTag, options...))
	router.Methods("PATCH").Path(part).Handler(server.NewChain(logger, s.updateTag, options...))
	router.Methods("DELETE").Path(part).Handler(server.NewChain(logger, s.deleteTag, options...))
	router.Methods("GET").Path(part).Handler(server.NewChain(logger, s.getTags, options...))

	router.Methods("GET").Path(part + "/to_transactions").Handler(server.NewChain(logger, s.getTagsToTransaction, options...))
	return router
}
