package endpoint

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"server/app/pkg/contextKeys"
	"server/app/pkg/datetime/date"
	"server/app/pkg/errors"
	"server/app/pkg/logging"
	"server/app/pkg/pointer"
	testingFunc2 "server/app/pkg/testingFunc"
	"server/app/services/transaction/model"
)

func TestDecodeUpdateReq(t *testing.T) {

	logging.Off()

	validJSON := testingFunc2.NewJSONUpdater(t, `{
		"id": 1,
		"amountFrom": 1.1,
		"amountTo": 1.1,
		"note": "name",
		"accountFromID": 1,
		"accountToID": 1,
		"dateTransaction": "2020-01-01",
		"isExecuted": true
	}`)

	validWant := &model.UpdateReq{
		ID:              1,
		AmountFrom:      pointer.Pointer(1.1),
		AmountTo:        pointer.Pointer(1.1),
		Note:            pointer.Pointer("name"),
		AccountFromID:   pointer.Pointer(uint32(1)),
		AccountToID:     pointer.Pointer(uint32(1)),
		DateTransaction: pointer.Pointer(date.NewDate(2020, 1, 1)),
		IsExecuted:      pointer.Pointer(true),
		UserID:          1,
		DeviceID:        "DeviceID",
	}

	for _, tt := range []struct {
		message, body string
		ctx           context.Context
		want          *model.UpdateReq
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
		{"3.Отсутствующее поле UserID в контексте",
			validJSON.Get(),
			testingFunc2.GeneralCtx.Delete(contextKeys.UserIDKey).Get(),
			nil,
			errors.BadRequest.New("-"),
		},
		{"4.Минимальный запрос", `{
				"id": 1
			}`,
			testingFunc2.GeneralCtx.Get(),
			&model.UpdateReq{
				ID:       1,
				UserID:   1,
				DeviceID: "DeviceID",
			},
			nil,
		},
		{"5.Отрицательное значение id",
			validJSON.Set("id", "-1").Get(),
			testingFunc2.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("uint"),
		},
		{"6.Отрицательное значение amountFrom",
			validJSON.Set("amountFrom", "-1.1").Get(),
			testingFunc2.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("amountFrom"),
		},
		{"7.Отрицательное значение amountTo",
			validJSON.Set("amountTo", "-1.1").Get(),
			testingFunc2.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("amountTo"),
		},
		{"8.Отрицательное значение accountFromID",
			validJSON.Set("accountFromID", "-1").Get(),
			testingFunc2.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("accountFromID"),
		},
		{"9.Отрицательное значение accountToID",
			validJSON.Set("accountToID", "-1").Get(),
			testingFunc2.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("accountToID"),
		},
		{"10.Невалидная дата",
			validJSON.Set("dateTransaction", "invalid").Get(),
			testingFunc2.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("time"),
		},
		{"11.Пустой запрос",
			"",
			testingFunc2.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("EOF"),
		},
		{"12.Отсутствующее поле DeviceID в контексте",
			validJSON.Get(),
			testingFunc2.GeneralCtx.Delete(contextKeys.DeviceIDKey).Get(),
			nil,
			errors.BadRequest.New("-"),
		},
	} {
		t.Run(tt.message, func(t *testing.T) {

			res, err := decodeUpdateTransactionReq(tt.ctx, httptest.NewRequest("", "/", strings.NewReader(tt.body)))
			if testingFunc2.CheckError(t, tt.err, err) {
				return
			}

			testingFunc2.CheckStruct(t, *tt.want, res, nil)
		})
	}
}
