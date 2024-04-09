package middleware

import (
	"context"
	"encoding/json"
	defErrors "errors"
	"net/http/httptest"
	"testing"

	"server/app/pkg/errors"
)

func TestEncodeErrorResponse(t *testing.T) {

	for _, tt := range []struct {
		name      string
		err       error
		wantError error
	}{
		{
			name:      "1.Empty error",
			err:       nil,
			wantError: errors.InternalServer.New(""),
		},
		{
			name:      "2.Custom error",
			err:       errors.BadRequest.New("test"),
			wantError: errors.BadRequest.New("test"),
		},
		{
			name:      "3.Not wrapped error",
			err:       defErrors.New("test"),
			wantError: errors.InternalServer.New("test"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			DefaultErrorEncoder(context.Background(), w, tt.err, func(error) {})

			wantCustomError := errors.CastError(tt.wantError)
			wantCode := int(wantCustomError.ErrorType)

			if w.Code != wantCode {
				t.Fatalf("Полученный httpCode: %v, ожидаемый: %v. Ошибка:%v", w.Code, wantCode, w.Body.String())
			}

			getCustomErr := errors.CustomError{}
			if err := json.NewDecoder(w.Body).Decode(&getCustomErr); err != nil {
				t.Fatalf("Ошибка декодирования: %v", err)
			}

			getCustomErr.InitialError = defErrors.New(getCustomErr.DevelopText)

			if !errors.As(getCustomErr, tt.wantError) {
				t.Fatalf("Полученная ошибка: %v, ожидаемая: %v", getCustomErr, tt.wantError)
			}

		})
	}
}
