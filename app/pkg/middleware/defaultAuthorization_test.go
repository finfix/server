package middleware

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"

	"server/app/pkg/contextKeys"
	"server/app/pkg/errors"
	"server/app/pkg/pointer"
	"server/app/pkg/testingFunc"
)

func TestAuthorization(t *testing.T) {

	type createTokenParams = struct {
		ttl      time.Duration
		userID   uint32
		deviceID string
	}

	validParams := createTokenParams{
		ttl:      time.Hour,
		userID:   1,
		deviceID: "wantDeviceID",
	}

	NewAuthMiddleware("test")

	for _, tt := range []struct {
		message string
		token   *string
		params  *createTokenParams
		err     error
	}{
		{
			"1.Валидный токен",
			nil,
			&createTokenParams{
				userID:   validParams.userID,
				deviceID: validParams.deviceID,
				ttl:      validParams.ttl,
			},
			nil,
		},
		{
			"2.Токен с истекшим сроком",
			nil,
			&createTokenParams{
				userID:   validParams.userID,
				deviceID: validParams.deviceID,
				ttl:      -time.Hour,
			},
			errors.Unauthorized.New("token is expired by 1h0m0s"),
		},
		{
			"3.Невалидный токен",
			pointer.Pointer("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUyNTg5ODMyMTEsInN1YiI6IjEifQ.TneMNhesU3VT0XVGb8dGg8zyyObrmPk_x9kdh-aJDwQ"),
			nil,
			errors.Unauthorized.Wrap(jwt.ErrSignatureInvalid),
		},
		{
			"4.Пустой токен",
			pointer.Pointer(""),
			nil,
			errors.Unauthorized.New("JWT-token is empty"),
		},
		{
			"5.Токен без DeviceID",
			nil,
			&createTokenParams{
				userID: validParams.userID,
				ttl:    validParams.ttl,
			},
			errors.Unauthorized.New("DeviceID is empty"),
		},
		{
			"6.Токен без UserID",
			nil,
			&createTokenParams{
				deviceID: validParams.deviceID,
				ttl:      validParams.ttl,
			},
			errors.Unauthorized.New("UserID is empty"),
		},
	} {
		t.Run(tt.message, func(t *testing.T) {

			tt := tt

			// Если токен не передан, то создаем его
			if tt.token == nil {

				// Создаем токен
				token, err := jwtManager.NewJWT(tt.params.userID, "test", tt.params.deviceID, tt.params.ttl)
				if err != nil {
					t.Fatalf("\nНе смогли создать JWV-токен: %v", err)
				}
				tt.token = &token
			}

			req := httptest.NewRequest("", "/", nil)
			req.Header.Add("Authorization", *tt.token)

			ctx, err := DefaultAuthorization(context.Background(), req)

			if testingFunc.CheckError(t, tt.err, err) {
				return
			}

			getUserID, ok := ctx.Value(contextKeys.UserIDKey).(uint32)
			if !ok {
				t.Fatalf("\nUserID не найден в контексте")
			}
			getDeviceID := ctx.Value(contextKeys.DeviceIDKey)
			if getDeviceID == nil {
				t.Fatalf("\nDeviceID не найден в контексте")
			}

			if validParams.userID != getUserID {
				t.Fatalf("\nUserID не совпадают: %v != %v", validParams.userID, getUserID)
			}
			if validParams.deviceID != getDeviceID {
				t.Fatalf("\nDeviceID не совпадают: %v != %v", validParams.deviceID, getDeviceID)
			}
		})
	}
}
