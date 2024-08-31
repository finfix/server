package testUtils

import "fmt"

func IgnoreError[T any](v T, err error) T {
	if err != nil {
		fmt.Println(err.Error())
	}
	return v
}
