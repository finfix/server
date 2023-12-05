package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime"

	"google.golang.org/grpc/status"

	"pkg/errors/pbError"
)

// Определяем типы ошибок
const (
	NoType         = ErrorType(2)
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

type ErrorType uint32

type CustomError struct {
	ErrorType   `json:"-"`
	HumanText   string  `json:"humanTextError"`
	DevelopText string  `json:"developerTextError"`
	Err         error   `json:"-"`
	Path        string  `json:"path"`
	Context     *string `json:"context,omitempty" validate:"required"`
}

func (error CustomError) Error() string {
	return error.Err.Error()
}

// Создаем новую ошибку
func (typ ErrorType) New(msg string) error {
	return typ.NewPathCtx(msg, 2, "")
}

// Создаем новую ошибку с контекстом
func (typ ErrorType) NewCtx(msg, context string, args ...any) error {
	return typ.NewPathCtx(msg, 2, context, args...)
}

// Создаем новую ошибку с выбором глубины пути (от 1)
func (typ ErrorType) NewPath(msg string, skip int) error {
	return typ.NewPathCtx(msg, skip+1, "")
}

// Создаем новую ошибку с выбором глубины пути (от 1) и контекстом
func (typ ErrorType) NewPathCtx(msg string, skip int, context string, args ...any) error {
	_, file, line, _ := runtime.Caller(skip)

	customErr := CustomError{
		ErrorType: typ,
		Err:       errors.New(msg),
		Path:      fmt.Sprintf("%v:%v", file, line),
	}
	if context != "" {
		context := fmt.Sprintf(context, args...)
		customErr.Context = &context
	}
	return customErr
}

// Оборачиваем дефолтную ошибку в кастомный тип
// Если в функцию передается кастомная ошибка, то она просто возвращается
func (typ ErrorType) Wrap(err error) error {
	return typ.WrapPathCtx(err, 2, "")
}

// Оборачиваем дефолтную ошибку в кастомный тип и задаем ей контекст
// Если в функцию передается кастомная ошибка, то она просто возвращается с добавлением контекста
func (typ ErrorType) WrapCtx(err error, context string, agrs ...any) error {
	return typ.WrapPathCtx(err, 2, context, agrs...)
}

// Оборачиваем дефолтную ошибку в кастомный тип с выбором глубины пути (от 1)
// Если в функцию передается кастомная ошибка, то она просто возвращается
func (typ ErrorType) WrapPath(err error, skip int) error {
	return typ.WrapPathCtx(err, skip+1, "")
}

// Оборачиваем дефолтную ошибку в кастомный тип с выбором глубины пути (от 1) и контекстом.
// Если в функцию передается кастомная ошибка, то она просто возвращается с добавлением контекста
func (typ ErrorType) WrapPathCtx(err error, skip int, context string, args ...any) error {

	if err == nil {
		return nil
	}

	err = ConvertGrpcErrorToCustomError(err)

	// Если это уже обернутая ошибка, добавляем контекст, меняем тип и возвращаем
	if customErr, ok := err.(CustomError); ok {
		//customErr.ErrorType = typ
		if customErr.Context != nil && context != "" {
			newContext := fmt.Sprintf("%s. %s", context, *customErr.Context)
			customErr.Context = &newContext
		}
		return customErr
	}

	// Если это новая ошибка, то оборачиваем ее в нашу кастомную
	_, file, line, _ := runtime.Caller(skip)

	customError := CustomError{
		ErrorType: typ,
		Err:       err,
		Path:      fmt.Sprintf("%v:%v", file, line),
	}
	if context != "" {
		context := fmt.Sprintf(context, args...)
		customError.Context = &context
	}
	return customError
}

func ConvertGrpcErrorToCustomError(err error) error {

	// Если это ошибка gRPC
	if st, ok := status.FromError(err); ok {

		// Если это наша кастомная ошибка, то достаем ее данные
		if details := st.Details(); len(details) > 0 {
			if errProto, ok := details[0].(*pbError.CustomError); ok {
				err = CustomError{
					ErrorType:   ErrorType(errProto.ErrorType),
					HumanText:   errProto.HumanText,
					Err:         fmt.Errorf(errProto.Err),
					Path:        errProto.Path,
					Context:     errProto.Context,
					DevelopText: errProto.DevelopText,
				}
				return err
			}
		}

		// Если это внутренняя ошибка gRPC, то получаем ее данные и оборачиваем
		err = errors.New(st.Message())
	}
	return err
}

// Добавляем в ошибку текст, который можно отдать пользователю
func AddHumanText(err error, message string) error {

	if customErr, ok := err.(CustomError); ok {

		if customErr.HumanText != "" {
			return err
		}

		return CustomError{
			ErrorType: customErr.ErrorType,
			Err:       customErr.Err,
			HumanText: message,
			Path:      customErr.Path,
			Context:   customErr.Context,
		}
	}

	_, file, line, _ := runtime.Caller(1)

	return CustomError{
		ErrorType: NoType,
		Err:       err,
		HumanText: message,
		Path:      fmt.Sprintf("%v:%v", file, line),
	}
}

// Получаем тип ошибки
func GetType(err error) ErrorType {

	if customErr, ok := err.(CustomError); ok {
		return customErr.ErrorType
	}

	return NoType
}

// Переводим ошибку в JSON
func Json(err error) ([]byte, error) {

	if customErr, ok := err.(CustomError); ok {

		customErr.DevelopText = customErr.Err.Error()

		byt, e := json.Marshal(customErr)
		if e != nil {
			return nil, InternalServer.Wrap(e)
		}

		return byt, nil
	}

	return nil, InternalServer.NewCtx("Дефолтная ошибка не обернута. Ошибка: %v", err.Error())
}

func As(get error, target any) bool {
	if target != nil || get != nil {
		return errors.As(get, &target)
	}
	return false
}
