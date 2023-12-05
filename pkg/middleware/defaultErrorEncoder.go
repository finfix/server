package middleware

import (
	"context"
	"fmt"
	"net/http"

	"pkg/errors"
)

func DefaultErrorEncoder(_ context.Context, w http.ResponseWriter, err error, loggingFunc func(error)) {

	if err == nil {
		err = errors.InternalServer.NewPath("В функцию DefaultErrorEncoder передан nil", 2)
	}

	err = errors.ConvertGrpcErrorToCustomError(err)

	err = addDefaultHumanText(err)

	loggingFunc(err)

	errorType := errors.GetType(err)

	w.Header().Set("Content-Type", "application/jsonapi; charset=utf-8")
	w.WriteHeader(int(errorType))

	byt, e := errors.Json(err)
	if e != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(e.Error()))
	}

	_, _ = w.Write(byt)
}

func addDefaultHumanText(err error) error {
	errorType := errors.GetType(err)

	switch errorType {
	case errors.BadRequest:
		err = errors.AddHumanText(err, "Введены неверные данные")
	case errors.InternalServer:
		err = errors.AddHumanText(err, "Произошла непредвиденная ошибка")
	case errors.NotFound:
		err = errors.AddHumanText(err, "Данные не найдены")
	case errors.Forbidden:
		err = errors.AddHumanText(err, "Доступ запрещен")
	case errors.Teapot:
		err = errors.AddHumanText(err, "Тут забыли написать текст ошибки")
	case errors.BadGateway:
		err = errors.AddHumanText(err, "Произошла ошибка на сервере внешнего сервиса")
	case errors.Unauthorized:
		err = errors.AddHumanText(err, "Пользователь не авторизован")
	case errors.NoType:
		err = errors.InternalServer.WrapCtx(err, "Ошибка не имеет обертки, путь неверный")
		err = errors.AddHumanText(err, "Произошла непредвиденная ошибка")
	case errors.LogicError:
		err = errors.InternalServer.WrapCtx(err, "Рабочая ошибка не должна попасть в middleware")
		err = errors.AddHumanText(err, "Произошла непредвиденная ошибка")
	}

	return err
}
