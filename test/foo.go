package test

import (
	"context"
)

type contextKey string

var contextKeyA = contextKey("foo")

func Set(ctx context.Context) context.Context {
	return context.WithValue(ctx, contextKeyA, "bar")
}

func Get(ctx context.Context) string {
	val := ctx.Value(id).(string)
	return val
}
