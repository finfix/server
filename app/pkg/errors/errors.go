package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"maps"

	"server/app/pkg/log/model"
	"server/app/pkg/stackTrace"
)

// Error - Кастомная структура ошибки
type Error struct {

	// Тип ошибки, в который зашит HTTP-код
	// В случае, если этот тип снова кладется в errors.Type.Wrap, эта переменная затирается
	// Оставить первоначатльный тип ошибки можно через errors.DontEraseErrorType()
	ErrorType ErrorType `json:"-"`

	// Первоначальная ошибка. Если необходимо завернуть эту ошибку через fmt.Errorf("%w", err), то
	// Необходимо воспользоваться errors.ErrorfOption("additionalInfo: %w")
	Err error `json:"-"`

	// Поскольку стандартный энкодер json в го не умеет нормально сериализовать тип ошибок, эта переменная
	// Используется для подставления значения Err прямо перед сериализацией ошибки в функции JSON
	DeveloperText string `json:"error"`

	// Человекочитаемый текст, который можно показать клиенту
	// Переменная настраивается через errors.HumanTextOption(messageWithFmt, args...)
	// Если значения нет, то автоматически проставляется шаблонными данными в функции middleware.DefaultErrorEncoder
	HumanText string `json:"humanText"`

	// Стектрейс от места враппинга ошибки. Если необходимо начать стектрейс с уровня выше, то
	// Необходимо воспользоваться errors.SkipThisCallOption(errors.<const>)
	// const = SkipThisCall - начать стектрейс на один уровень выше враппера errors.Type.Wrap по дереву
	// const = SkipPreviousCaller и остальные работают по аналогии, пропуская все больше уровней стека вызовов
	StackTrace []string `json:"path"`

	// Дополнительные параметры, направленные на дополнение ошибки контекстом, которые проставляются
	// Через errors.ParamsOption(key1, value1, key2, value2, ...)
	Params map[string]string `json:"parameters"`

	// Служебное поле, которое автоматически заполняется в функции middleware.DefaultErrorEncoder
	// вспомогательными данными из контекста
	UserInfo *model.UserInfo `json:"userInfo"`

	// Служебное поле, которое автоматически заполняется в функции middleware.DefaultErrorEncoder
	// вспомогательными данными из контекста
	SystemInfo model.SystemInfo `json:"systemInfo"`

	// Параметр, определяющий уровень логгирования ошибки в функции middleware.DefaultErrorEncoder
	// Настраивается через errors.LogAsOption(LogOption)
	LogAs LogOption `json:"-"`
}

// Error реализует протокол ошибок для использования нашей структуры в качестве error параметра
func (err Error) Error() string {
	return err.Err.Error()
}

// LogOption - Перечисление, необходимое для конкретизации уровня логгирования ошибки
type LogOption int

const (
	LogAsError LogOption = iota
	LogAsWarning
	LogNone
)

// TypeToLogOption - Дефолтные настройки для логгирования каждого типа ошибок
var TypeToLogOption = map[ErrorType]LogOption{
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

	skip := stackTrace.ThisCall
	if options.stackTrace != nil {
		skip = *options.stackTrace
	}

	// Создаем новую ошибку
	customErr := Error{
		SystemInfo: model.SystemInfo{ // Будет заполняться автоматически в функции логгирования
			Hostname: "",
			Version:  "",
			Build:    "",
			Env:      "",
		},
		ErrorType:     typ,
		DeveloperText: "",
		HumanText:     options.HumanText,
		Err:           errors.New(msg),
		StackTrace:    stackTrace.GetStackTrace(skip + 1),
		Params:        options.params,
		UserInfo:      nil,
		LogAs:         TypeToLogOption[typ],
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

	skip := stackTrace.ThisCall

	var customErr Error

	if As(err, &customErr) { // Если это уже обернутая ошибка

		// Если передан текст для пользователя, то затираем его
		if options.HumanText != "" {
			customErr.HumanText = options.HumanText
		}

		// Если передана глубина пути, то используем ее
		if options.stackTrace != nil {
			skip = *options.stackTrace
		}
		customErr.StackTrace = stackTrace.GetStackTrace(skip + 1)

		// Если передан текст ошибки, то добавляем его к исходной ошибке
		if options.errorf != nil {
			customErr.Err = fmt.Errorf("%w: %w", *options.errorf, customErr.Err)
		}

		if options.dontEraseErrorType == nil {
			customErr.ErrorType = typ
		}

	} else { // Если это не обернутая ошибка

		// Если передана глубина пути, то используем ее
		if options.stackTrace != nil {
			skip = *options.stackTrace
		}

		// Если это не обернутая ошибка, то создаем новую
		customErr = Error{
			SystemInfo: model.SystemInfo{ // Будет автоматически заполняться в функции логгирования
				Hostname: "",
				Version:  "",
				Build:    "",
				Env:      "",
			},
			ErrorType:     typ,
			DeveloperText: "",
			HumanText:     options.HumanText,
			Err:           err,
			StackTrace:    stackTrace.GetStackTrace(skip + 1),
			Params:        options.params,
			UserInfo:      nil,
			LogAs:         TypeToLogOption[typ],
		}

		if options.errorf != nil {
			customErr.Err = fmt.Errorf("%w: %w", *options.errorf, err)
		}
	}

	// Добавляем параметры
	maps.Copy(customErr.Params, options.params)

	if options.logAs != nil {
		customErr.LogAs = *options.logAs
	}

	return customErr
}

// CastError приводит приедшую ошибку к нашей кастомной ошибке, если пришедшая ошибка не кастомная
// То оборачиает ее и добавляет данные о том, что ошибка не обернута
func CastError(err error) Error {
	var customErr Error
	if !As(err, &customErr) {
		err = InternalServer.Wrap(err,
			SkipThisCallOption(),
			ParamsOption("error", "Ошибка не обернута, путь неверный"),
		)
		_ = As(err, &customErr)
	}
	return customErr
}

// JSON преобразует ошибку в json
func JSON(err Error) ([]byte, error) {

	// Подставляем значение ошибки в текстовую переменную DeveloperTextError, поскольку сериализатор не умеет
	// нормально обрабатывать тип error
	err.DeveloperText = err.Err.Error()
	byt, e := json.Marshal(err)
	if e != nil {
		return nil, InternalServer.Wrap(e)
	}

	return byt, nil
}

// As используется для вызова стандартной функции As
func As(get error, target any) bool {
	return errors.As(get, target)
}

// Unwrap используется для разворачивания завернутых с помощью fmt.Errorf("%w", err) ошибок
// default(default(1)) -> default(1)
// custom(default(default(1))) -> custom(default(1))
func Unwrap(err error) error {
	var customErr Error
	if As(err, &customErr) {
		customErr.Err = errors.Unwrap(customErr.Err)
		return customErr
	} else {
		return errors.Unwrap(err)
	}
}

// Is используется для проверки типов любой комбинации дефолтных и кастомных ошибок
func Is(err error, target error) bool {

	var customErr, customTarget Error
	if As(err, &customErr) {
		if As(target, &customTarget) {
			return errors.Is(customErr.Err, customTarget.Err) // custom - custom
		} else {
			return errors.Is(customErr.Err, target) // custom - default
		}
	} else {
		if As(target, &customTarget) {
			return errors.Is(err, customTarget.Err) // default - custom
		} else {
			return errors.Is(err, target) // default - default
		}
	}
}

func New(err string) error {
	return errors.New(err)
}
