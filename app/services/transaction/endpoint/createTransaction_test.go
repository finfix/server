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
	"server/app/services/transaction/model"
)

func TestDecodeCreateReq(t *testing.T) {

	log.Off()

	validJSON := testingFunc.NewJSONUpdater(t, `{
		"type": "income",
		"amountFrom": 1.1,
		"amountTo": 1.1,
		"note": "name",
		"accountFromID": 1,
		"accountToID": 1,
		"dateTransaction": "2020-01-01",
		"isExecuted": true,
		"datetimeCreate": "2020-01-01T01:01:01+0100"
	}`)

	validWant := &model.CreateTransactionReq{
		Type:            "income",
		AmountFrom:      decimal.NewFromFloat(1.1),
		AmountTo:        decimal.NewFromFloat(1.1),
		Note:            "name",
		AccountFromID:   1,
		AccountToID:     1,
		DateTransaction: datetime.NewDate(2020, 1, 1),
		IsExecuted:      pointer.Pointer(true),
		Necessary:       testingFunc.ValidNecessary,
		DatetimeCreate:  datetime.Time{Time: time.Date(2020, 1, 1, 1, 1, 1, 0, time.FixedZone("", 3600))},
		TagIDs:          nil, // TODO: Проверить
	}

	for _, tt := range []struct {
		message, body string
		ctx           context.Context
		want          *model.CreateTransactionReq
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
