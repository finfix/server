package endpoint

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/shopspring/decimal"

	"server/app/pkg/contextKeys"
	"server/app/pkg/errors"
	"server/app/pkg/log"
	"server/app/pkg/pointer"
	"server/app/pkg/testingFunc"
	"server/app/services/account/model"
)

func TestDecodeUpdateAccountReq(t *testing.T) {

	log.Off()

	validJSON := testingFunc.NewJSONUpdater(t, `{
		"id": 1,	
		"remainder": 1.1,
		"name": "name",	
		"iconID": 1,
		"visible": true,
		"accountGroupID": 1,
		"accountingInHeader": true,
		"accountingInCharts": true,
		"budget": {
			"amount": 1.1,
			"fixedSum": 1.1,
			"daysOffset": 1,
			"gradualFilling": true
		}
	}`)

	validWant := &model.UpdateAccountReq{
		Necessary:          testingFunc.ValidNecessary,
		ID:                 1,
		Remainder:          pointer.Pointer(decimal.NewFromFloat(1.1)),
		Name:               pointer.Pointer("name"),
		IconID:             pointer.Pointer(uint32(1)),
		Visible:            pointer.Pointer(true),
		AccountingInHeader: pointer.Pointer(true),
		AccountingInCharts: pointer.Pointer(true),
		Currency:           nil,
		ParentAccountID:    nil,
		Budget: model.UpdateAccountBudgetReq{
			Amount:         pointer.Pointer(decimal.NewFromFloat(1.1)),
			FixedSum:       pointer.Pointer(decimal.NewFromFloat(1.1)),
			DaysOffset:     pointer.Pointer(uint32(1)),
			GradualFilling: pointer.Pointer(true),
		},
	}

	for _, tt := range []struct {
		message, body string
		ctx           context.Context
		want          *model.UpdateAccountReq
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
		{"3.Отсутствующее поле UserID в контексте",
			validJSON.Get(),
			testingFunc.GeneralCtx.Delete(contextKeys.UserIDKey).Get(),
			nil,
			errors.BadRequest.New("-"),
		},
		{"4.Минимальный запрос", `{
				"id": 1,
				"accountGroupID": 1
			}`,
			testingFunc.GeneralCtx.Get(),
			&model.UpdateAccountReq{
				Necessary:          testingFunc.ValidNecessary,
				ID:                 1,
				Remainder:          nil,
				Name:               nil,
				IconID:             nil,
				Visible:            nil,
				AccountingInHeader: nil,
				AccountingInCharts: nil,
				Currency:           nil,
				ParentAccountID:    nil,
				Budget: model.UpdateAccountBudgetReq{
					Amount:         nil,
					FixedSum:       nil,
					DaysOffset:     nil,
					GradualFilling: nil,
				},
			},
			nil,
		},
		{"5.Отрицательное значение iconID",
			validJSON.Set("iconID", "-1").Get(),
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
			validJSON.Set("id", "-1").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("uint"),
		},
		{"8.Отсутствующее поле DeviceID в контексте",
			validJSON.Get(),
			testingFunc.GeneralCtx.Delete(contextKeys.DeviceIDKey).Get(),
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
