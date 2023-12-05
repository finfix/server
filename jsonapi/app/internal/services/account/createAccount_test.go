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

	"core/app/enum/accountType"
	"jsonapi/app/internal/services/account/model"
)

func TestDecodeAccountCreateReq(t *testing.T) {

	logging.Off()

	validJson := testingFunc.NewJSONUpdater(t, `{
		"budget": 1.1,
		"remainder": 1.1,
		"name": "name",
		"iconID": 1,
		"type": "expense",
		"currency": "USD",
		"accountGroupID": 1,
		"accounting": true,
		"gradualBudgetFilling": true
	}`)

	validWant := &model.CreateReq{
		Budget:               1.1,
		Remainder:            1.1,
		Name:                 "name",
		IconID:               1,
		Type:                 accountType.Expense,
		Currency:             "USD",
		AccountGroupID:       1,
		Accounting:           pointer.Pointer(true),
		GradualBudgetFilling: pointer.Pointer(true),
		UserID:               1,
		DeviceID:             "DeviceID",
	}

	for _, tt := range []struct {
		message, body string
		ctx           context.Context
		want          *model.CreateReq
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
		{"3.Отрицательное значение на поле iconID",
			validJson.Set("iconID", -1).Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("iconID"),
		},
		{"4.Отсутствующее поле accountGroupID",
			validJson.Delete("accountGroupID").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("accountGroupID"),
		},
		{"5.Отсутствующее поле name",
			validJson.Delete("name").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("name"),
		},
		{"6.Отсутствующее поле iconID",
			validJson.Delete("iconID").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("iconID"),
		},
		{"7.Отсутствующее поле type",
			validJson.Delete("type").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("type"),
		},
		{"8.С невалидным полем type",
			validJson.Set("type", "invalid").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("account type"),
		},
		{"9.Отсутствующее поле currency",
			validJson.Delete("currency").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("currency"),
		},
		{"10.Отсутствующее поле UserID в контексте",
			validJson.Get(),
			testingFunc.GeneralCtx.Delete("DeviceID").Get(),
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
			validJson.Delete("accounting").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("accounting"),
		},
		{"13.Отсутствующее поле DeviceID в контексте",
			validJson.Get(),
			testingFunc.GeneralCtx.Delete("DeviceID").Get(),
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
