package app

import (
	"fmt"
	"net/http"

	"github.com/haleyrc/api"
	"github.com/haleyrc/api/service"
)

func (s *Server) withScope(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// We have to short circuit our middleware so we don't end up in an
		// infinite redirect loop when a cookie is present but with an invalid
		// user ID. We could attempt to just not attach this middleware to the
		// logout route, but it results in a lot more code and it's easy to
		// forget to add it when we do need it, so this is the simpler solution.
		if r.URL.Path == "/auth/logout" {
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		cookie, err := r.Cookie("user")
		if err != nil && err != http.ErrNoCookie {
			fmt.Println(err)
		}

		var sc scope
		if cookie != nil {
			gur, err := s.Users.GetUser(ctx, service.GetUserRequest{
				ID: api.ID(cookie.Value),
			})
			if err != nil {
				s.Logger.Println("invalid user:", cookie.Value)
				http.Redirect(w, r, "/auth/logout", http.StatusMovedPermanently)
				return
			}
			sc.User = &gur.User
		}

		ctx = contextWithScope(ctx, sc)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
