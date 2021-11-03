package app

import (
	"net/http"
	"time"

	"github.com/haleyrc/api/service"
)

func (s *Server) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		sc := scopeFromContext(ctx)

		if err := r.ParseForm(); err != nil {
			panic(err)
		}

		name := r.FormValue("name")
		gur, err := s.Users.GetUser(ctx, service.GetUserRequest{Name: name})
		if err != nil {
			// TODO: Render a 400 if wasn't a user not found
			data := newPage(sc)
			data.Error(err.Error())
			s.templates.Login.Render(401, w, data)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "user",
			Value:    string(gur.User.ID),
			MaxAge:   int(time.Hour.Seconds()),
			Secure:   s.secure,
			HttpOnly: true,
			Path:     "/",
		})

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func (s *Server) loginPage() http.HandlerFunc {
	type PageData struct {
		page
	}
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		sc := scopeFromContext(ctx)

		if sc.User != nil {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			return
		}

		page := newPage(sc)
		s.templates.Login.Render(200, w, PageData{page: page})
	}
}

func (s *Server) logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:     "user",
			Value:    "",
			MaxAge:   -1,
			Secure:   s.secure,
			HttpOnly: true,
			Path:     "/",
		})
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}
