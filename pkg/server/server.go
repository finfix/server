package server

import (
	"context"
	"net/http"
)

type Chain struct {
	loggingRequest LoggingRequest
	before         []BeforeFunc
	send           SendFunc
	after          []AfterFunc
	encode         EncodeResponseFunc
	errorEncode    EncodeErrorFunc
	loggingFunc    LoggingFunc
}

func NewChain(send SendFunc, opts ...ServerOption) *Chain {
	s := &Chain{
		send: send,
	}
	for _, option := range opts {
		option(s)
	}
	return s
}

type ServerOption func(*Chain)

func ServerResponseEncoder(e EncodeResponseFunc) ServerOption {
	return func(s *Chain) { s.encode = e }
}

func ServerErrorEncoder(ee EncodeErrorFunc) ServerOption {
	return func(s *Chain) { s.errorEncode = ee }
}

func ServerLoggingRequest(l LoggingRequest) ServerOption {
	return func(s *Chain) { s.loggingRequest = l }
}

func ServerBefore(before ...BeforeFunc) ServerOption {
	return func(s *Chain) { s.before = append(s.before, before...) }
}

func ServerAfter(after ...AfterFunc) ServerOption {
	return func(s *Chain) { s.after = append(s.after, after...) }
}

func ServerErrorLoggingFunc(loggingFunc func(error)) ServerOption {
	return func(s *Chain) { s.loggingFunc = loggingFunc }
}

type LoggingRequest func(*http.Request)
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
