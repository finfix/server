package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"server/app/pkg/log"
	"server/app/pkg/panicRecover"
	"server/app/pkg/server/middleware"
)

type Chain struct {
	before      []BeforeFunc
	send        SendFunc
	after       []AfterFunc
	encode      EncodeResponseFunc
	errorEncode EncodeErrorFunc
}

func NewChain(send SendFunc, opts ...Option) *Chain {
	chain := &Chain{
		before:      nil,
		send:        send,
		after:       nil,
		encode:      middleware.DefaultResponseEncoder,
		errorEncode: middleware.DefaultErrorEncoder,
	}
	for _, option := range opts {
		option(chain)
	}
	return chain
}

type Option func(*Chain)

func ResponseEncoder(e EncodeResponseFunc) Option {
	return func(s *Chain) { s.encode = e }
}

func ErrorEncoder(ee EncodeErrorFunc) Option {
	return func(s *Chain) { s.errorEncode = ee }
}

func Before(before ...BeforeFunc) Option {
	return func(s *Chain) { s.before = append(s.before, before...) }
}

func After(after ...AfterFunc) Option {
	return func(s *Chain) { s.after = append(s.after, after...) }
}

type BeforeFunc func(context.Context, *http.Request) (context.Context, error)

type SendFunc func(context.Context, *http.Request) (any, error)

type AfterFunc func(context.Context, http.ResponseWriter) context.Context

type EncodeResponseFunc func(context.Context, http.ResponseWriter, any) error

type EncodeErrorFunc func(context.Context, http.ResponseWriter, error)

type LoggingFunc func(error)

func (s *Chain) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Обрабатываем панику, если она случилась
	defer panicRecover.PanicRecover(func(err error) {
		log.Error(context.Background(), err)
		middleware.DefaultErrorEncoder(context.Background(), w, err)
	})

	var (
		startTime = time.Now().Local()
		ctx       = r.Context()
		err       error
		res       any
	)

	ctx = log.SetTaskID(ctx)

	defer func() {
		log.Info(ctx, fmt.Sprintf("%v [%v]", r.URL.Path, strings.ToLower(r.Method)),
			log.ParamsOption(
				"method", r.Method,
				"path", r.URL.Path,
				"duration", time.Since(startTime),
			),
		)
	}()

	for _, f := range s.before {
		ctx, err = f(ctx, r)
		if err != nil {
			s.errorEncode(ctx, w, err)
			return
		}
	}

	res, err = s.send(ctx, r)
	if err != nil {
		s.errorEncode(ctx, w, err)
		return
	}

	if err = s.encode(ctx, w, res); err != nil {
		s.errorEncode(ctx, w, err)
		return
	}
}
