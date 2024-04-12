package endpoint

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"server/app/pkg/contextKeys"
	"server/app/pkg/errors"
	"server/app/pkg/logging"
	"server/app/pkg/testingFunc"
	"server/app/services/auth/model"
)

func TestDecodeSignUp(t *testing.T) {

	logging.Off()

	validJSON := testingFunc.NewJSONUpdater(t, `{
		"email": "qwerty@berubox.com",
		"password": "password",
		"name": "name"
	}`)

	validWant := &model.SignUpReq{
		Email:    "qwerty@berubox.com",
		Password: "password",
		Name:     "name",
		DeviceID: "DeviceID",
	}

	for _, tt := range []struct {
		message, body string
		ctx           context.Context
		want          *model.SignUpReq
		err           error
	}{
		{"1.Обычный запрос",
			validJSON.Get(),
			testingFunc.GeneralCtx.Get(),
			validWant,
			nil,
		},
		{"2.Невалидный json",
			testingFunc.InvalidJSON,
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("invalid"),
		},
		{"3.Невалидный email",
			validJSON.Set("email", "invalid").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("email"),
		},
		{"4.Отсутствующее поле email",
			validJSON.Delete("email").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("email"),
		},
		{"5.Отсутствующее поле password",
			validJSON.Delete("password").Get(),
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
		{"7.Отсутствующее поле name",
			validJSON.Delete("name").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("name"),
		},
		{"8.Отсутствующее поле DeviceID в контексте",
			validJSON.Get(),
			testingFunc.GeneralCtx.Delete(contextKeys.DeviceIDKey).Get(),
			nil,
			errors.BadRequest.New("-"),
		},
	} {
		t.Run(tt.message, func(t *testing.T) {

			res, err := decodeSignUpReq(tt.ctx, httptest.NewRequest("", "/", strings.NewReader(tt.body)))
			if testingFunc.CheckError(t, tt.err, err) {
				return
			}

			testingFunc.CheckStruct(t, *tt.want, res, nil)
		})
	}
}
