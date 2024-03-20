package validation

import (
	"reflect"

	"server/pkg/errors"

	"regexp"
)

func Mail(email string) error {
	matched, err := regexp.MatchString(`^([a-z0-9_-]+\.)*[a-z0-9_-]+@[a-z0-9_-]+(\.[a-z0-9_-]+)*\.[a-z]{2,6}$`, email)
	if err != nil {
		return errors.InternalServer.WrapCtx(err, "Email: %v", email)
	}
	if !matched {
		err = errors.BadRequest.NewCtx("Invalid email address provided", "Email: %v", email)
		return errors.AddHumanText(err, "Неверно введен адрес электронной почты")
	}
	return nil
}

func zeroValue(requestStruct any, tag string) error {

	if reflect.ValueOf(requestStruct).Kind() != reflect.Struct {
		return errors.InternalServer.NewPathCtx("Пришедший интерфейс не равен структуре", 2, "Тип структуры: %s", reflect.ValueOf(requestStruct).Kind().String())
	}

	// Получаем тип данных структуры (ждем обязательно структуру)
	t := reflect.TypeOf(requestStruct)

	// Получаем значение структуры
	v := reflect.ValueOf(requestStruct)

	// Проходимся по каждому полю структуры
	for i := 0; i < t.NumField(); i++ {

		// Получаем поле
		typeField := t.Field(i)

		switch typeField.Name {
		case "state", "sizeCache", "unknownFields":
			continue
		}

		// Получаем значение поля
		valField := v.Field(i)

		// Получаем то, что в теге validate
		reqTag := typeField.Tag.Get("validate")

		// Получаем тег json
		jsTag := typeField.Tag.Get("json")

		// Если поле равно нулю и тег validate = required
		if valField.IsZero() && reqTag == "required" {

			if tag != "" {
				tag += "."
			}

			return errors.BadRequest.NewPath("Required field \""+tag+jsTag+"\" is not filled", 2)
		}

		// Если тип поля структура
		if typeField.Type.Kind() == reflect.Struct {

			// Приводим к интерфейсу
			tt := valField.Interface()

			// Добавляем вложенность
			if len(tag) != 0 && i == 0 {
				tag = tag + "."
			}

			// Рекурсивно вызываем функцию для вложенной функции
			err := zeroValue(tt, tag+jsTag)

			// Если внутри структуры плохо
			if err != nil {
				return errors.BadRequest.NewCtx("Required field is not filled", "Поле: %v", tag+jsTag)
			}
		}
	}

	// Возвращаем гуд
	return nil
}

func ZeroValue(requestStruct any) error {
	return zeroValue(requestStruct, "")
}
