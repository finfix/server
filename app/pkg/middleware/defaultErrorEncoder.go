package middleware

import (
	"context"
	"net/http"

	"server/app/pkg/errors"
	"server/app/pkg/logging"
)

func DefaultErrorEncoder(_ context.Context, w http.ResponseWriter, er error, loggingFunc func(error)) {

	if er == nil {
		er = errors.InternalServer.New("В функцию DefaultErrorEncoder передана пустая ошибка", errors.Options{
			PathDepth: errors.SecondPathDepth,
		})
	}

	err := errors.CastError(er)
	err.HumanText = humanTextByLevel[err.ErrorType]

	loggingFunc(err)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(int(err.ErrorType))

	byt, er := errors.JSON(err)
	if er != nil {
		logging.GetLogger().Error(er)
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
