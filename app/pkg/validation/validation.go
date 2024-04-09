package validation

import (
	"reflect"
	"regexp"

	"server/app/pkg/errors"
)

func Mail(email string) error {
	matched, err := regexp.MatchString(`^([a-z0-9_-]+\.)*[a-z0-9_-]+@[a-z0-9_-]+(\.[a-z0-9_-]+)*\.[a-z]{2,6}$`, email)
	if err != nil {
		return errors.InternalServer.Wrap(err, errors.Options{
			Params: map[string]any{"email": email},
		})
	}
	if !matched {
		return errors.BadRequest.New("Invalid email address provided", errors.Options{
			Params:    map[string]any{"email": email},
			HumanText: "Неверно введен адрес электронной почты",
		})
	}
	return nil
}

func zeroValue(requestStruct any, tag string, depth int) error {

	if reflect.ValueOf(requestStruct).Kind() != reflect.Struct {
		return errors.InternalServer.New("Пришедший интерфейс не равен структуре", errors.Options{
			Params:    map[string]any{"Тип структуры": reflect.ValueOf(requestStruct).Kind().String()},
			PathDepth: errors.SecondPathDepth,
		})
	}

	// Получаем тип данных структуры (ждем обязательно структуру)
	reflectType := reflect.TypeOf(requestStruct)

	// Получаем значение структуры
	reflectValue := reflect.ValueOf(requestStruct)

	// Проходимся по каждому полю структуры
	for i := 0; i < reflectType.NumField(); i++ {

		// Получаем поле
		typeField := reflectType.Field(i)

		switch typeField.Name {
		case "state", "sizeCache", "unknownFields":
			continue
		}

		// Получаем значение поля
		valField := reflectValue.Field(i)

		// Получаем то, что в теге validate
		reqTag := typeField.Tag.Get("validate")

		// Получаем тег json
		jsTag := typeField.Tag.Get("json")

		// Если поле равно нулю и тег validate = required
		if valField.IsZero() && reqTag == "required" {

			if tag != "" {
				tag += "."
			}

			return errors.BadRequest.New("Required field is not filled", errors.Options{
				PathDepth: depth,
				Params:    map[string]any{"field": tag + jsTag},
			})
		}

		// Если тип поля структура
		if typeField.Type.Kind() == reflect.Struct {

			// Приводим к интерфейсу
			tt := valField.Interface()

			// Добавляем вложенность
			if len(tag) != 0 && i == 0 {
				tag += "."
			}

			// Рекурсивно вызываем функцию для вложенной функции
			err := zeroValue(tt, tag+jsTag, depth+1)

			// Если внутри структуры плохо
			if err != nil {
				return err
			}
		}
	}

	// Возвращаем гуд
	return nil
}

func ZeroValue(requestStruct any) error {
	return zeroValue(requestStruct, "", errors.ThirdPathDepth)
}
