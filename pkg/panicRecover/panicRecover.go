package panicRecover

import (
	"fmt"

	"server/pkg/errors"
)

func PanicRecover(handling func(err error)) {
	if r := recover(); r != nil {
		handling(errors.InternalServer.New(fmt.Sprintf("%v", r),
			errors.SkipThisCallOption(),
		))
	}
}
