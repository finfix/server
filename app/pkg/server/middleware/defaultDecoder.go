package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gorilla/schema"

	"server/app/pkg/errors"
	"server/app/pkg/validation"
	"server/app/services"
)

type DecodeMethod int

const (
	DecodeSchema DecodeMethod = iota + 1
	DecodeJSON
)

type Validable interface {
	Validate() error
}

type NecessarySettable interface {
	SetNecessary(services.NecessaryUserInformation) any
}

type Decodable interface {
	Validable
	NecessarySettable
}

func DefaultDecoder[T Decodable](
	ctx context.Context,
	r *http.Request,
	decodeSchema DecodeMethod,
	_ T,
) (req T, err error) {

	if reflect.ValueOf(req).Kind() != reflect.Struct {
		return req, errors.InternalServer.New("Пришедший интерфейс не равен структуре", []errors.Option{
			errors.ParamsOption("Тип интерфейса", reflect.ValueOf(req).Kind().String()),
			errors.PathDepthOption(errors.SecondPathDepth),
		}...)
	}

	switch decodeSchema {
	case DecodeSchema:
		err = schema.NewDecoder().Decode(&req, r.URL.Query())
	case DecodeJSON:
		err = json.NewDecoder(r.Body).Decode(&req)
	}
	if err != nil {
		return req, errors.BadRequest.Wrap(err)
	}

	if err = req.Validate(); err != nil {
		return req, err
	}

	necessaryInformation, err := services.ExtractNecessaryFromCtx(ctx)
	if err != nil {
		return req, err
	}
	reqAny := req.SetNecessary(necessaryInformation)

	var ok bool
	if req, ok = reqAny.(T); !ok {
		return req, errors.InternalServer.New("Не удалось привести интерфейс к структуре", []errors.Option{
			errors.ParamsOption("Тип интерфейса", reflect.ValueOf(reqAny).Kind().String()),
			errors.PathDepthOption(errors.SecondPathDepth),
		}...)
	}

	return req, validation.ZeroValue(req)
}
