package endpoint

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/shopspring/decimal"

	"server/app/pkg/contextKeys"
	"server/app/pkg/datetime"
	"server/app/pkg/errors"
	"server/app/pkg/log"
	"server/app/pkg/pointer"
	"server/app/pkg/testingFunc"
	"server/app/services/account/model"
	"server/app/services/account/model/accountType"
)

func TestDecodeAccountCreateReq(t *testing.T) {

	log.Off()

	validJSON := testingFunc.NewJSONUpdater(t, `{
		"remainder": 1.1,
		"name": "name",
		"iconID": 1,
		"type": "expense",
		"currency": "USD",
		"accountGroupID": 1,
		"accountingInHeader": true,
		"accountingInCharts": true,
		"gradualBudgetFilling": true,
		"datetimeCreate": "2020-01-01T01:01:01+0100",
		"budget": {
			"amount": 1.1,
			"fixedSum": 1.1,
			"daysOffset": 1,
			"gradualFilling": true
		}
	}`)

	validWant := &model.CreateAccountReq{
		Remainder:          decimal.NewFromFloat(1.1),
		Name:               "name",
		IconID:             1,
		Type:               accountType.Expense,
		Currency:           "USD",
		AccountGroupID:     1,
		AccountingInHeader: pointer.Pointer(true),
		AccountingInCharts: pointer.Pointer(true),
		DatetimeCreate:     datetime.Time{Time: time.Date(2020, 1, 1, 1, 1, 1, 0, time.FixedZone("", 3600))},
		Budget: model.CreateAccountBudgetReq{
			Amount:         decimal.NewFromFloat(1.1),
			FixedSum:       decimal.NewFromFloat(1.1),
			DaysOffset:     1,
			GradualFilling: pointer.Pointer(true),
		},
		Necessary: testingFunc.ValidNecessary,
	}

	for _, tt := range []struct {
		message, body string
		ctx           context.Context
		want          *model.CreateAccountReq
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
		{"12.Отсутствующее поле accountingInHeader",
			validJSON.Delete("accountingInHeader").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("accountingInHeader"),
		},
		{"13.Отсутствующее поле accountingInCharts",
			validJSON.Delete("accountingInCharts").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("accountingInCharts"),
		},
		{"14.Отсутствующее поле DeviceID в контексте",
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
