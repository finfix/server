package account

import (
	"context"
	"net/http/httptest"
	"testing"

	"logger/app/logging"
	"pkg/errors"
	"pkg/testingFunc"

	"jsonapi/app/internal/services/account/model"
)

func TestDecodeGetAccountGroupsReq(t *testing.T) {

	logging.Off()

	validWant := &model.GetAccountGroupsReq{
		UserID:   1,
		DeviceID: "DeviceID",
	}

	for _, tt := range []struct {
		message string
		ctx     context.Context
		want    *model.GetAccountGroupsReq
		err     error
	}{
		{"1.Обычный запрос",
			testingFunc.GeneralCtx.Get(),
			validWant,
			nil,
		},
		{"2.Отсутствующее поле UserID в контексте",
			testingFunc.GeneralCtx.Delete("UserID").Get(),
			nil,
			errors.BadRequest.New("-"),
		},
		{"3.Отсутствующее поле DeviceID в контексте",
			testingFunc.GeneralCtx.Delete("DeviceID").Get(),
			nil,
			errors.BadRequest.New("-"),
		},
	} {
		t.Run(tt.message, func(t *testing.T) {

			res, err := decodeGetAccountGroupsReq(tt.ctx, httptest.NewRequest("", "/", nil))
			if testingFunc.CheckError(t, tt.err, err) {
				return
			}

			testingFunc.CheckStruct(t, *tt.want, res, nil)
		})
	}
}
