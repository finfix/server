package chain

import (
	"context"
	"net/http"

	"server/pkg/contextKeys"
	"server/pkg/errors"
	"server/pkg/jwtManager"
)

func DefaultAuthorization(ctx context.Context, r *http.Request) (context.Context, error) {
	userID, deviceID, err := jwtManager.Parse(r.Header.Get("Authorization"))
	if err != nil {
		if deviceID != "" {
			ctx = contextKeys.SetDeviceID(ctx, deviceID)
		}
		if userID != 0 {
			ctx = contextKeys.SetUserID(ctx, userID)
		}
		return ctx, errors.Unauthorized.Wrap(err,
			errors.SkipThisCallOption(),
			errors.DontEraseErrorType(),
		)
	}
	if deviceID == "" {
		return ctx, errors.Unauthorized.New("DeviceID is empty",
			errors.SkipThisCallOption(),
		)
	}
	if userID == 0 {
		return ctx, errors.Unauthorized.New("UserID is empty",
			errors.SkipThisCallOption(),
		)
	}
	ctx = contextKeys.SetDeviceID(ctx, deviceID)
	ctx = contextKeys.SetUserID(ctx, userID)

	return ctx, nil
}
