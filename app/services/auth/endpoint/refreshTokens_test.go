package endpoint

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"server/app/pkg/errors"
	"server/app/pkg/log"
	"server/app/pkg/testingFunc"
	"server/app/services/auth/model"
)

func TestDecodeRefreshTokens(t *testing.T) {

	log.Off()

	validJSON := testingFunc.NewJSONUpdater(t, `{
		"token": "token"
	}`)

	for _, tt := range []struct {
		message, body string
		ctx           context.Context
		want          *model.RefreshTokensReq
		err           error
	}{
		{"1.Обычный запрос",
			validJSON.Get(),
			testingFunc.GeneralCtx.Get(),
			&model.RefreshTokensReq{
				Token:     "token",
				Necessary: testingFunc.ValidNecessary,
			},
			nil,
		},
		{"2.Невалидный json",
			testingFunc.InvalidJSON,
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("invalid"),
		},
		{"3.Пустой json",
			"{}",
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("token"),
		},
		{"4.Пустой запрос",
			"",
			testingFunc.GeneralCtx.Get(),
			nil,
			errors.BadRequest.New("EOF"),
		},
	} {
		t.Run(tt.message, func(t *testing.T) {

			res, err := decodeRefreshTokensReq(tt.ctx, httptest.NewRequest("", "/", strings.NewReader(tt.body)))
			if testingFunc.CheckError(t, tt.err, err) {
				return
			}

			testingFunc.CheckStruct(t, *tt.want, res, nil)
		})
	}
}
