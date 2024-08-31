package pointer

import (
	"testing"
)

func TestPointer(t *testing.T) {

	t.Run("1. Получение указателя на строку", func(t *testing.T) {

		str := "string"
		strPointer := Pointer("string")

		if *strPointer != str {
			t.Errorf("Полученное значение: %v, ожидаемое значение: %v", *strPointer, str)
		}
	})

}
