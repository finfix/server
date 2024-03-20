package endpoint

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"server/pkg/errors"
	"server/pkg/logging"
	"server/pkg/testingFunc"

	"server/app/internal/services/account/model"
)

func TestDecodeSwitchAccountsReq(t *testing.T) {

	logging.Off()

	validJson := testingFunc.NewJSONUpdater(t, `{
		"id_1": 1,
		"id_2": 2
	}`)

	validWant := &model.SwitchReq{
		ID1:      1,
		ID2:      2,
		UserID:   1,
		DeviceID: "DeviceID",
	}

	for _, tt := range []struct {
		message, body string
		ctx           context.Context
		want          *model.SwitchReq
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
		{"3.Пустой json",
			`{}`,
			context.Background(),
			nil,
			errors.BadRequest.New("id"),
		},
		{"4.Пустой запрос",
			``,
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("EOF"),
		},
		{"5.Отрицательное значение id_1",
			validJson.Set("id_1", "-1").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("uint"),
		},
		{"6.Отрицательное значение id_2",
			validJson.Set("id_2", "-1").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("uint"),
		},
		{"7.Отсутствующее поле DeviceID в контексте",
			validJson.Get(),
			testingFunc.GeneralCtx.Delete("DeviceID").Get(),
			nil,
			errors.BadRequest.New("-"),
		},
		{"8.Отсутствующее поле id_1",
			validJson.Delete("id_1").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("id_1"),
		},
		{"9.Отсутствующее поле id_2",
			validJson.Delete("id_2").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("id_2"),
		},
	} {
		t.Run(tt.message, func(t *testing.T) {

			res, err := decodeSwitchAccountsReq(tt.ctx, httptest.NewRequest("", "/", strings.NewReader(tt.body)))
			if testingFunc.CheckError(t, tt.err, err) {
				return
			}

			testingFunc.CheckStruct(t, *tt.want, res, nil)
		})
	}
}
