package middleware

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"server/app/pkg/contextKeys"
	"server/app/pkg/errors"
	"server/app/pkg/jwtManager"
)

func ExtractDataFromToken(ctx context.Context, r *http.Request) (context.Context, error) {
	ctx, err := DefaultAuthorization(ctx, r)
	if err != nil {
		if !errors.As(err, jwt.ValidationErrorExpired) {
			return ctx, err
		}
	}
	return ctx, nil
}

func DefaultAuthorization(ctx context.Context, r *http.Request) (context.Context, error) {
	userID, deviceID, err := jwtManager.Parse(r.Header.Get("Authorization"))
	if err != nil {
		ctx = contextKeys.SetDeviceID(ctx, deviceID)
		ctx = contextKeys.SetUserID(ctx, userID)
		return ctx, errors.Unauthorized.Wrap(err,
			errors.StackTraceOption(errors.PreviousCaller),
			errors.DontEraseErrorType(),
		)
	}
	if deviceID == "" {
		return ctx, errors.Unauthorized.New("DeviceID is empty",
			errors.StackTraceOption(errors.PreviousCaller),
		)
	}
	if userID == 0 {
		return ctx, errors.Unauthorized.New("UserID is empty",
			errors.StackTraceOption(errors.PreviousCaller),
		)
	}
	ctx = contextKeys.SetDeviceID(ctx, deviceID)
	ctx = contextKeys.SetUserID(ctx, userID)
	return ctx, nil
}
