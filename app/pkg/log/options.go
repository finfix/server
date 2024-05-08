package log

import "fmt"

type Option func(*options)

type options struct {
	// Дополнительные параметры
	params map[string]string
}

func ParamsOption(parameters ...any) Option {
	p := make(map[string]string)
	for i := 0; i < len(parameters); i += 2 {
		p[fmt.Sprintf("%v", parameters[i])] = fmt.Sprintf("%v", parameters[i+1])
	}
	if len(parameters)%2 != 0 {
		p[fmt.Sprintf("%v", parameters[len(parameters)-1])] = "param not found"
	}
	return func(o *options) { o.params = p }
}

func mergeOptions(opts ...Option) options {
	options := &options{
		params: nil,
	}

	for _, opt := range opts {
		opt(options)
	}

	return *options
}
