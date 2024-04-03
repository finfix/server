package endpoint

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"server/app/internal/services/auth/model"
	"server/pkg/errors"
	"server/pkg/logging"
	"server/pkg/testingFunc"
)

func TestDecodeSignUp(t *testing.T) {

	logging.Off()

	validJson := testingFunc.NewJSONUpdater(t, `{
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
			validJson.Get(),
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
		{"7.Отсутствующее поле name",
			validJson.Delete("name").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("name"),
		},
		{"8.Отсутствующее поле DeviceID в контексте",
			validJson.Get(),
			testingFunc.GeneralCtx.Delete("DeviceID").Get(),
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
