package slices

import (
	"fmt"

	"pkg/errors"
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

// GetFields возвращает массив значений полей из массива структур
// Example:
// AccountGroupsIDs := slice.GetFields(_accountGroups, func(ag model.AccountGroup) uint32 { return ag.ID })
func GetFields[K comparable, V any](slice []V, field func(V) K) []K {
	fields := make([]K, 0, len(slice))
	for _, v := range slice {
		fields = append(fields, field(v))
	}
	return fields
}

// In проверяет, содержится ли ХОТЬ ОДНО значение в переданном массиве
// Example:
// if slice.In(1, 1, 2, 3) {
func In[K comparable](value K, slice ...K) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// GetMapValueStruct Возвращает проверочную мапу всех возможных значений поля
// Example:
// AccountGroupsIDs := slice.GetMapValueStruct(_accountGroups, func(ag model.AccountGroup) uint32 { return ag.ID })
func GetMapValueStruct[K comparable, V any](slice []V, field func(V) K) map[K]struct{} {
	m := make(map[K]struct{}, len(slice))
	for _, v := range slice {
		m[field(v)] = struct{}{}
	}
	return m
}

// JoinExclusive возвращает массивы, содержащие только уникальные элементы
// (не работает с указателями, так как указатели являются ссылками на разные объекты, даже если значения объектов одинаковы)
// Example:
// leftObjectsExclusive, rightObjectsExclusive := slice.JoinExclusive(leftObjects, rightObjects)
func JoinExclusive[T comparable](leftObjects, rightObjects []T) (leftObjectsExclusive, rightObjectsExclusive []T) {
	leftObjectsMap := GetMapValueStruct(leftObjects, func(v T) T { return v })
	rightObjectsMap := GetMapValueStruct(rightObjects, func(v T) T { return v })

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

	return leftObjectsExclusive, rightObjectsExclusive
}

// First возвращает первый элемент массива
// Если массив пустой, возвращает nil
// example:
// firstElement := slice.First([]int{1, 2, 3})
// Возвращает указатель на копию первого элемента массива!
func First[T any](array []T) *T {
	if len(array) == 0 {
		return nil
	} else {
		return &array[0]
	}
}

// FirstWithError получает массив и ошибку (как правило в таком формате возвращают функции получения массива элементов)
// Если пришедшая ошибка не пустая, просто возвращаем ее
// Если пришедший массив пустой, возвращаем ошибку
// Если массив не пустой, то возвращаем первый элемент этого массива
// example:
// firstElement, err := slice.FirstWithError(s.GetArray())
func FirstWithError[T any](array []T, initialErr error) (value T, err error) {

	// Если пришедшая ошибка не пустая
	if initialErr != nil {

		// Возвращаем ее
		return value, initialErr
	}

	// Получаем первый элемент массива
	valuePtr := First(array)

	// Если элемента нет == массив пустой
	if valuePtr == nil {

		// Возвращаем ошибку
		return value, errors.NotFound.New("Значение массива не найдено",
			errors.ParamsOption("Type", fmt.Sprintf("%T", value)),
		)
	}

	// Возвращаем первый элемент массива
	return *valuePtr, nil
}
