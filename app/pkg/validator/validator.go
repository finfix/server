package validator

import (
	"github.com/go-playground/validator/v10"

	"server/app/pkg/errors"
)

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
		validatorErrs, ok := stdErr.(validator.ValidationErrors)
		if !ok {
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

	return nil
}
