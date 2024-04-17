package endpoint

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"server/app/pkg/contextKeys"
	"server/app/pkg/errors"
	"server/app/pkg/logging"
	"server/app/pkg/pointer"
	"server/app/pkg/testingFunc"
	"server/app/services/account/model"
	"server/app/services/account/model/accountType"
)

func TestDecodeGetAccountsReq(t *testing.T) {

	logging.Off()

	validParams := testingFunc.NewParamUpdater(map[string]string{
		"type":               "expense",
		"accountGroupIDs":    "1,2",
		"accountingInHeader": "true",
	})

	validWant := &model.GetAccountsReq{
		Type:               pointer.Pointer(accountType.Expense),
		AccountGroupIDs:    []uint32{1, 2},
		AccountingInHeader: pointer.Pointer(true),
		Necessary:          testingFunc.ValidNecessary,
	}

	for _, tt := range []struct {
		message string
		params  url.Values
		ctx     context.Context
		want    *model.GetAccountsReq
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
			&model.GetAccountsReq{
				Necessary: testingFunc.ValidNecessary,
			},
			nil,
		},
		{"7.Отсутствующее поле UserID в контексте",
			validParams.Get(),
			testingFunc.GeneralCtx.Delete(contextKeys.UserIDKey).Get(),
			nil,
			errors.BadRequest.New("-"),
		},
	} {
		t.Run(tt.message, func(t *testing.T) {

			res, err := decodeGetReq(tt.ctx, httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s?%s", part, tt.params.Encode()), nil))
			if testingFunc.CheckError(t, tt.err, err) {
				return
			}

			testingFunc.CheckStruct(t, *tt.want, res, nil)
		})
	}
}
