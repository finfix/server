package endpoint

import (
	"context"
	"net/http/httptest"
	"testing"

	"server/app/pkg/contextKeys"
	"server/app/pkg/errors"
	"server/app/pkg/logging"
	testingFunc2 "server/app/pkg/testingFunc"
	"server/app/services/account/model"
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
			testingFunc2.GeneralCtx.Get(),
			validWant,
			nil,
		},
		{"2.Отсутствующее поле UserID в контексте",
			testingFunc2.GeneralCtx.Delete(contextKeys.UserIDKey).Get(),
			nil,
			errors.BadRequest.New("-"),
		},
		{"3.Отсутствующее поле DeviceID в контексте",
			testingFunc2.GeneralCtx.Delete(contextKeys.DeviceIDKey).Get(),
			nil,
			errors.BadRequest.New("-"),
		},
	} {
		t.Run(tt.message, func(t *testing.T) {

			res, err := decodeGetAccountGroupsReq(tt.ctx, httptest.NewRequest("", "/", nil))
			if testingFunc2.CheckError(t, tt.err, err) {
				return
			}

			testingFunc2.CheckStruct(t, *tt.want, res, nil)
		})
	}
}
