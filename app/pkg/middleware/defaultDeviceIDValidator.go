package middleware

import (
	"context"
	"net/http"

	"server/app/pkg/contextKeys"
	"server/app/pkg/errors"
)

func DefaultDeviceIDValidator(ctx context.Context, r *http.Request) (context.Context, error) {
	deviceID := r.Header.Get("DeviceID")
	if deviceID == "" {
		return ctx, errors.BadRequest.New("DeviceID is empty")
	}
	return contextKeys.SetDeviceID(ctx, deviceID), nil
}
