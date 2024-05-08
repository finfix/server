package middleware

import (
	"context"
	"net/http"

	"server/app/pkg/errors"
	"server/app/pkg/log"
)

func DefaultErrorEncoder(ctx context.Context, w http.ResponseWriter, er error) {

	if er == nil {
		er = errors.InternalServer.New("В функцию DefaultErrorEncoder передана пустая ошибка", []errors.Option{
			errors.PathDepthOption(errors.SecondPathDepth),
		}...)
	}

	err := errors.CastError(er)
	err.HumanText = humanTextByLevel[err.ErrorType]
	err.AdditionalInfo = log.ExtractAdditionalInfo(ctx)

	switch err.LogAs {
	case errors.LogAsError:
		log.Error(ctx, err)
	case errors.LogAsWarning:
		log.Warning(ctx, err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(int(err.ErrorType))

	byt, er := errors.JSON(err)
	if er != nil {
		log.Error(ctx, er)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(er.Error()))
	}

	_, _ = w.Write(byt)
}

var humanTextByLevel = map[errors.ErrorType]string{
	errors.BadRequest:     "Введены неверные данные",
	errors.InternalServer: "Произошла непредвиденная ошибка",
	errors.NotFound:       "Данные не найдены",
	errors.Forbidden:      "Доступ запрещен",
	errors.Teapot:         "Разработчик забыл написать текст ошибки",
	errors.BadGateway:     "Произошла ошибка на сервере внешнего сервиса",
	errors.Unauthorized:   "Пользователь не авторизован",
	errors.ClientReject:   "Клиент отказался принимать данные",
	errors.LogicError:     "Произошла непредвиденная ошибка",
}
