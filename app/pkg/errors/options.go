package errors

import "fmt"

type Option func(*options)

type options struct {
	// Дополнительные данные для ошибки
	params map[string]string
	// Параметры
	pathDepth *int
	// Тип логирования
	logAs *LogOption
	// Текст для пользователя
	HumanText string
	// Дополнительный текст к исходной ошибке
	errMessage *string
	// Не затирать тип ошибки
	dontEraseErrorType *struct{}
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

func PathDepthOption(p int) Option {
	return func(o *options) { o.pathDepth = &p }
}

func LogAsOption(p LogOption) Option {
	return func(o *options) { o.logAs = &p }
}

func HumanTextOption(p string, args ...any) Option {
	humanText := fmt.Sprintf(p, args...)
	return func(o *options) { o.HumanText = humanText }
}

func ErrMessageOption(p string) Option {
	return func(o *options) { o.errMessage = &p }
}

func DontEraseErrorType() Option {
	return func(o *options) { o.dontEraseErrorType = &struct{}{} }
}

func mergeOptions(opts ...Option) options {
	options := &options{
		params:             nil,
		pathDepth:          nil,
		logAs:              nil,
		HumanText:          "",
		dontEraseErrorType: nil,
		errMessage:         nil,
	}

	for _, opt := range opts {
		opt(options)
	}

	return *options
}
