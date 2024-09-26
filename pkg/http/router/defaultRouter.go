package router

import (
	"net/http/pprof"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"pkg/http/middleware"
)

func NewRouter() *chi.Mux {

	r := chi.NewRouter()

	// middlewares
	r.Use(
		middleware.PanicRecover,
		middleware.ResponseTime,
		middleware.RequestID,
		middleware.Logger,
	)

	// prometheus
	r.Handle("/metrics", promhttp.Handler())

	// healthCheck
	r.Get("/health", HealthCheck)

	// pprof
	r.HandleFunc("/debug/pprof/*", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	return r
}
