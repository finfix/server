package testUtils

import (
	"context"

	"server/pkg/contextKeys"
)

func NewCtxBuilder() CtxBuilder {
	return CtxBuilder{
		values: make(map[contextKeys.ContextKey]any),
	}
}

type CtxBuilder struct {
	values map[contextKeys.ContextKey]any
}

func (b CtxBuilder) Set(key contextKeys.ContextKey, value any) CtxBuilder {
	newValues := copyMap(b.values)
	newValues[key] = value
	b.values = newValues
	return b
}

func (b CtxBuilder) Delete(key contextKeys.ContextKey) CtxBuilder {
	newValues := copyMap(b.values)
	delete(newValues, key)
	b.values = newValues
	return b
}

func (b CtxBuilder) Get() context.Context {
	ctx := context.Background()
	for key, value := range b.values {
		ctx = context.WithValue(ctx, key, value)
	}
	return ctx
}

func copyMap[K comparable, V any](m map[K]V) map[K]V {
	newMap := make(map[K]V)
	for key, value := range m {
		newMap[key] = value
	}
	return newMap
}
