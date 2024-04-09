package server

import (
	"context"
	"net/http"
)

type Chain struct {
	loggingRequest LoggingRequestFunc
	before         []BeforeFunc
	send           SendFunc
	after          []AfterFunc
	encode         EncodeResponseFunc
	errorEncode    EncodeErrorFunc
	loggingFunc    LoggingFunc
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

func LoggingRequest(l LoggingRequestFunc) Option {
	return func(s *Chain) { s.loggingRequest = l }
}

func Before(before ...BeforeFunc) Option {
	return func(s *Chain) { s.before = append(s.before, before...) }
}

func After(after ...AfterFunc) Option {
	return func(s *Chain) { s.after = append(s.after, after...) }
}

func ErrorLoggingFunc(loggingFunc func(error)) Option {
	return func(s *Chain) { s.loggingFunc = loggingFunc }
}

type LoggingRequestFunc func(*http.Request)
type BeforeFunc func(context.Context, *http.Request) (context.Context, error)
type SendFunc func(context.Context, *http.Request) (any, error)
type AfterFunc func(context.Context, http.ResponseWriter) context.Context
type EncodeResponseFunc func(context.Context, http.ResponseWriter, any) error
type EncodeErrorFunc func(context.Context, http.ResponseWriter, error, func(error))
type LoggingFunc func(error)

type ErrorHandler interface {
	Handle(ctx context.Context, err error)
}

func (s *Chain) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var (
		ctx = r.Context()
		err error
		res any
	)

	if s.loggingRequest != nil {
		s.loggingRequest(r)
	}

	for _, f := range s.before {
		ctx, err = f(ctx, r)
		if err != nil {
			s.errorEncode(ctx, w, err, s.loggingFunc)
			return
		}
	}

	res, err = s.send(ctx, r)
	if err != nil {
		s.errorEncode(ctx, w, err, s.loggingFunc)
		return
	}

	if err = s.encode(ctx, w, res); err != nil {
		s.errorEncode(ctx, w, err, s.loggingFunc)
		return
	}
}
