package testUtils

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"pkg/errors"
)

// CheckError производит проверку ошибок на соответствие типов
func CheckError(t *testing.T, wantErr error, gotErr error, compareErrors bool) bool {
	err, wasPassedError := checkError(wantErr, gotErr, compareErrors)
	if err != nil {
		t.Fatal(err)
	}
	return wasPassedError
}

var (
	errWantErrorIsEmpty      = errors.New("Тестируемая функция должна была вернуть ошибку")
	errGotErrorIsNotEmpty    = errors.New("Тестируемая функция вернула не ожидаемую ошибку")
	errGotErrorIsNotWrapped  = errors.New("Тестируемая ошибка в не обернута в кастомный тип")
	errWantErrorIsNotWrapped = errors.New("Ожидаемая ошибка в декларации теста не обернута в кастомный тип")
	errCustomTypesMismatch   = errors.New("Кастомные типы ошибок не совпадают")
	errWrappedTypesMismatch  = errors.New("Обернутые типы ошибок не совпадают")
)

func checkError(wantErr error, gotErr error, compareErrors bool) (err error, wasPassedErrors bool) {
	// Если обе пришедшие ошибки пустые
	if gotErr == nil && wantErr == nil {

		// Проверять нечего, все окей, выходим и уведомляем, что в функцию не были переданы ошибки
		return nil, wasPassedErrors
	}

	wasPassedErrors = true

	// Декларируем переменные для хранения кастомных типов ошибок
	var wantCustomErr errors.Error
	var gotCustomErr errors.Error

	// Если одна из ошибок пустая, а другая нет
	if (gotErr == nil) != (wantErr == nil) {

		// Если не пришла ошибка, но должна была
		if wantErr != nil {
			return fmt.Errorf("\n\n%w:\n%v\n\n", errWantErrorIsEmpty, printErrorWithIndent(wantErr)), wasPassedErrors

		} else { // Если пришла ошибка, но не должна была

			// Если можем закастить ошибку в кастомный тип
			if errors.As(gotErr, &gotCustomErr) {
				return fmt.Errorf("\n\n%w:\n%v\n\n", errGotErrorIsNotEmpty, printErrorWithIndent(gotErr)), wasPassedErrors
			} else { // Если не можем закастить ошибку в кастомный тип
				return fmt.Errorf("\n\n%w: %w: %v\n\n", errGotErrorIsNotEmpty, errGotErrorIsNotWrapped, printErrorWithIndent(gotErr)), wasPassedErrors
			}
		}
	}

	// Если пришедшая целевая ошибка не обернута
	if !errors.As(wantErr, &wantCustomErr) {
		return fmt.Errorf("\n\n%w:\n%v\n\n", errWantErrorIsNotWrapped, printErrorWithIndent(wantErr)), wasPassedErrors
	}

	// Если пришедшая полученная ошибка не обернута
	if !errors.As(gotErr, &gotCustomErr) {
		return fmt.Errorf("\n\n%w:\n%v\n\n", errGotErrorIsNotWrapped, printErrorWithIndent(gotErr)), wasPassedErrors
	}

	// Проверяем соответствие кастомных типов ошибок (соответственно HTTP-кодов)
	if wantCustomErr.ErrorType != gotCustomErr.ErrorType {
		return fmt.Errorf("\n\n%w: %v вместо %v:\n%v\n\n", errCustomTypesMismatch, wantCustomErr.ErrorType, gotCustomErr.ErrorType, printErrorWithIndent(gotErr)), wasPassedErrors
	}

	// Если передано требование, проверяем ошибки на дефолтные типы (значит где-то может быть завязана логика на этих ошибках)
	if compareErrors {
		if !errors.Is(gotErr, wantErr) {
			return fmt.Errorf("\n\n%w: %v:\n%v\n\n", errWrappedTypesMismatch, wantErr, printErrorWithIndent(gotErr)), wasPassedErrors
		}
	}

	// Все проверки пройдены, выходим и уведомляем, что в функцию были переданы ошибки
	return nil, wasPassedErrors
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

func printErrorWithIndent(err error) string {
	if err == nil {
		return ""
	}
	var customError errors.Error
	if errors.As(err, &customError) {
		customError.DeveloperText = customError.Error()
		return string(IgnoreError(json.MarshalIndent(customError, "", "\t")))
	} else {
		return err.Error()

	}
}
