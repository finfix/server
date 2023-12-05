package crontab

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"logger/app/logging"
	"pkg/errors"
	"pkg/testingFunc"
)

func TestAuthorization(t *testing.T) {

	logging.Off()

	for _, tt := range []struct {
		message string
		header  map[string]string
		err     error
	}{
		//{"1.Обычный запрос",
		//	map[string]string{"MySecretKey": "0"},
		//	nil,
		//},
		{"2.Неверный ключ",
			map[string]string{"MySecretKey": "00"},
			errors.Forbidden.New("incorrect"),
		},
		//{"3.Пустой ключ",
		//	map[string]string{"MySecretKey": ""},
		//	errors.Forbidden.New("incorrect"),
		//},
	} {
		t.Run(tt.message, func(t *testing.T) {
			r := httptest.NewRequest("", "/", strings.NewReader(""))
			for k, v := range tt.header {
				r.Header.Set(k, v)
			}

			_, err := authorization(context.Background(), r)
			if testingFunc.CheckError(t, tt.err, err) {
				return
			}
		})
	}
}
