package endpoint

import (
	"context"
	"fmt"
	"net/http/httptest"
	"net/url"
	"testing"

	"server/app/enum/transactionType"
	"server/app/internal/services/transaction/model"
	"server/pkg/contextKeys"
	"server/pkg/datetime/date"
	"server/pkg/errors"
	"server/pkg/logging"
	"server/pkg/pointer"
	"server/pkg/testingFunc"
)

func TestDecodeGetReq(t *testing.T) {

	logging.Off()

	validParams := testingFunc.NewParamUpdater(map[string]string{
		"offset":    "1",
		"limit":     "100",
		"accountID": "1",
		"type":      "income",
		"dateFrom":  "2020-01-01",
		"dateTo":    "2020-01-02",
	})

	validWant := &model.GetReq{
		Offset:    pointer.Pointer(uint32(1)),
		Limit:     pointer.Pointer(uint32(100)),
		AccountID: pointer.Pointer(uint32(1)),
		Type:      pointer.Pointer(transactionType.Income),
		DateFrom:  pointer.Pointer(date.NewDate(2020, 1, 1)),
		DateTo:    pointer.Pointer(date.NewDate(2020, 1, 2)),
		UserID:    1,
		DeviceID:  "DeviceID",
	}

	for _, tt := range []struct {
		message string
		params  url.Values
		ctx     context.Context
		want    *model.GetReq
		err     error
	}{
		{"1.Обычный запрос",
			validParams.Get(),
			testingFunc.GeneralCtx.Get(),
			validWant,
			nil,
		},
		{"2.Отсутствующее поле UserID в контексте",
			validParams.Get(),
			testingFunc.GeneralCtx.Delete(contextKeys.UserIDKey).Get(),
			nil,
			errors.BadRequest.New("-"),
		},
		{"3.Отсутствующее поле DeviceID в контексте",
			validParams.Get(),
			testingFunc.GeneralCtx.Delete(contextKeys.DeviceIDKey).Get(),
			nil,
			errors.BadRequest.New("-"),
		},
		{"4.Пустой запрос",
			nil,
			testingFunc.GeneralCtx.Get(),
			&model.GetReq{
				UserID:   1,
				DeviceID: "DeviceID",
			},
			nil,
		},
		{"5.Невалидный тип транзакции",
			validParams.Set("type", "invalid").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("transaction type"),
		},
		{"6.Отрицательное поле list",
			validParams.Set("list", "-1").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("list"),
		},
		{"7.Отрицательное поле accountID",
			validParams.Set("accountID", "-1").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("accountID"),
		},
		{"8.Невалидная дата dateFrom",
			validParams.Set("dateFrom", "invalid").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("date"),
		},
		{"9.Невалидная дата dateTo",
			validParams.Set("dateTo", "invalid").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("date"),
		},
		{"10.Вторая дата раньше первой",
			validParams.Set("dateFrom", "2020-01-02").Set("dateTo", "2020-01-01").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("less"),
		},
		{"11.Вторая дата равна первой",
			validParams.Set("dateFrom", "2020-01-02").Set("dateTo", "2020-01-02").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("less"),
		},
	} {
		t.Run(tt.message, func(t *testing.T) {

			res, err := decodeGetTransactionsReq(tt.ctx, httptest.NewRequest("GET", fmt.Sprintf("%s?%s", "/transaction", tt.params.Encode()), nil))
			if testingFunc.CheckError(t, tt.err, err) {
				return
			}

			testingFunc.CheckStruct(t, *tt.want, res, nil)
		})
	}
}
