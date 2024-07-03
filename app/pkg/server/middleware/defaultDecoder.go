package middleware

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

type validatorProtocol interface {
	Validate() error
}

func DefaultDecoder(
	ctx context.Context,
	r *http.Request,
	decodeSchema DecodeMethod,
	dest any,
) (err error) {

	reflectVar := reflect.ValueOf(dest)
	if reflectVar.Kind() != reflect.Ptr || reflectVar.Elem().Kind() != reflect.Struct {
		return errors.InternalServer.New("Пришедший интерфейс является указателем на структуру",
			errors.ParamsOption("Тип интерфейса", reflectVar.Kind().String()),
			errors.PathDepthOption(errors.SecondPathDepth))
	}

	switch decodeSchema {
	case DecodeSchema:
		err = schema.NewDecoder().Decode(dest, r.URL.Query())
	case DecodeJSON:
		err = json.NewDecoder(r.Body).Decode(dest)
	}
	if err != nil {
		return errors.BadRequest.Wrap(
			err,
			errors.PathDepthOption(errors.SecondPathDepth),
		)
	}

	// Если структура реализует интерфейс валидатора, то валидируем ее с помощью функции
	if v, ok := dest.(validatorProtocol); ok {
		if err = v.Validate(); err != nil {
			return errors.BadRequest.Wrap(
				err,
				errors.PathDepthOption(errors.SecondPathDepth),
			)
		}
	}

	// Получаем необходимую для каждого запроса информацию из контекста
	necessaryInformation, err := services.ExtractNecessaryFromCtx(ctx)
	if err != nil {
		return errors.BadRequest.Wrap(
			err,
			errors.PathDepthOption(errors.SecondPathDepth),
		)
	}

	if err = setNessessary(necessaryInformation, dest); err != nil {
		return errors.InternalServer.Wrap(err,
			errors.PathDepthOption(errors.SecondPathDepth),
		)
	}

	if err = validator.Validate(dest); err != nil {
		return errors.BadRequest.Wrap(err,
			errors.PathDepthOption(errors.SecondPathDepth),
		)
	}

	return nil
}

func setNessessary(necessaryInformation services.NecessaryUserInformation, dest any) error {

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
