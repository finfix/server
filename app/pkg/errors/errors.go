package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"runtime"
	"strings"
)

const (
	defaultPathDepth = iota + 1
	SecondPathDepth
	ThirdPathDepth
	FourthPathDepth
)

type CustomError struct {
	ErrorType      `json:"-"`
	HumanText      string            `json:"humanTextError"`
	DevelopText    string            `json:"developerTextError"`
	InitialError   error             `json:"-"`
	Path           []string          `json:"path"`
	Params         map[string]string `json:"parameters,omitempty" validate:"required"`
	AdditionalInfo map[string]string `json:"additionalInfo,omitempty"`
	LogAs          LogOption         `json:"-"`
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
		ErrorType:      typ,
		HumanText:      options.HumanText,
		DevelopText:    msg,
		InitialError:   errors.New(msg),
		Path:           GetPath(skip + 1),
		Params:         options.params,
		AdditionalInfo: nil,
		LogAs:          TypeToLogOption[typ],
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

		// Если передана глубина пути, то используем ее
		if options.pathDepth != nil {
			skip = *options.pathDepth
		}
		customErr.Path = GetPath(skip)

		// Если передан текст ошибки, то добавляем его к исходной ошибке
		if options.errMessage != nil {
			customErr.InitialError = fmt.Errorf("%v: %v", options.errMessage, customErr.InitialError)
		}

		if options.dontEraseErrorType == nil {
			customErr.ErrorType = typ
		}

	} else { // Если это не обернутая ошибка

		// Если передана глубина пути, то используем ее
		if options.pathDepth != nil {
			skip = *options.pathDepth
		}

		// Если это не обернутая ошибка, то создаем новую
		customErr = CustomError{
			ErrorType:      typ,
			HumanText:      options.HumanText,
			DevelopText:    err.Error(),
			InitialError:   err,
			Path:           GetPath(skip + 1),
			Params:         options.params,
			AdditionalInfo: nil,
			LogAs:          TypeToLogOption[typ],
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
		err = InternalServer.Wrap(errors.New("Ошибка не обернута, путь неверный"), []Option{
			PathDepthOption(SecondPathDepth),
		}...)
		customErr, _ = err.(CustomError)
	}
	return customErr
}

func GetPath(skip int) []string {
	var pcs [32]uintptr
	n := runtime.Callers(skip, pcs[:])
	var path []string
	for i := skip; i < n; i++ {
		_, file, line, _ := runtime.Caller(i)
		if strings.Contains(file, "coin") || strings.Contains(file, "Coin") {
			path = append(path, fmt.Sprintf("%v:%v", file, line))
		}
	}
	return path
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

func Is(get error, target error) bool {
	return errors.Is(get, target)
}
