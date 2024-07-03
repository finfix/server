package slices

import (
	"fmt"

	"server/app/pkg/errors"
	"server/app/pkg/maps"
)

// ToMap возращает map, где ключом является поле структуры, а значением сама структура
// Example:
// AccountGroupsMap := slice.ToMap(_accountGroups, func(ag model.AccountGroup) uint32 { return ag.ID })
func ToMap[K comparable, V any](slice []V, field func(V) K) map[K]V {
	mapBySlise := make(map[K]V, len(slice))
	for _, v := range slice {
		mapBySlise[field(v)] = v
	}
	return mapBySlise
}

func GetUniqueByField[K comparable, V any](slice []V, field func(V) K) []V {
	return maps.ToSlice(ToMap(slice, field))
}

func GetFields[K comparable, V any](slice []V, field func(V) K) []K {
	fields := make([]K, 0, len(slice))
	for _, v := range slice {
		fields = append(fields, field(v))
	}
	return fields
}

func In[K comparable](value K, slice ...K) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func GetMapValueStruct[K comparable, V any](slice []V, field func(V) K) map[K]struct{} {
	m := make(map[K]struct{}, len(slice))
	for _, v := range slice {
		m[field(v)] = struct{}{}
	}
	return m
}

func JoinExclusive(leftObjects, rightObjects []uint32) (leftObjectsExclusive, rightObjectsExclusive []uint32) {
	leftObjectsMap := GetMapValueStruct(leftObjects, func(v uint32) uint32 { return v })
	rightObjectsMap := GetMapValueStruct(rightObjects, func(v uint32) uint32 { return v })

	for _, leftObject := range leftObjects {
		if _, ok := rightObjectsMap[leftObject]; !ok {
			leftObjectsExclusive = append(leftObjectsExclusive, leftObject)
		}
	}

	for _, rightObject := range rightObjects {
		if _, ok := leftObjectsMap[rightObject]; !ok {
			rightObjectsExclusive = append(rightObjectsExclusive, rightObject)
		}
	}

	return
}

func First[T any](array []T) *T {
	if len(array) == 0 {
		return nil
	} else {
		value := array[0]
		return &value
	}
}

func FirstWithError[T any](array []T) (value T, err error) {
	valuePtr := First(array)
	if valuePtr == nil {
		return value, errors.NotFound.New("Значение массива не найдено",
			errors.ParamsOption("Type", fmt.Sprintf("%T", value)),
		)
	}
	return *valuePtr, nil
}
