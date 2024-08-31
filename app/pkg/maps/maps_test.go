package maps

import (
	"reflect"
	"testing"
)

func TestJoin(t *testing.T) {
	type args[K comparable, V any] struct {
		leftMap  map[K]V
		rightMap map[K]V
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want map[K]V
	}
	tests := []testCase[string, int]{
		{
			name: "1. Объединение двух пустых карт",
			args: args[string, int]{
				leftMap:  map[string]int{},
				rightMap: map[string]int{},
			},
			want: map[string]int{},
		},
		{
			name: "2. Объединение двух карт с одинаковыми ключами",
			args: args[string, int]{
				leftMap:  map[string]int{"a": 1, "b": 2},
				rightMap: map[string]int{"a": 3, "b": 4},
			},
			want: map[string]int{"a": 3, "b": 4},
		},
		{
			name: "3. Объединение двух карт с разными ключами",
			args: args[string, int]{
				leftMap:  map[string]int{"a": 1, "b": 2},
				rightMap: map[string]int{"c": 3, "d": 4},
			},
			want: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
		},
		{
			name: "4. Объединение пустой карты с непустой",
			args: args[string, int]{
				leftMap:  map[string]int{},
				rightMap: map[string]int{"a": 1, "b": 2},
			},
			want: map[string]int{"a": 1, "b": 2},
		},
		{
			name: "5. Объединение неинициализированной карты с непустой",
			args: args[string, int]{
				leftMap:  nil,
				rightMap: map[string]int{"a": 1, "b": 2},
			},
			want: map[string]int{"a": 1, "b": 2},
		},
		{
			name: "6. Объединение двух неинициализированных карт",
			args: args[string, int]{
				leftMap:  nil,
				rightMap: nil,
			},
			want: map[string]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Join(tt.args.leftMap, tt.args.rightMap)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeys(t *testing.T) {
	type args[K comparable, V any] struct {
		mapa map[K]V
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want []K
	}
	tests := []testCase[string, string]{
		{
			name: "1. Получение ключей из пустой карты",
			args: args[string, string]{mapa: map[string]string{}},
			want: []string{},
		},
		{
			name: "2. Получение ключей из непустой карты",
			args: args[string, string]{mapa: map[string]string{"a": "1", "b": "2"}},
			want: []string{"a", "b"},
		},
		{
			name: "3. Получение ключей из неинициализированной мапы",
			args: args[string, string]{mapa: nil},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Keys(tt.args.mapa)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValues(t *testing.T) {
	type args[K comparable, V any] struct {
		mapa map[K]V
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want []V
	}
	tests := []testCase[string, string]{
		{
			name: "1. Получение значений из пустой карты",
			args: args[string, string]{mapa: map[string]string{}},
			want: []string{},
		},
		{
			name: "2. Получение значений из непустой карты",
			args: args[string, string]{mapa: map[string]string{"a": "1", "b": "2"}},
			want: []string{"1", "2"},
		},
		{
			name: "3. Получение значений из неинициализированной мапы",
			args: args[string, string]{mapa: nil},
			want: []string{},
		},
		{
			name: "4. Получение значений из карты с одинаковыми значениями",
			args: args[string, string]{mapa: map[string]string{"a": "1", "b": "1"}},
			want: []string{"1", "1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Values(tt.args.mapa); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}
