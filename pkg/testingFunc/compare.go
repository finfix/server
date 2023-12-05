package testingFunc

import (
	"testing"

	"pkg/errors"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func CheckError(t *testing.T, wantErr error, getErr error) bool {
	if getErr != nil || wantErr != nil {

		wantType := errors.GetType(wantErr)

		switch {
		case getErr == nil && wantErr != nil:
			t.Fatalf("\n\nДолжна быть ошибка: %v\n\n", wantErr)
		case wantType == errors.NoType:
			customErr := castError(t, getErr)
			t.Fatalf("\n\nОшибки быть не должно. Ошибка: %v\n%v\n\n", getErr, customErr.Path)
		case wantType != errors.GetType(getErr):
			customErr := castError(t, getErr)
			t.Fatalf("\n\nДолжен быть другой тип ошибки: %v вместо %v. Ошибка: %v\n%v\n\n", errors.GetType(wantErr), errors.GetType(getErr), getErr, customErr.Path)
		case !errors.As(getErr, wantErr):
			customErr := castError(t, getErr)
			t.Fatalf("\n\nДолжен быть другой текст ошибки: %v. Ошибка: %v\n%v\n\n", wantErr, getErr, customErr.Path)
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
