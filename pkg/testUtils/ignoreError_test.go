package testUtils

import (
	"testing"
)

func TestIgnoreError(t *testing.T) {
	type args[T any] struct {
		v   T
		err error
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "1. Ошибка пустая",
			args: args[int]{
				v:   1,
				err: nil,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IgnoreError(tt.args.v, tt.args.err)
			if got != tt.want {
				t.Errorf("IgnoreError() = %v, want %v", got, tt.want)
			}
		})
	}
}
