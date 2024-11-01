package chain

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"

	"pkg/contextKeys"
	"pkg/errors"
	"pkg/jwtManager"
	"pkg/pointer"
	"pkg/testUtils"
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

	for _, tt := range []struct {
		message        string
		token          *string
		params         *createTokenParams
		err            error
		needCompareErr bool
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
			false,
		},
		{
			"2.Токен с истекшим сроком",
			nil,
			&createTokenParams{
				userID:   validParams.userID,
				deviceID: validParams.deviceID,
				ttl:      -time.Hour,
			},
			errors.Unauthorized.Wrap(jwtManager.ErrUserUnauthorized),
			true,
		},
		{
			"3.Невалидный токен",
			pointer.Pointer("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUyNTg5ODMyMTEsInN1YiI6IjEifQ.TneMNhesU3VT0XVGb8dGg8zyyObrmPk_x9kdh-aJDwQ"),
			&createTokenParams{
				userID:   0,
				deviceID: "",
				ttl:      -time.Hour,
			},
			errors.Unauthorized.Wrap(jwt.ErrSignatureInvalid),
			false,
		},
		{
			"4.Пустой токен",
			pointer.Pointer(""),
			&createTokenParams{
				userID:   0,
				deviceID: "",
				ttl:      time.Hour,
			},
			errors.Unauthorized.New("JWT-token is empty"),
			false,
		},
		{
			"5.Токен без DeviceID",
			nil,
			&createTokenParams{
				userID:   validParams.userID,
				ttl:      validParams.ttl,
				deviceID: "",
			},
			errors.Unauthorized.New("DeviceID is empty"),
			false,
		},
		{
			"6.Токен без UserID",
			nil,
			&createTokenParams{
				deviceID: validParams.deviceID,
				ttl:      validParams.ttl,
				userID:   0,
			},
			errors.Unauthorized.New("UserID is empty"),
			false,
		},
	} {
		t.Run(tt.message, func(t *testing.T) {

			tt := tt

			jwtManager.Init([]byte("test"), tt.params.ttl, 0)

			// Если токен не передан, то создаем его
			if tt.token == nil {

				// Создаем токен
				token, err := jwtManager.GenerateToken(jwtManager.AccessToken, tt.params.userID, tt.params.deviceID)
				if err != nil {
					t.Fatalf("\nНе смогли создать JWV-токен: %v", err)
				}
				tt.token = &token
			}

			req := httptest.NewRequest("", "/", nil)
			req.Header.Add("Authorization", *tt.token)

			ctx, err := DefaultAuthorization(context.Background(), req)

			if testUtils.CheckError(t, tt.err, err, tt.needCompareErr) {
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
