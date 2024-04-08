package endpoint

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"server/app/services/transaction/model"
	"server/pkg/contextKeys"
	"server/pkg/errors"
	"server/pkg/logging"
	"server/pkg/testingFunc"
)

func TestDecodeDeleteReq(t *testing.T) {

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
			testingFunc.GeneralCtx.Delete(contextKeys.UserIDKey).Get(),
			nil,
			errors.BadRequest.New("-"),
		},
		{"5.Отсутствующее поле DeviceID в контексте",
			validParams.Get(),
			testingFunc.GeneralCtx.Delete(contextKeys.DeviceIDKey).Get(),
			nil,
			errors.BadRequest.New("-"),
		},
	} {
		t.Run(tt.message, func(t *testing.T) {

			res, err := decodeDeleteTransactionReq(tt.ctx, httptest.NewRequest(http.MethodDelete, fmt.Sprintf("%s?%s", part, tt.params.Encode()), nil))
			if testingFunc.CheckError(t, tt.err, err) {
				return
			}

			testingFunc.CheckStruct(t, *tt.want, res, nil)
		})
	}
}
