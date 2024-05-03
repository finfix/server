package middleware

import (
	"context"
	"net/http"

	"server/app/pkg/contextKeys"
	"server/app/pkg/errors"
	"server/app/pkg/jwtManager"
)

func ExtractDataFromToken(ctx context.Context, r *http.Request) (context.Context, error) {
	userID, deviceID, err := jwtManager.Parse(r.Header.Get("Authorization"))
	if err != nil {
		if !errors.As(err, errors.Unauthorized) {
			return ctx, err
		}
	}
	if deviceID == "" {
		return ctx, errors.Unauthorized.New("DeviceID is empty")
	}
	if userID == 0 {
		return ctx, errors.Unauthorized.New("UserID is empty")
	}
	ctx = context.WithValue(ctx, contextKeys.DeviceIDKey, deviceID)
	ctx = context.WithValue(ctx, contextKeys.UserIDKey, userID)
	return ctx, nil
}

func DefaultAuthorization(ctx context.Context, r *http.Request) (context.Context, error) {
	userID, deviceID, err := jwtManager.Parse(r.Header.Get("Authorization"))
	if err != nil {
		return ctx, err
	}
	if deviceID == "" {
		return ctx, errors.Unauthorized.New("DeviceID is empty")
	}
	if userID == 0 {
		return ctx, errors.Unauthorized.New("UserID is empty")
	}
	ctx = context.WithValue(ctx, contextKeys.DeviceIDKey, deviceID)
	ctx = context.WithValue(ctx, contextKeys.UserIDKey, userID)
	return ctx, nil
}
