package endpoint

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"server/app/internal/services/account/model"
	"server/pkg/errors"
	"server/pkg/logging"
	"server/pkg/pointer"
	"server/pkg/testingFunc"
)

func TestDecodeUpdateAccountReq(t *testing.T) {

	logging.Off()

	validJson := testingFunc.NewJSONUpdater(t, `{
		"id": 1,	
		"remainder": 1.1,
		"name": "name",	
		"iconID": 1,
		"visible": true,
		"accountGroupID": 1,
		"accounting": true,
		"budget": {
			"amount": 1.1,
			"fixedSum": 1.1,
			"daysOffset": 1,
			"gradualFilling": true
		}
	}`)

	validWant := &model.UpdateReq{
		ID:         1,
		Remainder:  pointer.Pointer(1.1),
		Name:       pointer.Pointer("name"),
		IconID:     pointer.Pointer(uint32(1)),
		Visible:    pointer.Pointer(true),
		Accounting: pointer.Pointer(true),
		UserID:     1,
		DeviceID:   "DeviceID",
		Budget: model.UpdateBudgetReq{
			Amount:         pointer.Pointer(1.1),
			FixedSum:       pointer.Pointer(1.1),
			DaysOffset:     pointer.Pointer(uint32(1)),
			GradualFilling: pointer.Pointer(true),
		},
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
				ID:       1,
				UserID:   1,
				DeviceID: "DeviceID",
			},
			nil,
		},
		{"5.Отрицательное значение iconID",
			validJson.Set("iconID", "-1").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("uint"),
		},
		{"6.Пустой запрос",
			"",
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("EOF"),
		},
		{"7.Отрицательное значение id",
			validJson.Set("id", "-1").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("uint"),
		},
		{"8.Отсутствующее поле DeviceID в контексте",
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
