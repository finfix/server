package account

import (
	"context"
	"fmt"
	"net/http/httptest"
	"net/url"
	"testing"

	"logger/app/logging"
	"pkg/errors"
	"pkg/testingFunc"

	"jsonapi/app/internal/services/account/model"
)

func TestDecodeDeleteAccountReq(t *testing.T) {

	logging.Off()

	validParams := testingFunc.NewParamUpdater(map[string]string{
		"id": "1",
	})

	validWant := &model.DeleteReq{
		ID:       1,
		UserID:   1,
		DeviceID: "DeviceID",
	}

	for _, tt := range []struct {
		message string
		params  url.Values
		ctx     context.Context
		want    *model.DeleteReq
		err     error
	}{
		{"1.Обычный запрос",
			validParams.Get(),
			testingFunc.GeneralCtx.Get(),
			validWant,
			nil,
		},
		{"2.Пустой запрос",
			testingFunc.NewParamUpdater(map[string]string{}).Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("EOF"),
		},
		{"3.Отрицательное значение id",
			validParams.Set("id", "-1").Get(),
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("uint"),
		},
		{"4.Отсутствующее поле UserID в контексте",
			validParams.Get(),
			testingFunc.GeneralCtx.Delete("UserID").Get(),
			nil,
			errors.BadRequest.New("-"),
		},
		{"5.Отсутствующее поле DeviceID в контексте",
			validParams.Get(),
			testingFunc.GeneralCtx.Delete("DeviceID").Get(),
			nil,
			errors.BadRequest.New("-"),
		},
	} {
		t.Run(tt.message, func(t *testing.T) {

			res, err := decodeDeleteAccountReq(tt.ctx, httptest.NewRequest("DELETE", fmt.Sprintf("%s?%s", "/account", tt.params.Encode()), nil))
			if testingFunc.CheckError(t, tt.err, err) {
				return
			}

			testingFunc.CheckStruct(t, *tt.want, res, nil)
		})
	}
}
