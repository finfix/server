package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"runtime"
)

const (
	defaultPathDepth = 1
	SecondPathDepth  = 2
	ThirdPathDepth   = 3
)

type CustomError struct {
	ErrorType    `json:"-"`
	HumanText    string            `json:"humanTextError"`
	DevelopText  string            `json:"developerTextError"`
	InitialError error             `json:"-"`
	Path         string            `json:"path"`
	Params       map[string]string `json:"parameters,omitempty" validate:"required"`
	TaskID       *string           `json:"taskID,omitempty"`
	LogAs        LogOption         `json:"-"`
}

func (err CustomError) Error() string {
	return err.InitialError.Error()
}

type LogOption int

const (
	LogAsError LogOption = iota
	LogAsWarning
)

var TypeToLogOption = map[ErrorType]LogOption{
	LogicError:     LogAsError,
	BadRequest:     LogAsWarning,
	NotFound:       LogAsWarning,
	Teapot:         LogAsWarning,
	InternalServer: LogAsError,
	Forbidden:      LogAsWarning,
	Unauthorized:   LogAsWarning,
	ClientReject:   LogAsWarning,
	BadGateway:     LogAsWarning,
}

// New создает новую ошибку
func (typ ErrorType) New(msg string, opts ...Option) error {

	options := mergeOptions(opts...)

	skip := defaultPathDepth
	if options.pathDepth != nil {
		skip = *options.pathDepth
	}

	// Создаем новую ошибку
	customErr := CustomError{
		ErrorType:    typ,
		HumanText:    options.HumanText,
		DevelopText:  msg,
		InitialError: errors.New(msg),
		Path:         getPath(skip),
		Params:       options.params,
		TaskID:       nil,
		LogAs:        TypeToLogOption[typ],
	}

	// Если передан тип логирования, то добавляем его
	if options.logAs != nil {
		customErr.LogAs = *options.logAs
	}

	return customErr
}

// Wrap оборачивает ошибку
func (typ ErrorType) Wrap(err error, opts ...Option) error {

	options := mergeOptions(opts...)

	skip := defaultPathDepth

	var (
		customErr CustomError
		ok        bool
	)

	if customErr, ok = err.(CustomError); ok { // Если это уже обернутая ошибка

		// Если передан текст для пользователя, то затираем его
		if options.HumanText != "" {
			customErr.HumanText = options.HumanText
		}

		// Затираем путь, если передано в опциях
		if options.erasePath {
			// Если передана глубина пути, то используем ее
			if options.pathDepth != nil {
				skip = *options.pathDepth
			}
			customErr.Path = getPath(skip)
		}

		// Если передан текст ошибки, то добавляем его к исходной ошибке
		if options.errMessage != nil {
			customErr.InitialError = fmt.Errorf("%v: %v", options.errMessage, customErr.InitialError)
		}

		customErr.ErrorType = typ

	} else { // Если это не обернутая ошибка

		// Если передана глубина пути, то используем ее
		if options.pathDepth != nil {
			skip = *options.pathDepth
		}

		// Если это не обернутая ошибка, то создаем новую
		customErr = CustomError{
			ErrorType:    typ,
			HumanText:    options.HumanText,
			DevelopText:  err.Error(),
			InitialError: err,
			Path:         getPath(skip),
			Params:       options.params,
			TaskID:       nil,
			LogAs:        TypeToLogOption[typ],
		}

		if options.errMessage != nil {
			customErr.InitialError = fmt.Errorf("%v: %v", options.errMessage, err)
		} else {
			customErr.InitialError = err
		}
	}

	// Добавляем параметры
	maps.Copy(customErr.Params, options.params)

	if options.logAs != nil {
		customErr.LogAs = *options.logAs
	}

	return customErr
}

func CastError(err error) CustomError {

	customErr, ok := err.(CustomError)
	if !ok {
		err = InternalServer.Wrap(err, []Option{
			ErrMessageOption("Ошибка не обернута, путь неверный"),
			PathDepthOption(SecondPathDepth),
		}...)
		customErr, _ = err.(CustomError)
	}
	return customErr
}

func getPath(skip int) string {
	_, file, line, _ := runtime.Caller(skip + 1)
	return fmt.Sprintf("%v:%v", file, line)
}

// JSON преобразует ошибку в json
func JSON(err CustomError) ([]byte, error) {

	byt, e := json.Marshal(err)
	if e != nil {
		return nil, InternalServer.Wrap(e)
	}

	return byt, nil
}

func As(get error, target any) bool {
	if target != nil || get != nil {
		return errors.As(get, &target)
	}
	return false
}
