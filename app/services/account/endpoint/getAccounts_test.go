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
	testingFunc2 "server/app/pkg/testingFunc"
	"server/app/services/account/model"
	"server/app/services/account/model/accountType"
)

func TestDecodeGetAccountsReq(t *testing.T) {

	logging.Off()

	validParams := testingFunc2.NewParamUpdater(map[string]string{
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
			testingFunc2.GeneralCtx.Get(),
			validWant,
			nil,
		},
		{"2.Невалидное поле type",
			validParams.Set("type", "invalid").Get(),
			testingFunc2.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("type"),
		},
		{"3.Отрицательное значение поля accountGroupID",
			validParams.Set("accountGroupID", "-1").Get(),
			testingFunc2.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("accountGroupID"),
		},
		{"4.Отсутствующее поле DeviceID в контексте",
			validParams.Get(),
			testingFunc2.GeneralCtx.Delete("DeviceID").Get(),
			nil,
			errors.BadRequest.New("-"),
		},
		{"6.Пустой запрос",
			nil,
			testingFunc2.GeneralCtx.Get(),
			&model.GetReq{
				UserID:   1,
				DeviceID: "DeviceID",
			},
			nil,
		},
		{"7.Отсутствующее поле UserID в контексте",
			validParams.Get(),
			testingFunc2.GeneralCtx.Delete(contextKeys.UserIDKey).Get(),
			nil,
			errors.BadRequest.New("-"),
		},
	} {
		t.Run(tt.message, func(t *testing.T) {

			res, err := decodeGetReq(tt.ctx, httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s?%s", part, tt.params.Encode()), nil))
			if testingFunc2.CheckError(t, tt.err, err) {
				return
			}

			testingFunc2.CheckStruct(t, *tt.want, res, nil)
		})
	}
}
