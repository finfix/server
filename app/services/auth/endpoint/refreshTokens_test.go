package endpoint

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"server/app/services/auth/model"
	"server/pkg/errors"
	"server/pkg/logging"
	"server/pkg/testingFunc"
)

func TestDecodeRefreshTokens(t *testing.T) {

	logging.Off()

	validJSON := testingFunc.NewJSONUpdater(t, `{
		"token": "token"
	}`)

	for _, tt := range []struct {
		message, body string
		want          *model.RefreshTokensReq
		err           error
	}{
		{"1.Обычный запрос",
			validJSON.Get(),
			&model.RefreshTokensReq{
				Token: "token",
			},
			nil,
		},
		{"2.Невалидный json",
			testingFunc.InvalidJSON,
			nil,
			errors.BadRequest.New("invalid"),
		},
		{"3.Пустой json",
			"{}",
			nil,
			errors.BadRequest.New("token"),
		},
		{"4.Пустой запрос",
			"",
			nil,
			errors.BadRequest.New("EOF"),
		},
	} {
		t.Run(tt.message, func(t *testing.T) {

			res, err := decodeRefreshTokensReq(context.Background(), httptest.NewRequest("", "/", strings.NewReader(tt.body)))
			if testingFunc.CheckError(t, tt.err, err) {
				return
			}

			testingFunc.CheckStruct(t, *tt.want, res, nil)
		})
	}
}
