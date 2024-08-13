package validator

import (
	"github.com/go-playground/validator/v10"

	"server/app/pkg/errors"
)

type validatorProtocol interface {
	Validate() error
}

// Синглтон переменная валидатора
var validate = validator.New()

// Validate валидирует полученную структуру по тегам в декларативном формате
func Validate(data any) error {

	var (
		fields, tags []string
		values       []any
	)

	// Валидируем структуру
	stdErr := validate.Struct(data)

	// Если валидация прошла с ошибкой
	if stdErr != nil {

		// Приводим полученную ошибку к внутренней ошибке валидатора
		var validatorErrs validator.ValidationErrors
		if !errors.As(stdErr, &validatorErrs) {
			return errors.InternalServer.New("Не смогли закастить ошибку валидатора",
				errors.SkipThisCallOption(),
			)
		}

		// Проходимся по каждой ошибке валидации
		for _, validatorErr := range validatorErrs {

			// Заполняем дебаг-данными
			fields = append(fields, validatorErr.Field())
			tags = append(tags, validatorErr.Tag())
			values = append(values, validatorErr.Value())
		}

		return errors.BadRequest.Wrap(validatorErrs,
			errors.SkipThisCallOption(),
			errors.ParamsOption(
				"fields", fields,
				"tags", tags,
				"values", values,
			),
		)
	}

	// Если структура реализует интерфейс валидатора, то валидируем ее с помощью функции
	if v, ok := data.(validatorProtocol); ok {
		if err := v.Validate(); err != nil {
			return errors.BadRequest.Wrap(err,
				errors.SkipThisCallOption(),
			)
		}
	}

	return nil
}
