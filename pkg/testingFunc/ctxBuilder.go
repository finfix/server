package testingFunc

import (
	"context"
)

func NewCtxBuilder(map[string]any) ctxBuilder {
	return ctxBuilder{
		values: make(map[string]any),
	}
}

type ctxBuilder struct {
	values map[string]any
}

func (b ctxBuilder) Set(key string, value any) ctxBuilder {
	newValues := copyMap(b.values)
	newValues[key] = value
	b.values = newValues
	return b
}

func (b ctxBuilder) Delete(key string) ctxBuilder {
	newValues := copyMap(b.values)
	delete(newValues, key)
	b.values = newValues
	return b
}

func (b ctxBuilder) Get() context.Context {
	ctx := context.Background()
	for key, value := range b.values {
		ctx = context.WithValue(ctx, key, value)
	}
	return ctx
}

func copyMap[T any](m map[string]T) map[string]T {
	newMap := make(map[string]T)
	for key, value := range m {
		newMap[key] = value
	}
	return newMap
}
