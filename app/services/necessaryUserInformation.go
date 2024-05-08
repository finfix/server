package services

import (
	"context"

	"server/app/pkg/contextKeys"
	"server/app/pkg/errors"
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
