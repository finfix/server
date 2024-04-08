package endpoint

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"server/app/services/transaction/model"
	"server/pkg/contextKeys"
	"server/pkg/datetime/date"
	"server/pkg/errors"
	"server/pkg/logging"
	"server/pkg/pointer"
	"server/pkg/testingFunc"
)

func TestDecodeCreateReq(t *testing.T) {

	logging.Off()

	validJSON := testingFunc.NewJSONUpdater(t, `{
		"type": "income",
		"amountFrom": 1.1,
		"amountTo": 1.1,
		"note": "name",
		"accountFromID": 1,
		"accountToID": 1,
		"dateTransaction": "2020-01-01",
		"isExecuted": true
	}`)

	validWant := &model.CreateReq{
		Type:            "income",
		AmountFrom:      1.1,
		AmountTo:        1.1,
		Note:            "name",
		AccountFromID:   1,
		AccountToID:     1,
		DateTransaction: date.NewDate(2020, 1, 1),
		IsExecuted:      pointer.Pointer(true),
		UserID:          1,
		DeviceID:        "DeviceID",
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
		{"3.Отрицательное значение поля amountFrom",
			validJSON.Set("amountFrom", -1.1).Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("amountFrom"),
		},
		{"4.Отрицательное значение поля amountTo",
			validJSON.Set("amountTo", -1.1).Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("amountTo"),
		},
		{"5.Отрицательное значение поля accountFromID",
			validJSON.Set("accountFromID", -1).Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("accountFromID"),
		},
		{"6.Отрицательное значение поля accountToID",
			validJSON.Set("accountToID", -1).Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("accountToID"),
		},
		{"7.Невалидная дата",
			validJSON.Set("dateTransaction", "invalid").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("time"),
		},
		{"8.Невалидный тип транзакции",
			validJSON.Set("type", "invalid").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("transaction type"),
		},
		{"9.Отсутствующее поле type",
			validJSON.Delete("type").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("type"),
		},
		{"10.Отсутствующее поле amountFrom",
			validJSON.Delete("amountFrom").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("amountFrom"),
		},
		{"11.Отсутствующее поле amountTo",
			validJSON.Delete("amountTo").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("amountTo"),
		},
		{"12.Отсутствующее поле accountFromID",
			validJSON.Delete("accountFromID").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("accountFromID"),
		},
		{"13.Отсутствующее поле accountToID",
			validJSON.Delete("accountToID").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("accountToID"),
		},
		{"14.Отсутствующее поле dateTransaction",
			validJSON.Delete("dateTransaction").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("date"),
		},
		{"15.Отсутствующее поле isExecuted",
			validJSON.Delete("isExecuted").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("isExecuted"),
		},
		{"16.Отсутствующее поле UserID в контексте",
			validJSON.Get(),
			testingFunc.GeneralCtx.Delete(contextKeys.UserIDKey).Get(),
			nil,
			errors.BadRequest.New(""),
		},
		{"17.Пустой запрос",
			"",
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("EOF"),
		},
		{"18.Отсутствующее поле DeviceID в контексте",
			validJSON.Get(),
			testingFunc.GeneralCtx.Delete(contextKeys.DeviceIDKey).Get(),
			nil,
			errors.BadRequest.New(""),
		},
	} {
		t.Run(tt.message, func(t *testing.T) {

			res, err := decodeCreateTransactionReq(tt.ctx, httptest.NewRequest("", "/", strings.NewReader(tt.body)))
			if testingFunc.CheckError(t, tt.err, err) {
				return
			}

			testingFunc.CheckStruct(t, *tt.want, res, nil)
		})
	}
}
