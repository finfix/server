package endpoint

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"server/app/pkg/errors"
	"server/app/pkg/log"
	"server/app/pkg/testingFunc"
)

func TestAuthorization(t *testing.T) {

	log.Off()

	for _, tt := range []struct {
		message string
		header  map[string]string
		err     error
	}{
		/*{"1.Обычный запрос",
			 map[string]string{"AdminSecretKey": "0"},
			 nil,
		 },*/
		{"2.Неверный ключ",
			map[string]string{"AdminSecretKey": "00"},
			errors.Forbidden.New("incorrect"),
		},
		/*{"3.Пустой ключ",
			map[string]string{"AdminSecretKey": ""},
			errors.Forbidden.New("incorrect"),
		},*/
	} {
		t.Run(tt.message, func(t *testing.T) {
			r := httptest.NewRequest("", "/", strings.NewReader(""))
			for k, v := range tt.header {
				r.Header.Set(k, v)
			}

			_, err := authorizationWithAdminKey(context.Background(), r)
			if testingFunc.CheckError(t, tt.err, err) {
				return
			}
		})
	}
}
