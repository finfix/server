package decoder

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"

	"server/pkg/errors"
	"server/pkg/reflectUtils"
	"server/pkg/validator"
	"server/internal/services"
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
	if err = reflectUtils.CheckPointerToStruct(dest); err != nil {
		return err
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
	if err = services.SetNecessary(necessaryInformation, dest); err != nil {
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
