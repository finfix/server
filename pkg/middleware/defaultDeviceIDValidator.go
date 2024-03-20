package middleware

import (
	"context"
	"net/http"

	"server/pkg/errors"
)

func DefaultDeviceIDValidator(ctx context.Context, r *http.Request) (context.Context, error) {
	if r.Header.Get("DeviceID") == "" {
		return ctx, errors.BadRequest.New("DeviceID is empty")
	}
	return context.WithValue(ctx, "DeviceID", r.Header.Get("DeviceID")), nil
}
