package log

type Option func(*options)

type options struct {
	// Дополнительные параметры
	params map[string]string
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

func mergeOptions(opts ...Option) options {
	options := &options{
		params: nil,
	}

	for _, opt := range opts {
		opt(options)
	}

	return *options
}
