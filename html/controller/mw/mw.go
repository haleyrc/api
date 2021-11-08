package mw

import (
	"context"
	"errors"
	"net/http"

	"github.com/haleyrc/api"
)

type Middleware func(http.Handler) http.Handler
type CookieJar interface {
	Get(r *http.Request, key string) (string, error)
}
type UserRepo interface {
	GetUser(context.Context, api.UserQuery) (api.User, error)
}
type ErrorFunc func(context.Context, error) string

type key int

const (
	userKey key = iota
)

func ContextWithUser(ctx context.Context, user *api.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func UserFromContext(ctx context.Context) *api.User {
	tmp := ctx.Value(userKey)
	if tmp == nil {
		return nil
	}
	return tmp.(*api.User)
}

func Authenticate(cookies CookieJar, users UserRepo, errorFunc ErrorFunc) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			id, err := cookies.Get(r, "user")
			if err != nil {
				if !errors.Is(err, http.ErrNoCookie) {
					errorFunc(ctx, err)
				}
				next.ServeHTTP(w, r)
				return
			}

			user, err := users.GetUser(ctx, api.UserQuery{
				ID: api.ID(id),
			})
			if err != nil {
				errorFunc(ctx, err)
				next.ServeHTTP(w, r)
				return
			}

			ctx = ContextWithUser(ctx, &user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RequireAuthenticated(redirectTo string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			user := UserFromContext(ctx)
			if user == nil {
				http.Redirect(w, r, redirectTo, http.StatusMovedPermanently)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
