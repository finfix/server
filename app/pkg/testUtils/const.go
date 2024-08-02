package testUtils

import (
	"server/app/pkg/contextKeys"
	"server/app/services"
)

const (
	InvalidJSON = "{invalid}"
)

var (
	GeneralCtx = NewCtxBuilder().
		Set(contextKeys.DeviceIDKey, "DeviceID").
		Set(contextKeys.UserIDKey, uint32(1))

	ValidNecessary = services.NecessaryUserInformation{
		UserID:   1,
		DeviceID: "DeviceID",
	}
)
