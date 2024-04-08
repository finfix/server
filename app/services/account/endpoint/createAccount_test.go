package endpoint

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"server/app/services/account/model"
	"server/app/services/account/model/accountType"
	"server/pkg/contextKeys"
	"server/pkg/errors"
	"server/pkg/logging"
	"server/pkg/pointer"
	"server/pkg/testingFunc"
)

func TestDecodeAccountCreateReq(t *testing.T) {

	logging.Off()

	validJSON := testingFunc.NewJSONUpdater(t, `{
		"remainder": 1.1,
		"name": "name",
		"iconID": 1,
		"type": "expense",
		"currency": "USD",
		"accountGroupID": 1,
		"accounting": true,
		"gradualBudgetFilling": true,
		"budget": {
			"amount": 1.1,
			"fixedSum": 1.1,
			"daysOffset": 1,
			"gradualFilling": true
		}
	}`)

	validWant := &model.CreateReq{
		Remainder:      1.1,
		Name:           "name",
		IconID:         1,
		Type:           accountType.Expense,
		Currency:       "USD",
		AccountGroupID: 1,
		Accounting:     pointer.Pointer(true),
		Budget: model.CreateBudgetReq{
			Amount:         1.1,
			FixedSum:       1.1,
			DaysOffset:     1,
			GradualFilling: pointer.Pointer(true),
		},
		UserID:   1,
		DeviceID: "DeviceID",
	}

	for _, tt := range []struct {
		message, body string
		ctx           context.Context
		want          *model.CreateReq
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
		{"3.Отрицательное значение на поле iconID",
			validJSON.Set("iconID", -1).Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("iconID"),
		},
		{"4.Отсутствующее поле accountGroupID",
			validJSON.Delete("accountGroupID").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("accountGroupID"),
		},
		{"5.Отсутствующее поле name",
			validJSON.Delete("name").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("name"),
		},
		{"6.Отсутствующее поле iconID",
			validJSON.Delete("iconID").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("iconID"),
		},
		{"7.Отсутствующее поле type",
			validJSON.Delete("type").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("type"),
		},
		{"8.С невалидным полем type",
			validJSON.Set("type", "invalid").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("account type"),
		},
		{"9.Отсутствующее поле currency",
			validJSON.Delete("currency").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("currency"),
		},
		{"10.Отсутствующее поле UserID в контексте",
			validJSON.Get(),
			testingFunc.GeneralCtx.Delete(contextKeys.DeviceIDKey).Get(),
			nil,
			errors.BadRequest.New("-"),
		},
		{"11.Пустой запрос",
			"",
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("EOF"),
		},
		{"12.Отсутствующее поле accounting",
			validJSON.Delete("accounting").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("accounting"),
		},
		{"13.Отсутствующее поле DeviceID в контексте",
			validJSON.Get(),
			testingFunc.GeneralCtx.Delete(contextKeys.DeviceIDKey).Get(),
			nil,
			errors.BadRequest.New("-"),
		},
	} {
		t.Run(tt.message, func(t *testing.T) {

			res, err := decodeCreateAccountReq(tt.ctx, httptest.NewRequest("", "/", strings.NewReader(tt.body)))
			if testingFunc.CheckError(t, tt.err, err) {
				return
			}

			testingFunc.CheckStruct(t, *tt.want, res, nil)
		})
	}
}
