package log

import (
	"server/app/pkg/pointer"
	"server/app/pkg/stackTrace"
)

type Option func(*options)

type options struct {
	// Дополнительные параметры
	params map[string]string
	// Насколько надо скипнуть текущий вызов в стеке вызовов
	stackTraceSkip *int
}

func ParamsOption(keyVal ...string) Option {
	p := make(map[string]string)
	for i := 0; i < len(keyVal); i += 2 {
		p[keyVal[i]] = keyVal[i+1]
	}
	if len(keyVal)%2 != 0 {
		p[keyVal[len(keyVal)-1]] = "param not found"
	}
	return func(o *options) { o.params = p }
}

func SkipThisCallOption() Option {
	return func(o *options) { o.stackTraceSkip = pointer.Pointer(stackTrace.SkipThisCall) }
}

func SkipPreviousCallerOption() Option {
	return func(o *options) { o.stackTraceSkip = pointer.Pointer(stackTrace.SkipPreviousCaller) }
}

func Skip2PreviousCallersOption() Option {
	return func(o *options) { o.stackTraceSkip = pointer.Pointer(stackTrace.Skip2PreviousCallers) }
}

func Skip3PreviousCallersOption() Option {
	return func(o *options) { o.stackTraceSkip = pointer.Pointer(stackTrace.Skip3PreviousCallers) }
}

func mergeOptions(opts ...Option) options {
	options := &options{
		params: make(map[string]string),
	}

	for _, opt := range opts {
		opt(options)
	}

	return *options
}
