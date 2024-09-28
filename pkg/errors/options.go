package errors

import (
	"fmt"

	"pkg/pointer"
	"pkg/stackTrace"
)

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
	// Дополнительная ошибка для errors.Is к исходной ошибке
	errorf *error
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

func SkipThisCallOption() Option {
	return func(o *options) { o.stackTrace = pointer.Pointer(stackTrace.SkipThisCall) }
}

func SkipPreviousCallerOption() Option {
	return func(o *options) { o.stackTrace = pointer.Pointer(stackTrace.SkipPreviousCaller) }
}

func Skip2PreviousCallersOption() Option {
	return func(o *options) { o.stackTrace = pointer.Pointer(stackTrace.Skip2PreviousCallers) }
}

func LogAsOption(p LogOption) Option {
	return func(o *options) { o.logAs = &p }
}

func HumanTextOption(p string, args ...any) Option {
	humanText := fmt.Sprintf(p, args...)
	return func(o *options) { o.HumanText = humanText }
}

func ErrorfOption(err error) Option {
	return func(o *options) { o.errorf = &err }
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
		errorf:             nil,
	}

	for _, opt := range opts {
		opt(options)
	}

	return *options
}
