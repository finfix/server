package account

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"logger/app/logging"
	"pkg/errors"
	"pkg/pointer"
	"pkg/testingFunc"

	"jsonapi/app/internal/services/account/model"
)

func TestDecodeUpdateAccountReq(t *testing.T) {

	logging.Off()

	validJson := testingFunc.NewJSONUpdater(t, `{
		"id": 1,
		"budget": 1,	
		"remainder": 1.1,
		"name": "name",	
		"iconID": 1,
		"visible": true,
		"accountGroupID": 1,
		"accounting": true
	}`)

	validWant := &model.UpdateReq{
		ID:             1,
		Budget:         pointer.Pointer(int32(1)),
		Remainder:      pointer.Pointer(1.1),
		Name:           pointer.Pointer("name"),
		IconID:         pointer.Pointer(uint32(1)),
		Visible:        pointer.Pointer(true),
		AccountGroupID: pointer.Pointer(uint32(1)),
		Accounting:     pointer.Pointer(true),
		UserID:         1,
		DeviceID:       "DeviceID",
	}

	for _, tt := range []struct {
		message, body string
		ctx           context.Context
		want          *model.UpdateReq
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
		{"3.Отсутствующее поле UserID в контексте",
			validJson.Get(),
			testingFunc.GeneralCtx.Delete("UserID").Get(),
			nil,
			errors.BadRequest.New("-"),
		},
		{"4.Минимальный запрос", `{
				"id": 1,
				"accountGroupID": 1
			}`,
			testingFunc.GeneralCtx.Get(),
			&model.UpdateReq{
				ID:             1,
				AccountGroupID: pointer.Pointer(uint32(1)),
				UserID:         1,
				DeviceID:       "DeviceID",
			},
			nil,
		},
		{"5.Отрицательное значение iconID",
			validJson.Set("iconID", "-1").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("uint"),
		},
		{"6.Отрицательное значение accountGroupID",
			validJson.Set("accountGroupID", "-1").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("uint"),
		},
		{"7.Пустой запрос",
			"",
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("EOF"),
		},
		{"8.Отрицательное значение id",
			validJson.Set("id", "-1").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("uint"),
		},
		{"9.Отсутствующее поле DeviceID в контексте",
			validJson.Get(),
			testingFunc.GeneralCtx.Delete("DeviceID").Get(),
			nil,
			errors.BadRequest.New("-"),
		},
	} {
		t.Run(tt.message, func(t *testing.T) {

			res, err := decodeUpdateAccountReq(tt.ctx, httptest.NewRequest("", "/", strings.NewReader(tt.body)))
			if testingFunc.CheckError(t, tt.err, err) {
				return
			}

			testingFunc.CheckStruct(t, *tt.want, res, nil)
		})
	}
}
