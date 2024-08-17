package testUtils

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"server/app/pkg/errors"
)

// CheckError производит проверку ошибок на соответствие типов
func CheckError(t *testing.T, wantErr error, gotErr error, compareErrors bool) bool {

	// Если обе пришедшие ошибки пустые
	if gotErr == nil && wantErr == nil {

		// Проверять нечего, все окей, выходим и уведомляем, что в функцию не были переданы ошибки
		return false
	}

	// Декларируем переменные для хранения кастомных типов ошибок
	var wantCustomErr errors.Error
	var gotCustomErr errors.Error

	// Если одна из ошибок пустая, а другая нет
	if (gotErr == nil) != (wantErr == nil) {

		// Если не пришла ошибка, но должна была
		if wantErr != nil {
			t.Fatalf("\n\nТестируемая функция должна была вернуть ошибку: %v\n\n", wantErr)

		} else { // Если пришла ошибка, но не должна была

			// Если можем закастить ошибку в кастомный тип
			if errors.As(wantErr, &wantCustomErr) {
				t.Fatalf("\n\nТестируемая функция вернула не ожидаемую ошибку. Ошибка: %v\n%v\n\n", gotErr, gotCustomErr.StackTrace[0])
			} else { // Если не можем закастить ошибку в кастомный тип
				t.Fatalf("\n\nТестируемая функция вернула не ожидаемую ошибку. Также ошибка не обернута в кастомный тип. Ошибка: %v\n\n", gotErr)
			}
		}
	}

	// Если пришедшая целевая ошибка не обернута
	if !errors.As(wantErr, &wantCustomErr) {
		t.Fatalf("\n\nОжидаемая ошибка в декларации теста не обернута в кастомный тип. Ошибка: %v\n\n", wantErr)
	}

	// Если пришедшая полученная ошибка не обернута
	if !errors.As(gotErr, &gotCustomErr) {
		t.Fatalf("\n\nПолученная ошибка из тестируемой функции не обернута в кастомный тип. Ошибка: %v\n\n", gotErr)
	}

	// Проверяем соответствие кастомных типов ошибок (соответственно HTTP-кодов)
	if wantCustomErr.ErrorType != gotCustomErr.ErrorType {
		t.Fatalf("\n\nДолжен быть другой кастомный тип ошибки: %v вместо %v. Ошибка: %v\n%v\n\n", wantCustomErr.ErrorType, gotCustomErr.ErrorType, gotErr, gotCustomErr.StackTrace[0])
	}

	// Если передано требование, проверяем ошибки на дефолтные типы (значит где-то может быть завязана логика на этих ошибках)
	if compareErrors {
		if !errors.Is(gotErr, wantErr) {
			t.Fatalf("\n\nДолжен быть другой тип ошибки: %v. Ошибка: %v\n%v\n\n", wantErr, gotErr, gotCustomErr.StackTrace[0])
		}
	}

	// Все проверки пройдены, выходим и уведомляем, что в функцию были переданы ошибки
	return true
}

func CheckStruct(t *testing.T, want, get, typ any) {
	var diff string

	if typ == nil {
		diff = cmp.Diff(get, want)
	} else {
		diff = cmp.Diff(get, want, cmpopts.IgnoreUnexported(typ))
	}

	if diff != "" {
		t.Fatalf("\nДолжны быть другие значения структуры: \n%s", diff)
	}
}
