package slices

import (
	"testing"
)

func TestJoinExclusive(t *testing.T) {
	type args struct {
		left  []uint32
		right []uint32
	}
	type result struct {
		leftExclusive  []uint32
		rightExclusive []uint32
	}
	tests := []struct {
		name   string
		args   args
		result result
	}{
		{
			name: "1. Обычный случай",
			args: args{
				left:  []uint32{1, 2, 3, 4, 5},
				right: []uint32{3, 4, 5, 6, 7},
			},
			result: result{
				leftExclusive:  []uint32{1, 2},
				rightExclusive: []uint32{6, 7},
			},
		},
		{
			name: "2. Пустые слайсы",
			args: args{
				left:  []uint32{},
				right: []uint32{},
			},
			result: result{
				leftExclusive:  []uint32{},
				rightExclusive: []uint32{},
			},
		},
		{
			name: "3. Пустой левый слайс",
			args: args{
				left:  []uint32{},
				right: []uint32{1, 2, 3},
			},
			result: result{
				leftExclusive:  []uint32{},
				rightExclusive: []uint32{1, 2, 3},
			},
		},
		{
			name: "4. Пустой правый слайс",
			args: args{
				left:  []uint32{1, 2, 3},
				right: []uint32{},
			},
			result: result{
				leftExclusive:  []uint32{1, 2, 3},
				rightExclusive: []uint32{},
			},
		},
		{
			name: "5. Полное совпадение",
			args: args{
				left:  []uint32{1, 2, 3},
				right: []uint32{1, 2, 3},
			},
			result: result{
				leftExclusive:  []uint32{},
				rightExclusive: []uint32{},
			},
		},
		{
			name: "6. Без совпадений",
			args: args{
				left:  []uint32{1, 2, 3},
				right: []uint32{4, 5, 6},
			},
			result: result{
				leftExclusive:  []uint32{1, 2, 3},
				rightExclusive: []uint32{4, 5, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			leftExclusive, rightExclusive := JoinExclusive(tt.args.left, tt.args.right)
			if len(leftExclusive) != len(tt.result.leftExclusive) {
				t.Fatalf("\n\nОжидалось: %v\nПолучено: %v\n\n", tt.result.leftExclusive, leftExclusive)
			}
			if len(rightExclusive) != len(tt.result.rightExclusive) {
				t.Fatalf("\n\nОжидалось: %v\nПолучено: %v\n\n", tt.result.rightExclusive, rightExclusive)
			}
			for i := range leftExclusive {
				if leftExclusive[i] != tt.result.leftExclusive[i] {
					t.Fatalf("\n\nОжидалось: %v\nПолучено: %v\n\n", tt.result.leftExclusive, leftExclusive)
				}
			}
			for i := range rightExclusive {
				if rightExclusive[i] != tt.result.rightExclusive[i] {
					t.Fatalf("\n\nОжидалось: %v\nПолучено: %v\n\n", tt.result.rightExclusive, rightExclusive)
				}
			}

		})
	}
}
