package testingFunc

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"server/app/pkg/errors"
)

func FirstReturnedElement[T any](t T, another ...any) T {
	return t
}

func CheckError(t *testing.T, wantErr error, gotErr error) bool {

	var (
		customWantErr  errors.Error
		customGotError errors.Error
	)

	switch {

	// Если обе ошибки пустые
	case gotErr == nil && wantErr == nil:
		return false

	// Если не пришла ошибка, но должна была
	case gotErr == nil && wantErr != nil:
		t.Fatalf("Должна быть ошибка: \n%v\n\n", wantErr)

	// Если пришла ошибка, но не должна была
	case wantErr == nil && gotErr != nil:
		if errors.As(gotErr, &customGotError) {
			t.Fatalf("Ошибки быть не должно. Ошибка: \n%v\n%v\n\n", FirstReturnedElement(errors.JSON(customGotError)), customGotError.StackTrace)
		} else {
			t.Fatalf("Ошибки быть не должно. Ошибка: \n%v\n\n", gotErr)
		}

	case wantErr != nil && gotErr != nil:

		// Проверяем совпадение типов ошибок
		if errors.As(wantErr, &customWantErr) && errors.As(gotErr, &customGotError) {
			if customGotError.ErrorType != customWantErr.ErrorType {
				t.Fatalf("Ошибки разных типов. Тип полученной ошибки: %v, тип ожидаемой ошибки: %v\n\n", customGotError.ErrorType, customWantErr.ErrorType)
			}
		}

		// Проверяем совпадение ошибок
		if !errors.Is(gotErr, wantErr) {
			t.Fatalf("Разные ошибки. \nПолученная ошибка: %v\nожидаемая ошибка: %v\n\n", gotErr, wantErr)
		}
	}

	return true
}

func castError(t *testing.T, err error) errors.Error {
	var customErr errors.Error
	ok := errors.As(err, &customErr)
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
