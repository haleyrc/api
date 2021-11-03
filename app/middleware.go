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

		var sc scope

		cookie, err := r.Cookie("user")
		if err != nil {
			if err != http.ErrNoCookie {
				fmt.Println(err)
			}
		} else {
			gur, err := s.Users.GetUser(ctx, service.GetUserRequest{
				ID: api.ID(cookie.Value),
			})
			if err != nil {
				fmt.Printf("invalid user: %s\n", cookie.Value)
				http.Redirect(w, r, "/auth/logout", http.StatusMovedPermanently)
				return
			}
			sc.User = &gur.User
		}

		ctx = contextWithScope(ctx, sc)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func findUser(users map[string]*api.User, id string) (*api.User, bool) {
	for _, user := range users {
		if string(user.ID) == id {
			return user, true
		}
	}
	return nil, false
}
