package testUtils

import (
	"context"

	"pkg/log"
)

func IgnoreError[T any](v T, err error) T {
	if err != nil {
		log.Fatal(context.Background(), err)
	}
	return v
}
