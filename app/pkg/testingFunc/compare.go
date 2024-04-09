package testingFunc

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"server/app/pkg/errors"
)

func CheckError(t *testing.T, wantErr error, gotErr error) bool {
	if gotErr != nil || wantErr != nil {

		wantCustomErr := errors.CastError(wantErr)
		gotCustomErr := errors.CastError(gotErr)

		switch {
		case gotErr == nil && wantErr != nil:
			t.Fatalf("\n\nДолжна быть ошибка: %v\n\n", wantErr)
		case wantCustomErr.ErrorType == 0:
			customErr := castError(t, gotErr)
			t.Fatalf("\n\nОшибки быть не должно. Ошибка: %v\n%v\n\n", gotErr, customErr.Path)
		case wantCustomErr.ErrorType != gotCustomErr.ErrorType:
			customErr := castError(t, gotErr)
			t.Fatalf("\n\nДолжен быть другой тип ошибки: %v вместо %v. Ошибка: %v\n%v\n\n", wantCustomErr.ErrorType, gotCustomErr.ErrorType, gotErr, customErr.Path)
		case !errors.As(gotErr, wantErr):
			customErr := castError(t, gotErr)
			t.Fatalf("\n\nДолжен быть другой текст ошибки: %v. Ошибка: %v\n%v\n\n", wantErr, gotErr, customErr.Path)
		}

		return true
	}
	return false
}

func castError(t *testing.T, err error) errors.CustomError {
	customErr, ok := err.(errors.CustomError)
	if !ok {
		t.Fatalf("\n\nОшибка не обернута в кастомный тип. Ошибка: %v\n\n", err)
	}
	return customErr
}

func CheckStruct(t *testing.T, want, get, typ any) {
	var diff string

	if typ == nil {
		diff = cmp.Diff(get, want)
	} else {
		diff = cmp.Diff(get, want, cmpopts.IgnoreUnexported(typ))
	}

	if diff != "" {
		t.Fatalf("\nДолжны быть другие значения структуры: \n" + diff)
	}
}
