package transaction

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"logger/app/logging"
	"pkg/datetime/date"
	"pkg/errors"
	"pkg/pointer"
	"pkg/testingFunc"

	"jsonapi/app/internal/services/transaction/model"
)

func TestDecodeUpdateReq(t *testing.T) {

	logging.Off()

	validJson := testingFunc.NewJSONUpdater(t, `{
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
				"id": 1
			}`,
			testingFunc.GeneralCtx.Get(),
			&model.UpdateReq{
				ID:       1,
				UserID:   1,
				DeviceID: "DeviceID",
			},
			nil,
		},
		{"5.Отрицательное значение id",
			validJson.Set("id", "-1").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("uint"),
		},
		{"6.Отрицательное значение amountFrom",
			validJson.Set("amountFrom", "-1.1").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("amountFrom"),
		},
		{"7.Отрицательное значение amountTo",
			validJson.Set("amountTo", "-1.1").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("amountTo"),
		},
		{"8.Отрицательное значение accountFromID",
			validJson.Set("accountFromID", "-1").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("accountFromID"),
		},
		{"9.Отрицательное значение accountToID",
			validJson.Set("accountToID", "-1").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("accountToID"),
		},
		{"10.Невалидная дата",
			validJson.Set("dateTransaction", "invalid").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("time"),
		},
		{"11.Пустой запрос",
			"",
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("EOF"),
		},
		{"12.Отсутствующее поле DeviceID в контексте",
			validJson.Get(),
			testingFunc.GeneralCtx.Delete("DeviceID").Get(),
			nil,
			errors.BadRequest.New("-"),
		},
	} {
		t.Run(tt.message, func(t *testing.T) {

			res, err := decodeUpdateTransactionReq(tt.ctx, httptest.NewRequest("", "/", strings.NewReader(tt.body)))
			if testingFunc.CheckError(t, tt.err, err) {
				return
			}

			testingFunc.CheckStruct(t, *tt.want, res, nil)
		})
	}
}
