package slices

import (
	"reflect"
	"testing"

	"server/app/pkg/errors"
	"server/app/pkg/pointer"
	"server/app/pkg/testUtils"
)

func TestJoinExclusive(t *testing.T) {
	type args[T any] struct {
		left  []T
		right []T
	}
	type result[T any] struct {
		leftExclusive  []T
		rightExclusive []T
	}
	type testCase[T any] struct {
		name   string
		args   args[T]
		result result[T]
	}
	tests := []testCase[int]{
		{
			name: "1. Обычный случай",
			args: args[int]{
				left:  []int{1, 2, 3, 4, 5},
				right: []int{3, 4, 5, 6, 7},
			},
			result: result[int]{
				leftExclusive:  []int{1, 2},
				rightExclusive: []int{6, 7},
			},
		},
		{
			name: "2. Пустые слайсы",
			args: args[int]{
				left:  []int{},
				right: []int{},
			},
			result: result[int]{
				leftExclusive:  []int{},
				rightExclusive: []int{},
			},
		},
		{
			name: "3. Пустой левый слайс",
			args: args[int]{
				left:  []int{},
				right: []int{1, 2, 3},
			},
			result: result[int]{
				leftExclusive:  []int{},
				rightExclusive: []int{1, 2, 3},
			},
		},
		{
			name: "4. Пустой правый слайс",
			args: args[int]{
				left:  []int{1, 2, 3},
				right: []int{},
			},
			result: result[int]{
				leftExclusive:  []int{1, 2, 3},
				rightExclusive: []int{},
			},
		},
		{
			name: "5. Полное совпадение",
			args: args[int]{
				left:  []int{1, 2, 3},
				right: []int{1, 2, 3},
			},
			result: result[int]{
				leftExclusive:  []int{},
				rightExclusive: []int{},
			},
		},
		{
			name: "6. Без совпадений",
			args: args[int]{
				left:  []int{1, 2, 3},
				right: []int{4, 5, 6},
			},
			result: result[int]{
				leftExclusive:  []int{1, 2, 3},
				rightExclusive: []int{4, 5, 6},
			},
		},
		{
			name: "7. Неинициализированные слайсы",
			args: args[int]{
				left:  nil,
				right: nil,
			},
			result: result[int]{
				leftExclusive:  []int{},
				rightExclusive: []int{},
			},
		},
	}
	num1 := pointer.Pointer(1)
	num2 := pointer.Pointer(2)
	num3 := pointer.Pointer(3)
	num4 := pointer.Pointer(4)
	num5 := pointer.Pointer(5)
	num6 := pointer.Pointer(6)
	num7 := pointer.Pointer(7)

	tests2 := []testCase[*int]{
		{
			name: "8. Обычный случай (указатели)",
			args: args[*int]{
				left:  []*int{num1, num2, num3, num4, num5},
				right: []*int{num3, num4, num5, num6, num7},
			},
			result: result[*int]{
				leftExclusive:  []*int{num1, num2},
				rightExclusive: []*int{num6, num7},
			},
		},
		{
			name: "9. Пустые слайсы (указатели)",
			args: args[*int]{
				left:  []*int{},
				right: []*int{},
			},
			result: result[*int]{
				leftExclusive:  []*int{},
				rightExclusive: []*int{},
			},
		},
		{
			name: "10. Пустой левый слайс (указатели)",
			args: args[*int]{
				left:  []*int{},
				right: []*int{num1, num2, num3},
			},
			result: result[*int]{
				leftExclusive:  []*int{},
				rightExclusive: []*int{num1, num2, num3},
			},
		},
		{
			name: "11. Пустой правый слайс (указатели)",
			args: args[*int]{
				left:  []*int{num1, num2, num3},
				right: []*int{},
			},
			result: result[*int]{
				leftExclusive:  []*int{num1, num2, num3},
				rightExclusive: []*int{},
			},
		},
		{
			name: "12. Полное совпадение (указатели)",
			args: args[*int]{
				left:  []*int{num1, num2, num3},
				right: []*int{num1, num2, num3},
			},
			result: result[*int]{
				leftExclusive:  []*int{},
				rightExclusive: []*int{},
			},
		},
		{
			name: "13. Без совпадений (указатели)",
			args: args[*int]{
				left:  []*int{num1, num2, num3},
				right: []*int{num4, num5, num6},
			},
			result: result[*int]{
				leftExclusive:  []*int{num1, num2, num3},
				rightExclusive: []*int{num4, num5, num6},
			},
		},
		{
			name: "14. Совпадение с nil (указатели)",
			args: args[*int]{
				left:  []*int{nil, num2, num3},
				right: []*int{nil, num2, num3},
			},
			result: result[*int]{
				leftExclusive:  []*int{},
				rightExclusive: []*int{},
			},
		},
		{
			name: "15. Без совпадений с nil (указатели)",
			args: args[*int]{
				left:  []*int{nil, num2, num3},
				right: []*int{num4, num5, num6},
			},
			result: result[*int]{
				leftExclusive:  []*int{nil, num2, num3},
				rightExclusive: []*int{num4, num5, num6},
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
	for _, tt := range tests2 {
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

func TestFirst(t *testing.T) {
	type args[T any] struct {
		array []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want *T
	}
	tests := []testCase[int]{
		{
			name: "1. Получение первого элемента из пустого массива",
			args: args[int]{
				array: []int{},
			},
			want: nil,
		},
		{
			name: "2. Получение первого элемента из массива с одним элементом",
			args: args[int]{
				array: []int{1},
			},
			want: pointer.Pointer(1),
		},
		{
			name: "3. Получение первого элемента из массива с несколькими элементами",
			args: args[int]{
				array: []int{1, 2, 3},
			},
			want: pointer.Pointer(1),
		},
		{
			name: "4. Получение первого элемента из неинициализированного массива",
			args: args[int]{
				array: nil,
			},
			want: nil,
		},
	}
	num1 := pointer.Pointer(1)
	num2 := pointer.Pointer(2)
	num3 := pointer.Pointer(3)
	tests2 := []testCase[*int]{
		{
			name: "5. Получение первого элемента из пустого массива (указатели)",
			args: args[*int]{
				array: []*int{},
			},
			want: nil,
		},
		{
			name: "6. Получение первого элемента из массива с одним элементом (указатели)",
			args: args[*int]{
				array: []*int{num1},
			},
			want: &num1,
		},
		{
			name: "7. Получение первого элемента из массива с несколькими элементами (указатели)",
			args: args[*int]{
				array: []*int{num1, num2, num3},
			},
			want: &num1,
		},
		{
			name: "8. Получение первого элемента из неинициализированного массива (указатели)",
			args: args[*int]{
				array: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := First(tt.args.array)
			if got == nil || tt.want == nil {
				if got != tt.want {
					t.Errorf("First() = %v, want %v", got, tt.want)
				}
			} else {
				if *got != *tt.want {
					t.Errorf("First() = %v, want %v", *got, *tt.want)
				}
			}
		})
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			got := First(tt.args.array)
			switch {
			case got == nil || tt.want == nil:
				if got != tt.want {
					t.Errorf("First() = %v, want %v", got, tt.want)
				}
			case *got != *tt.want:
				t.Errorf("First() = %v, want %v", *got, *tt.want)
			case **got != **tt.want:
				t.Errorf("First() = %v, want %v", **got, **tt.want)

			}
		})
	}
}

func TestFirstWithError(t *testing.T) {

	err := errors.New("error")
	type args[T any] struct {
		array      []T
		initialErr error
	}
	type testCase[T any] struct {
		name               string
		args               args[T]
		wantValue          T
		wantErr            error
		compareErrorsTypes bool
	}
	num1 := pointer.Pointer(1)
	num2 := pointer.Pointer(2)
	num3 := pointer.Pointer(3)
	tests := []testCase[int]{
		{
			name: "1. Получение первого элемента из пустого массива",
			args: args[int]{
				array:      []int{},
				initialErr: nil,
			},
			wantValue:          0,
			wantErr:            errors.NotFound.New(""),
			compareErrorsTypes: false,
		},
		{
			name: "2. Получение первого элемента из массива с одним элементом",
			args: args[int]{
				array:      []int{1},
				initialErr: nil,
			},
			wantValue: 1,
		},
		{
			name: "3. Получение первого элемента из массива с несколькими элементами",
			args: args[int]{
				array:      []int{1, 2, 3},
				initialErr: nil,
			},
			wantValue: 1,
		},
		{
			name: "4. Получение первого элемента из неинициализированного массива",
			args: args[int]{
				array:      nil,
				initialErr: nil,
			},
			wantValue:          0,
			wantErr:            errors.NotFound.New(""),
			compareErrorsTypes: false,
		},
		{
			name: "5. Получение первого элемента из пустого массива с ошибкой",
			args: args[int]{
				array:      []int{},
				initialErr: errors.InternalServer.Wrap(err),
			},
			wantValue:          0,
			wantErr:            errors.InternalServer.Wrap(err),
			compareErrorsTypes: true,
		},
	}
	tests2 := []testCase[*int]{
		{
			name: "6. Получение первого элемента из пустого массива (указатели)",
			args: args[*int]{
				array:      []*int{},
				initialErr: nil,
			},
			wantValue:          nil,
			wantErr:            errors.NotFound.New(""),
			compareErrorsTypes: false,
		},
		{
			name: "7. Получение первого элемента из массива с одним элементом (указатели)",
			args: args[*int]{
				array:      []*int{num1},
				initialErr: nil,
			},
			wantValue: num1,
		},
		{
			name: "8. Получение первого элемента из массива с несколькими элементами (указатели)",
			args: args[*int]{

				array:      []*int{num1, num2, num3},
				initialErr: nil,
			},
			wantValue: num1,
		},
		{
			name: "9. Получение первого элемента из неинициализированного массива (указатели)",
			args: args[*int]{
				array:      nil,
				initialErr: nil,
			},
			wantErr:            errors.NotFound.New(""),
			compareErrorsTypes: false,
		},
		{
			name: "10. Получение первого элемента из пустого массива с ошибкой (указатели)",

			args: args[*int]{
				array:      []*int{},
				initialErr: errors.InternalServer.Wrap(err),
			},
			wantValue:          nil,
			wantErr:            errors.InternalServer.Wrap(err),
			compareErrorsTypes: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotErr := FirstWithError(tt.args.array, tt.args.initialErr)
			testUtils.CheckError(t, tt.wantErr, gotErr, tt.compareErrorsTypes)
			if gotValue != tt.wantValue {
				t.Errorf("FirstWithError() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotErr := FirstWithError(tt.args.array, tt.args.initialErr)
			testUtils.CheckError(t, tt.wantErr, gotErr, tt.compareErrorsTypes)
			if gotValue != tt.wantValue {
				t.Errorf("FirstWithError() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestGetFields(t *testing.T) {
	type testStruct struct {
		Field1 string
		Field2 int
	}

	type args[V any, K comparable] struct {
		slice []V
		field func(V) K
	}
	type testCase[V any, K comparable] struct {
		name string
		args args[V, K]
		want []K
	}
	tests := []testCase[testStruct, string]{
		{
			name: "1. Получение полей из пустого слайса",
			args: args[testStruct, string]{
				slice: []testStruct{},
				field: func(v testStruct) string { return v.Field1 },
			},
			want: []string{},
		},
		{
			name: "2. Получение полей из непустого слайса",
			args: args[testStruct, string]{
				slice: []testStruct{
					{Field1: "a", Field2: 1},
					{Field1: "b", Field2: 2},
				},
				field: func(v testStruct) string { return v.Field1 },
			},
			want: []string{"a", "b"},
		},
		{
			name: "3. Получение полей из неинициализированного слайса",
			args: args[testStruct, string]{
				slice: nil,
				field: func(v testStruct) string { return v.Field1 },
			},
			want: []string{},
		},
	}
	type testStructPointer struct {
		Field1 *string
		Field2 int
	}
	a := pointer.Pointer("a")
	b := pointer.Pointer("b")
	tests2 := []testCase[*testStructPointer, *string]{
		{
			name: "4. Получение полей из пустого слайса (указатели)",
			args: args[*testStructPointer, *string]{
				slice: []*testStructPointer{},
				field: func(v *testStructPointer) *string { return v.Field1 },
			},
			want: []*string{},
		},
		{
			name: "5. Получение полей из непустого слайса (указатели)",
			args: args[*testStructPointer, *string]{
				slice: []*testStructPointer{
					{Field1: a, Field2: 1},
					{Field1: b, Field2: 2},
				},
				field: func(v *testStructPointer) *string { return v.Field1 },
			},
			want: []*string{a, b},
		},
		{
			name: "6. Получение полей из слайса с nil (указатели)",
			args: args[*testStructPointer, *string]{
				slice: []*testStructPointer{
					{Field1: nil, Field2: 1},
					{Field1: b, Field2: 2},
				},
				field: func(v *testStructPointer) *string { return v.Field1 },
			},
			want: []*string{nil, b},
		},
		{
			name: "7. Получение полей из неинициализированного слайса (указатели)",
			args: args[*testStructPointer, *string]{
				slice: nil,
				field: func(v *testStructPointer) *string { return v.Field1 },
			},
			want: []*string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFields(tt.args.slice, tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFields() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFields(tt.args.slice, tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMapValueStruct(t *testing.T) {
	type testStruct struct {
		Field1 string
		Field2 int
	}

	type args[V any, K comparable] struct {
		slice []V
		field func(V) K
	}
	type testCase[V any, K comparable] struct {
		name string
		args args[V, K]
		want map[K]struct{}
	}
	tests := []testCase[testStruct, string]{
		{
			name: "1. Получение значений из пустого слайса",
			args: args[testStruct, string]{
				slice: []testStruct{},
				field: func(v testStruct) string { return v.Field1 },
			},
			want: map[string]struct{}{},
		},
		{
			name: "2. Получение значений из непустого слайса",
			args: args[testStruct, string]{
				slice: []testStruct{
					{Field1: "a", Field2: 1},
					{Field1: "b", Field2: 2},
				},
				field: func(v testStruct) string { return v.Field1 },
			},
			want: map[string]struct{}{
				"a": {},
				"b": {},
			},
		},
		{
			name: "3. Получение значений из слайса с дубликатами",
			args: args[testStruct, string]{
				slice: []testStruct{
					{Field1: "a", Field2: 1},
					{Field1: "b", Field2: 2},
					{Field1: "a", Field2: 3},
				},
				field: func(v testStruct) string { return v.Field1 },
			},
			want: map[string]struct{}{
				"a": {},
				"b": {},
			},
		},
		{
			name: "4. Получение значений из неинициализированного слайса",
			args: args[testStruct, string]{
				slice: nil,
				field: func(v testStruct) string { return v.Field1 },
			},
			want: map[string]struct{}{},
		},
	}

	type testStructPointer struct {
		Field1 *string
		Field2 int
	}
	a := pointer.Pointer("a")
	b := pointer.Pointer("b")
	a2 := pointer.Pointer("a")
	tests2 := []testCase[*testStructPointer, *string]{
		{
			name: "5. Получение значений из пустого слайса (указатели)",
			args: args[*testStructPointer, *string]{
				slice: []*testStructPointer{},
				field: func(v *testStructPointer) *string { return v.Field1 },
			},
			want: map[*string]struct{}{},
		},
		{
			name: "6. Получение значений из непустого слайса (указатели)",
			args: args[*testStructPointer, *string]{
				slice: []*testStructPointer{
					{Field1: a, Field2: 1},
					{Field1: b, Field2: 2},
				},
				field: func(v *testStructPointer) *string { return v.Field1 },
			},
			want: map[*string]struct{}{
				a: {},
				b: {},
			},
		},
		{
			name: "7. Получение значений из слайса с дубликатами (указатели)",
			args: args[*testStructPointer, *string]{
				slice: []*testStructPointer{
					{Field1: a, Field2: 1},
					{Field1: b, Field2: 2},
					{Field1: a2, Field2: 3},
				},
				field: func(v *testStructPointer) *string { return v.Field1 },
			},
			want: map[*string]struct{}{
				a:  {},
				b:  {},
				a2: {},
			},
		},
		{
			name: "8. Получение значений из слайса с nil (указатели)",
			args: args[*testStructPointer, *string]{
				slice: []*testStructPointer{
					{Field1: nil, Field2: 1},
					{Field1: b, Field2: 2},
				},
				field: func(v *testStructPointer) *string { return v.Field1 },
			},
			want: map[*string]struct{}{
				nil: {},
				b:   {},
			},
		},
		{
			name: "9. Получение значений из неинициализированного слайса (указатели)",
			args: args[*testStructPointer, *string]{
				slice: nil,
				field: func(v *testStructPointer) *string { return v.Field1 },
			},
			want: map[*string]struct{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetMapValueStruct(tt.args.slice, tt.args.field)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMapValueStruct() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			got := GetMapValueStruct(tt.args.slice, tt.args.field)
			for key := range got {
				if _, ok := tt.want[key]; !ok {
					t.Errorf("GetMapValueStruct() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestIn(t *testing.T) {
	type args[K comparable] struct {
		value K
		slice []K
	}
	type testCase[K comparable] struct {
		name string
		args args[K]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "1. Поиск значения в пустом слайсе",
			args: args[int]{
				value: 1,
				slice: []int{},
			},
			want: false,
		},
		{
			name: "2. Поиск значения в слайсе с одним элементом",
			args: args[int]{
				value: 1,
				slice: []int{1},
			},
			want: true,
		},
		{
			name: "3. Поиск значения в слайсе с несколькими элементами",
			args: args[int]{
				value: 1,
				slice: []int{1, 2, 3},
			},
			want: true,
		},
		{
			name: "4. Поиск значения в слайсе с несколькими элементами (не найдено)",
			args: args[int]{
				value: 4,
				slice: []int{1, 2, 3},
			},
			want: false,
		},
		{
			name: "5. Поиск значения в слайсе с дубликатами",
			args: args[int]{
				value: 1,
				slice: []int{1, 2, 3, 1},
			},
			want: true,
		},
		{
			name: "6. Поиск значения в неинициализированном слайсе",
			args: args[int]{
				value: 1,
				slice: nil,
			},
			want: false,
		},
	}
	num1 := pointer.Pointer(1)
	num2 := pointer.Pointer(2)
	num3 := pointer.Pointer(3)
	tests2 := []testCase[*int]{
		{
			name: "7. Поиск значения в пустом слайсе (указатели)",
			args: args[*int]{
				value: num1,
				slice: []*int{},
			},
			want: false,
		},
		{
			name: "8. Поиск значения в слайсе с одним элементом (указатели)",
			args: args[*int]{
				value: num1,
				slice: []*int{num1},
			},
			want: true,
		},
		{
			name: "9. Поиск значения в слайсе с несколькими элементами (указатели)",
			args: args[*int]{
				value: num1,
				slice: []*int{num1, num2, num3},
			},
			want: true,
		},
		{
			name: "10. Поиск значения в слайсе с несколькими элементами (указатели) (не найдено)",
			args: args[*int]{
				value: pointer.Pointer(4),
				slice: []*int{num1, num2, num3},
			},
			want: false,
		},
		{
			name: "11. Поиск значения в слайсе с дубликатами (указатели)",
			args: args[*int]{
				value: num1,
				slice: []*int{num1, num2, num3, num1},
			},
			want: true,
		},
		{
			name: "12. Поиск значения в слайсе с nil (указатели)",
			args: args[*int]{
				value: nil,
				slice: []*int{nil, num2, num3},
			},
			want: true,
		},
		{
			name: "13. Поиск значения в неинициализированном слайсе (указатели)",
			args: args[*int]{
				value: num1,
				slice: nil,
			},
			want: false,
		},
		{
			name: "14. Поиск значения в неинициализированном слайсе с nil (указатели)",
			args: args[*int]{
				value: nil,
				slice: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := In(tt.args.value, tt.args.slice...)
			if got != tt.want {
				t.Errorf("In() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			got := In(tt.args.value, tt.args.slice...)
			if got != tt.want {
				t.Errorf("In() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMap(t *testing.T) {
	type testStruct struct {
		Field1 string
		Field2 int
	}

	type args[V any, K comparable] struct {
		slice []V
		field func(V) K
	}
	type testCase[V any, K comparable] struct {
		name string
		args args[V, K]
		want map[K]V
	}
	tests := []testCase[testStruct, int]{
		{
			name: "1. Преобразование пустого слайса",
			args: args[testStruct, int]{
				slice: []testStruct{},
				field: func(v testStruct) int { return v.Field2 },
			},
			want: map[int]testStruct{},
		},
		{
			name: "2. Преобразование непустого слайса",
			args: args[testStruct, int]{
				slice: []testStruct{
					{Field1: "a", Field2: 1},
					{Field1: "b", Field2: 2},
				},
				field: func(v testStruct) int { return v.Field2 },
			},
			want: map[int]testStruct{
				1: {Field1: "a", Field2: 1},
				2: {Field1: "b", Field2: 2},
			},
		},
		{
			name: "3. Преобразование неинициализированного слайса",
			args: args[testStruct, int]{
				slice: nil,
				field: func(v testStruct) int { return v.Field2 },
			},
			want: map[int]testStruct{},
		},
		{
			name: "4. Преобразование слайса с дубликатами",
			args: args[testStruct, int]{
				slice: []testStruct{
					{Field1: "a", Field2: 1},
					{Field1: "b", Field2: 2},
					{Field1: "c", Field2: 1},
				},
				field: func(v testStruct) int { return v.Field2 },
			},
			want: map[int]testStruct{
				1: {Field1: "c", Field2: 1},
				2: {Field1: "b", Field2: 2},
			},
		},
	}
	type testStructPointer struct {
		Field1 *string
		Field2 int
	}
	a := pointer.Pointer("a")
	b := pointer.Pointer("b")
	tests2 := []testCase[testStructPointer, *string]{
		{
			name: "5. Преобразование пустого слайса (указатели)",
			args: args[testStructPointer, *string]{
				slice: []testStructPointer{},
				field: func(v testStructPointer) *string { return v.Field1 },
			},
			want: map[*string]testStructPointer{},
		},
		{
			name: "6. Преобразование непустого слайса (указатели)",
			args: args[testStructPointer, *string]{
				slice: []testStructPointer{
					{Field1: a, Field2: 1},
					{Field1: b, Field2: 2},
				},
				field: func(v testStructPointer) *string { return v.Field1 },
			},
			want: map[*string]testStructPointer{
				a: {Field1: a, Field2: 1},
				b: {Field1: b, Field2: 2},
			},
		},
		{
			name: "7. Преобразование слайса с nil (указатели)",
			args: args[testStructPointer, *string]{
				slice: []testStructPointer{
					{Field1: nil, Field2: 1},
					{Field1: b, Field2: 2},
				},
				field: func(v testStructPointer) *string { return v.Field1 },
			},
			want: map[*string]testStructPointer{
				nil: {Field1: nil, Field2: 1},
				b:   {Field1: b, Field2: 2},
			},
		},
		{
			name: "8. Преобразование неинициализированного слайса (указатели)",
			args: args[testStructPointer, *string]{
				slice: nil,
				field: func(v testStructPointer) *string { return v.Field1 },
			},
			want: map[*string]testStructPointer{},
		},
		{
			name: "9. Преобразование слайса с дубликатами (указатели)",
			args: args[testStructPointer, *string]{
				slice: []testStructPointer{
					{Field1: a, Field2: 1},
					{Field1: b, Field2: 2},
					{Field1: a, Field2: 3},
				},
				field: func(v testStructPointer) *string { return v.Field1 },
			},
			want: map[*string]testStructPointer{
				a: {Field1: a, Field2: 3},
				b: {Field1: b, Field2: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToMap(tt.args.slice, tt.args.field)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			got := ToMap(tt.args.slice, tt.args.field)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
