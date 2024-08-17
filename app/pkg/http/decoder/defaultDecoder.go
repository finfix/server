package decoder

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gorilla/schema"

	"server/app/pkg/errors"
	"server/app/pkg/validator"
	"server/app/services"
)

type DecodeMethod int

const (
	DecodeSchema DecodeMethod = iota + 1
	DecodeJSON
)

func Decoder(
	ctx context.Context,
	r *http.Request,
	dest any,
	decodeSchemas ...DecodeMethod,
) (err error) {

	// Проверяем типы данных
	reflectVar := reflect.ValueOf(dest)
	if reflectVar.Kind() != reflect.Ptr || reflectVar.Elem().Kind() != reflect.Struct {
		return errors.InternalServer.New("Пришедший интерфейс является указателем на структуру",
			errors.ParamsOption("Тип интерфейса", reflectVar.Kind().String()),
			errors.SkipThisCallOption())
	}

	// Проходимся по каждому
	for _, decodeSchema := range decodeSchemas {
		switch decodeSchema {
		case DecodeSchema:
			err = schema.NewDecoder().Decode(dest, r.URL.Query())
		case DecodeJSON:
			err = json.NewDecoder(r.Body).Decode(dest)
		default:
			break
		}
		if err != nil {
			return errors.BadRequest.Wrap(
				err,
				errors.SkipThisCallOption(),
			)
		}
	}

	// Получаем необходимую для каждого запроса информацию из контекста
	necessaryInformation, err := services.ExtractNecessaryFromCtx(ctx)
	if err != nil {
		return errors.BadRequest.Wrap(
			err,
			errors.SkipThisCallOption(),
		)
	}

	// Заполняем необходимую для каждого запроса информацию в структуру
	if err = SetNecessary(necessaryInformation, dest); err != nil {
		return errors.InternalServer.Wrap(err,
			errors.SkipThisCallOption(),
		)
	}

	// Валидируем получившуюся структуру
	if err = validator.Validate(dest); err != nil {
		return errors.BadRequest.Wrap(err,
			errors.SkipThisCallOption(),
		)
	}

	return nil
}

func SetNecessary(necessaryInformation services.NecessaryUserInformation, dest any) error {

	// Получаем указатель на структуру
	reflectVar := reflect.ValueOf(dest).Elem()

	// Ищем поле с именем "Necessary"
	necessaryField := reflectVar.FieldByName("Necessary")

	// Если такого поля нет, тогда выходим из функции
	if !necessaryField.IsValid() {
		return nil
	}

	// Проверяем, является ли поле экспортированным и можно ли его устанавливать
	if !necessaryField.CanSet() {
		return errors.InternalServer.New(
			"Поле Necessary является неэкспортируемым",
		)
	}

	// Получаем значение структуры necessaryData с использованием отражения
	necessaryValue := reflect.ValueOf(necessaryInformation)

	// Устанавливаем значение поля
	necessaryField.Set(necessaryValue)

	return nil
}
