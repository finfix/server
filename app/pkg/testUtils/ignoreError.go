package testUtils

func IgnoreError[T any](v T, _ error) T {
	return v
}
