package testingFunc

import (
	"server/app/pkg/contextKeys"
)

const (
	InvalidJSON = "{invalid}"
)

var (
	GeneralCtx = NewCtxBuilder().
		Set(contextKeys.DeviceIDKey, "DeviceID").
		Set(contextKeys.UserIDKey, uint32(1))
)
