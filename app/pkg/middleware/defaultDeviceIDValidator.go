package middleware

import (
	"context"
	"net/http"

	"server/app/pkg/contextKeys"
	"server/app/pkg/errors"
)

func DefaultDeviceIDValidator(ctx context.Context, r *http.Request) (context.Context, error) {
	if r.Header.Get("DeviceID") == "" {
		return ctx, errors.BadRequest.New("DeviceID is empty")
	}
	return context.WithValue(ctx, contextKeys.DeviceIDKey, r.Header.Get("DeviceID")), nil
}
