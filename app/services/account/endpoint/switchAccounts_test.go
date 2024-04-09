package endpoint

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"server/app/pkg/contextKeys"
	"server/app/pkg/errors"
	"server/app/pkg/logging"
	testingFunc2 "server/app/pkg/testingFunc"
	"server/app/services/account/model"
)

func TestDecodeSwitchAccountsReq(t *testing.T) {

	logging.Off()

	validJSON := testingFunc2.NewJSONUpdater(t, `{
		"id1": 1,
		"id2": 2
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
			validJSON.Get(),
			testingFunc2.GeneralCtx.Get(),
			validWant,
			nil,
		},
		{"2.Невалидный json",
			testingFunc2.InvalidJSON,
			testingFunc2.GeneralCtx.Get(),
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
			testingFunc2.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("EOF"),
		},
		{"5.Отрицательное значение id1",
			validJSON.Set("id1", "-1").Get(),
			testingFunc2.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("uint"),
		},
		{"6.Отрицательное значение id2",
			validJSON.Set("id2", "-1").Get(),
			testingFunc2.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("uint"),
		},
		{"7.Отсутствующее поле DeviceID в контексте",
			validJSON.Get(),
			testingFunc2.GeneralCtx.Delete(contextKeys.DeviceIDKey).Get(),
			nil,
			errors.BadRequest.New("-"),
		},
		{"8.Отсутствующее поле id1",
			validJSON.Delete("id1").Get(),
			testingFunc2.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("id1"),
		},
		{"9.Отсутствующее поле id2",
			validJSON.Delete("id2").Get(),
			testingFunc2.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("id2"),
		},
	} {
		t.Run(tt.message, func(t *testing.T) {

			res, err := decodeSwitchAccountsReq(tt.ctx, httptest.NewRequest("", "/", strings.NewReader(tt.body)))
			if testingFunc2.CheckError(t, tt.err, err) {
				return
			}

			testingFunc2.CheckStruct(t, *tt.want, res, nil)
		})
	}
}
