package chain

import (
	"context"
	"encoding/json"
	defErrors "errors"
	"net/http/httptest"
	"testing"

	"pkg/errors"
	"pkg/log"
	"pkg/testUtils"
)

func TestEncodeErrorResponse(t *testing.T) {

	log.Off()

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
			DefaultErrorEncoder(context.Background(), w, tt.err)

			wantCustomError := errors.CastError(tt.wantError)
			wantCode := int(wantCustomError.ErrorType)

			if w.Code != wantCode {
				t.Fatalf("Полученный httpCode: %v, ожидаемый: %v. Ошибка:%v", w.Code, wantCode, w.Body.String())
			}

			var getCustomErr errors.Error
			if err := json.NewDecoder(w.Body).Decode(&getCustomErr); err != nil {
				t.Fatalf("Ошибка декодирования: %v", err)
			}
			getCustomErr.ErrorType = errors.ErrorType(w.Code)

			testUtils.CheckError(t, tt.wantError, getCustomErr, false)

		})
	}
}
