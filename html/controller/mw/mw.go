package mw

import (
	"context"

	"github.com/haleyrc/api"
)

type key int

const (
	userKey key = iota
)

func UserFromContext(ctx context.Context) *api.User {
	tmp := ctx.Value(userKey)
	if tmp == nil {
		return nil
	}
	return tmp.(*api.User)
}
