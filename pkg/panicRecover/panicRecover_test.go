package panicRecover

import (
	"testing"

	"pkg/errors"
	"pkg/testUtils"
)

func TestPanicRecover(t *testing.T) {

	t.Run("1. Тестирование перехвата паники", func(t *testing.T) {
		defer PanicRecover(func(err error) {
			testUtils.CheckError(t, errors.InternalServer.New(""), err, false)
		})

		panic("test")
	})

	someFunc := func() {
		panic("test")
	}

	t.Run("2. Тестирование перехвата паники внутри функции", func(t *testing.T) {
		defer PanicRecover(func(err error) {
			testUtils.CheckError(t, errors.InternalServer.New(""), err, false)
		})

		someFunc()
	})
}
