package app

import (
	"context"

	"github.com/haleyrc/api"
)

type key int

const (
	_ key = iota
	scopeKey
)

type scope struct {
	User *api.User
}

func contextWithScope(ctx context.Context, s scope) context.Context {
	return context.WithValue(ctx, scopeKey, s)
}

func scopeFromContext(ctx context.Context) scope {
	tmp := ctx.Value(scopeKey)
	if tmp == nil {
		return scope{}
	}
	return tmp.(scope)
}
