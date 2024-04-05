package panicRecover

import (
	"fmt"
	"runtime"

	"server/pkg/errors"
	"server/pkg/strings"
)

func getErrorFromPanic(r interface{}) error {
	var pcs [32]uintptr
	n := runtime.Callers(errors.ThirdPathDepth, pcs[:])
	textErr := fmt.Sprintf("%v", r)

	for i := 3; i < n; i++ {
		_, file, line, _ := runtime.Caller(i)
		if strings.StringsContains(file, []string{"coin", "Coin"}) {
			err := errors.InternalServer.New(textErr, errors.Options{PathDepth: errors.SecondPathDepth + i})
			return err
		}
	}
	return errors.InternalServer.New(textErr, errors.Options{ErrMessage: "Не удалось получить путь из стека"})
}

func PanicRecover(handling func(err error)) {
	if r := recover(); r != nil {
		err := getErrorFromPanic(r)
		handling(err)
	}
}
