package endpoint

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"server/app/internal/services/auth/model"
	"server/pkg/contextKeys"
	"server/pkg/errors"
	"server/pkg/logging"
	"server/pkg/testingFunc"
)

func TestDecodeAuthReq(t *testing.T) {

	logging.Off()

	validJson := testingFunc.NewJSONUpdater(t, `{
		"email": "qwerty@berubox.com",
		"password": "password"
	}`)

	for _, tt := range []struct {
		message, body string
		ctx           context.Context
		want          *model.SignInReq
		err           error
	}{
		{"1.Обычный запрос",
			validJson.Get(),
			testingFunc.GeneralCtx.Get(),
			&model.SignInReq{
				Email:    "qwerty@berubox.com",
				Password: "password",
				DeviceID: "DeviceID",
			},
			nil,
		},
		{"2.Невалидный json",
			testingFunc.InvalidJSON,
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("invalid"),
		},
		{"3.Невалидный email",
			validJson.Set("email", "invalid").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("email"),
		},
		{"4.Отсутствующее поле email",
			validJson.Delete("email").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("email"),
		},
		{"5.Отсутствующее поле password",
			validJson.Delete("password").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("password"),
		},
		{"6.Пустой запрос",
			"",
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("EOF"),
		},
		{"7.Запрос с пустым полем DeviceID в контексте",
			validJson.Get(),
			testingFunc.GeneralCtx.Delete(contextKeys.DeviceIDKey).Get(),
			nil,
			errors.BadRequest.New("-"),
		},
	} {
		t.Run(tt.message, func(t *testing.T) {

			res, err := decodeSignInReq(tt.ctx, httptest.NewRequest("", "/", strings.NewReader(tt.body)))
			if testingFunc.CheckError(t, tt.err, err) {
				return
			}

			testingFunc.CheckStruct(t, *tt.want, res, nil)
		})
	}
}
