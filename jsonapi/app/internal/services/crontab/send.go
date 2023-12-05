package crontab

import (
	"logger/app/logging"
	"pkg/errors"
	"pkg/middleware"
	"pkg/server"

	pb "core/app/proto/pbCrontab"
	"jsonapi/app/internal/config"

	"context"
	"net/http"

	"github.com/gorilla/mux"
)

var part = "/crontab"

type service struct {
	client pb.CrontabClient
}

func authorization(ctx context.Context, r *http.Request) (context.Context, error) {

	if r.Header.Get("MySecretKey") != config.GetConfig().SecretKey {
		return nil, errors.Forbidden.NewCtx("MySecretKey is incorrect", "IP address: %v", r.Header.Get("X-Real-IP"))
	}

	return ctx, nil
}

func NewService(client pb.CrontabClient) http.Handler {

	s := &service{
		client: client,
	}

	options := []server.ServerOption{
		server.ServerLoggingRequest(logging.DefaultRequestLoggerFunc),
		server.ServerBefore(authorization),
		server.ServerResponseEncoder(middleware.DefaultResponseEncoder),
		server.ServerErrorEncoder(middleware.DefaultErrorEncoder),
		server.ServerErrorLoggingFunc(logging.DefaultErrorLoggerFunc),
	}

	r := mux.NewRouter()

	r.Methods("POST").Path(part + "/updCurrencies").Handler(server.NewChain(s.updateCurrencies, options...))
	return r
}
