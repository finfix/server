package services

import (
	"context"
	"reflect"

	"server/app/pkg/contextKeys"
	"server/app/pkg/errors"
	"server/app/pkg/reflectUtils"
)

type NecessaryUserInformation struct {
	UserID   uint32 `json:"-" schema:"-" validate:"required" minimum:"1"` // Идентификатор пользователя
	DeviceID string `json:"-" schema:"-" validate:"required"`             // Идентификатор устройства
}

func ExtractNecessaryFromCtx(ctx context.Context) (necessary NecessaryUserInformation, err error) {
	userID := contextKeys.GetUserID(ctx)
	if userID == nil {
		return necessary, errors.BadRequest.New("user id not found or not uint32")
	}
	deviceID := contextKeys.GetDeviceID(ctx)
	if deviceID == nil {
		return necessary, errors.BadRequest.New("device id not found or not string")
	}
	return NecessaryUserInformation{
		UserID:   *userID,
		DeviceID: *deviceID,
	}, nil
}

func SetNecessary(necessaryInformation NecessaryUserInformation, dest any) error {

	// Проверяем типы данных
	if err := reflectUtils.CheckPointerToStruct(dest); err != nil {
		return err
	}

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
