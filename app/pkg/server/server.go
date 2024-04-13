package server

import (
	"context"
	"net/http"
	"time"

	"server/app/pkg/logging"
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
		send: send,
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

type ErrorHandler interface {
	Handle(ctx context.Context, err error)
}

func (s *Chain) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	startTime := time.Now().Local()
	defer func() {
		logging.GetLogger().Info("Request %v %v %v", r.Method, r.URL.Path, time.Since(startTime))
	}()

	var (
		ctx = r.Context()
		err error
		res any
	)

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