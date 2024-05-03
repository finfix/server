package endpoint

import (
	"context"
	"net/http/httptest"
	"testing"

	"server/app/pkg/contextKeys"
	"server/app/pkg/errors"
	"server/app/pkg/log"
	"server/app/pkg/testingFunc"
	"server/app/services/account/model"
)

func TestDecodeGetAccountGroupsReq(t *testing.T) {

	log.Off()

	validWant := &model.GetAccountGroupsReq{
		Necessary: testingFunc.ValidNecessary,
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
			testingFunc.GeneralCtx.Delete(contextKeys.UserIDKey).Get(),
			nil,
			errors.BadRequest.New("-"),
		},
		{"3.Отсутствующее поле DeviceID в контексте",
			testingFunc.GeneralCtx.Delete(contextKeys.DeviceIDKey).Get(),
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
