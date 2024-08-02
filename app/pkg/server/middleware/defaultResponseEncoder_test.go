package middleware

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"server/app/pkg/testingFunc"
)

func TestEncodeResponse(t *testing.T) {

	for _, tt := range []struct {
		name string
		res  any
		err  error
	}{
		{
			name: "1.With response byte",
			res:  []byte("test"),
			err:  nil,
		},
		{
			name: "2.With response struct",
			res: struct {
				Test string `json:"test"`
			}{Test: "test"},
			err: nil,
		},
		{
			name: "3.With response string",
			res:  "test",
			err:  nil,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {

			tt := tt

			w := httptest.NewRecorder()

			err := DefaultResponseEncoder(context.Background(), w, tt.res)
			if testUtils.CheckError(t, tt.err, err) {
				return
			}

			get := tt.res

			if err := json.NewDecoder(w.Body).Decode(&get); err != nil {
				t.Fatalf("Ошибка декодирования: %v", err)
			}

			// Костыль
			byt, err := json.Marshal(tt.res)
			if err != nil {
				t.Fatalf("Ошибка маршалинга: %v", err)
			}
			if err := json.Unmarshal(byt, &tt.res); err != nil {
				t.Fatalf("Ошибка декодирования: %v", err)
			}

			testUtils.CheckStruct(t, tt.res, get, nil)
		})
	}
}
