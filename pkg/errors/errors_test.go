package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestIs(t *testing.T) {

	var (
		err        = errors.New("err")
		wrappedErr = fmt.Errorf("wrappedErr: %w", err)
		customErr  = InternalServer.Wrap(err)
	)

	var (
		anotherErr       = errors.New("err")
		anotherCustomErr = InternalServer.Wrap(anotherErr)
	)

	var tests = []struct {
		name   string
		err    error
		target error
		want   bool
	}{
		{
			"1. Сравнение двух дефолтных ошибок", // default(1) - default(1)
			err,
			err,
			true,
		},
		{
			"2. Сравнение обернутой ошибки с помощью fmt.Errorf с первоначальной ошибкой", // default(default(1)) - default(1)
			wrappedErr,
			err,
			true,
		},
		{
			"3. Сравнение двух разных ошибок с одинаковым текстом", // default(1) - default(2)
			err,
			anotherErr,
			false,
		},
		{
			"4. Сравнение кастомной обернутой ошибки и дефолтной", // custom(default(1)) - default(1)
			customErr,
			err,
			true,
		},
		{
			"5. Сравнение кастомной обернутой ошибки и другой ошибки с одинаковым текстом", // custom(default(1)) - default(2)
			customErr,
			anotherErr,
			false,
		},
		{
			"6. Сравнение кастомной обернутой ошибки и той же ошибки, обернутой с помощью fmt.Errorf", // custom(default(1)) - default(default(1))
			customErr,
			anotherErr,
			false,
		},
		{
			"7. Сравнение дефолтной ошибки с такой же, завернутой в кастомную", // default(1) - custom(default(1))
			err,
			customErr,
			true,
		},
		{
			"8. Сравнение двух разных кастомных ошибок", // custom(default(1)) - custom(default(2))
			customErr,
			anotherCustomErr,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Is(tt.err, tt.target); got != tt.want {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}
