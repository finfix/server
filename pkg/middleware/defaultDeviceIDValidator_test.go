package middleware

import (
	"context"
	"net/http/httptest"
	"testing"

	"server/pkg/errors"
	"server/pkg/testingFunc"
)

func TestGetDeviceID(t *testing.T) {

	for _, tt := range []struct {
		name     string
		deviceID string
		err      error
	}{
		{
			name:     "1.Empty DeviceID",
			deviceID: "",
			err:      errors.BadRequest.New("DeviceID is empty"),
		},
		{
			name:     "2.With DeviceID",
			deviceID: "test",
			err:      nil,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {

			req := httptest.NewRequest("", "/", nil)

			if tt.deviceID != "" {
				req.Header.Add("DeviceID", tt.deviceID)
			}

			ctx, err := DefaultDeviceIDValidator(context.Background(), req)

			if testingFunc.CheckError(t, tt.err, err) {
				return
			}

			getDeviceID := ctx.Value("DeviceID").(string)

			if tt.deviceID != getDeviceID {
				t.Fatalf("\nDeviceID не совпадают: %v != %v", tt.deviceID, getDeviceID)
			}
		})
	}
}
