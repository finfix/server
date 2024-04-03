package endpoint

import (
	"context"
	"fmt"
	"net/http/httptest"
	"net/url"
	"testing"

	"server/app/enum/accountType"
	"server/app/internal/services/account/model"
	"server/pkg/errors"
	"server/pkg/logging"
	"server/pkg/pointer"
	"server/pkg/testingFunc"
)

func TestDecodeGetAccountsReq(t *testing.T) {

	logging.Off()

	validParams := testingFunc.NewParamUpdater(map[string]string{
		"type":            "expense",
		"accountGroupIDs": "1,2",
		"accounting":      "true",
	})

	validWant := &model.GetReq{
		Type:            pointer.Pointer(accountType.Expense),
		AccountGroupIDs: []uint32{1, 2},
		Accounting:      pointer.Pointer(true),
		UserID:          1,
		DeviceID:        "DeviceID",
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
		{"2.Невалидное поле type",
			validParams.Set("type", "invalid").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("type"),
		},
		{"3.Отрицательное значение поля accountGroupID",
			validParams.Set("accountGroupID", "-1").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("accountGroupID"),
		},
		{"4.Отсутствующее поле DeviceID в контексте",
			validParams.Get(),
			testingFunc.GeneralCtx.Delete("DeviceID").Get(),
			nil,
			errors.BadRequest.New("-"),
		},
		{"6.Пустой запрос",
			nil,
			testingFunc.GeneralCtx.Get(),
			&model.GetReq{
				UserID:   1,
				DeviceID: "DeviceID",
			},
			nil,
		},
		{"7.Отсутствующее поле UserID в контексте",
			validParams.Get(),
			testingFunc.GeneralCtx.Delete("UserID").Get(),
			nil,
			errors.BadRequest.New("-"),
		},
	} {
		t.Run(tt.message, func(t *testing.T) {

			res, err := decodeGetReq(tt.ctx, httptest.NewRequest("GET", fmt.Sprintf("%s?%s", "/account", tt.params.Encode()), nil))
			if testingFunc.CheckError(t, tt.err, err) {
				return
			}

			testingFunc.CheckStruct(t, *tt.want, res, nil)
		})
	}
}
