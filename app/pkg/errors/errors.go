package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"net/http"
	"runtime"
)

const (
	LogicError     = ErrorType(1)
	BadRequest     = ErrorType(http.StatusBadRequest)
	NotFound       = ErrorType(http.StatusNotFound)
	Teapot         = ErrorType(http.StatusTeapot)
	InternalServer = ErrorType(http.StatusInternalServerError)
	Forbidden      = ErrorType(http.StatusForbidden)
	Unauthorized   = ErrorType(http.StatusUnauthorized)
	ClientReject   = ErrorType(http.StatusNotAcceptable)
	BadGateway     = ErrorType(http.StatusBadGateway)
)

const (
	defaultPathDepth = 1
	SecondPathDepth  = 2
	ThirdPathDepth   = 3
)

type ErrorType uint32

type CustomError struct {
	ErrorType    `json:"-"`
	HumanText    string         `json:"humanTextError"`
	DevelopText  string         `json:"developerTextError"`
	InitialError error          `json:"-"`
	Path         string         `json:"path"`
	Params       map[string]any `json:"parameters,omitempty" validate:"required"`
	TaskID       *string        `json:"taskID,omitempty"`
	LogAs        LogOption      `json:"-"`
}

func (err CustomError) Error() string {
	return err.InitialError.Error()
}

type LogOption int

const (
	LogAsError LogOption = iota
	LogAsWarning
)

type Options struct {
	// Дополнительные данные для ошибки
	Params map[string]any
	// Параметры
	PathDepth int
	// Тип логирования
	LogAs *LogOption
	// Текст для пользователя
	HumanText string

	// Затирать путь
	ErasePath bool
	// Дополнительный текст к исходной ошибке
	ErrMessage string
}

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
func (typ ErrorType) New(msg string, opts ...Options) error {

	var options Options
	if len(opts) > 0 {
		options = opts[0]
	}

	skip := defaultPathDepth

	// Если передана глубина пути, то используем ее
	if options.PathDepth != 0 {
		skip = options.PathDepth
	}

	// Создаем новую ошибку
	customErr := CustomError{
		ErrorType:    typ,
		InitialError: errors.New(msg),
		Path:         getPath(skip),
		Params:       options.Params,
		HumanText:    options.HumanText,
		LogAs:        TypeToLogOption[typ],
	}

	// Если передан тип логирования, то добавляем его
	if options.LogAs != nil {
		customErr.LogAs = *options.LogAs
	}

	return customErr
}

// Wrap оборачивает ошибку
func (typ ErrorType) Wrap(err error, opts ...Options) error {

	var options Options
	if len(opts) > 0 {
		options = opts[0]
	}

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
		if options.ErasePath {
			// Если передана глубина пути, то используем ее
			if options.PathDepth != 0 {
				skip = options.PathDepth
			}
			customErr.Path = getPath(skip)
		}

		// Если передан текст ошибки, то добавляем его к исходной ошибке
		if options.ErrMessage != "" {
			customErr.InitialError = fmt.Errorf("%v: %v", options.ErrMessage, customErr.InitialError)
		}

		customErr.ErrorType = typ

	} else { // Если это не обернутая ошибка

		// Если передана глубина пути, то используем ее
		if options.PathDepth != 0 {
			skip = options.PathDepth
		}

		// Если это не обернутая ошибка, то создаем новую
		customErr = CustomError{
			ErrorType:    typ,
			InitialError: err,
			Path:         getPath(skip),
			Params:       options.Params,
			HumanText:    options.HumanText,
			LogAs:        TypeToLogOption[typ],
		}

		if options.ErrMessage != "" {
			customErr.InitialError = fmt.Errorf("%v: %v", options.ErrMessage, err)
		} else {
			customErr.InitialError = err
		}
	}

	// Добавляем параметры
	maps.Copy(customErr.Params, options.Params)

	if options.LogAs != nil {
		customErr.LogAs = *options.LogAs
	}

	return customErr
}

func CastError(err error) CustomError {

	customErr, ok := err.(CustomError)
	if !ok {
		err = InternalServer.Wrap(err, Options{
			ErrMessage: "Ошибка не обернута, путь неверный",
			PathDepth:  SecondPathDepth,
		})
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

	err.DevelopText = err.Error()

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
