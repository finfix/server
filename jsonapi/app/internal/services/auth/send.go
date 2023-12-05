package auth

import (
	"logger/app/logging"
	"pkg/middleware"
	"pkg/server"

	"auth/app/proto/pbAuth"

	"net/http"

	"github.com/gorilla/mux"
)

var part = "/auth"

type service struct {
	client pbAuth.AuthClient
}

func NewService(client pbAuth.AuthClient) http.Handler {

	s := &service{
		client: client,
	}

	options := []server.ServerOption{
		server.ServerLoggingRequest(logging.DefaultRequestLoggerFunc),
		server.ServerBefore(middleware.DefaultDeviceIDValidator),
		server.ServerResponseEncoder(middleware.DefaultResponseEncoder),
		server.ServerErrorEncoder(middleware.DefaultErrorEncoder),
		server.ServerErrorLoggingFunc(logging.DefaultErrorLoggerFunc),
	}

	r := mux.NewRouter()

	r.Methods("POST").Path(part + "/signIn").Handler(server.NewChain(s.signIn, options...))
	r.Methods("POST").Path(part + "/signUp").Handler(server.NewChain(s.signUp, options...))
	r.Methods("POST").Path(part + "/refreshTokens").Handler(server.NewChain(s.refreshTokens, options...))
	return r
}
