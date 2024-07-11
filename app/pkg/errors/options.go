package errors

import "fmt"

type Option func(*options)

type options struct {
	// Дополнительные данные для добавления контекста ошибки и дополнительных данных
	params map[string]string
	// Параметр, указывающий, сколько вызовов стека относительно текущего вызова вверх пропустить
	stackTrace *int
	// Тип логирования
	logAs *LogOption
	// Текст для пользователя
	HumanText string
	// Дополнительный текст к исходной ошибке
	errorfMessage *string
	// Параметр, указывающий, что тип ошибки затирать не надо при wrapping'е кастомной ошибки
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

func StackTraceOption(p int) Option {
	return func(o *options) { o.stackTrace = &p }
}

func LogAsOption(p LogOption) Option {
	return func(o *options) { o.logAs = &p }
}

func HumanTextOption(p string, args ...any) Option {
	humanText := fmt.Sprintf(p, args...)
	return func(o *options) { o.HumanText = humanText }
}

func ErrorfOption(p string) Option {
	return func(o *options) { o.errorfMessage = &p }
}

func DontEraseErrorType() Option {
	return func(o *options) { o.dontEraseErrorType = &struct{}{} }
}

func mergeOptions(opts ...Option) options {
	var options = &options{
		params:             nil,
		stackTrace:         nil,
		logAs:              nil,
		HumanText:          "",
		dontEraseErrorType: nil,
		errorfMessage:      nil,
	}

	for _, opt := range opts {
		opt(options)
	}

	return *options
}
