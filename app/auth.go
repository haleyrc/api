package app

import (
	"context"
	"net/http"

	"github.com/haleyrc/api/errors"
	"github.com/haleyrc/api/pages"
	"github.com/haleyrc/api/service"
)

func (s *Server) DoLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		s.RenderError(400, err.Error())(w, r)
		return
	}
	f := s.doLogin(r.Context(), r.FormValue("name"))
	f(w, r)
}

func (s *Server) doLogin(ctx context.Context, name string) http.HandlerFunc {
	gur, err := s.Users.GetUser(ctx, service.GetUserRequest{
		Name: name,
	})
	if err != nil {
		switch errors.Kind(err) {
		case errors.KindResourceNotFound:
			var data pages.Login
			data.Error("Invalid username or password.")
			return s.Render(401, s.templates.Login, data)
		default:
			return s.RenderError(500, err.Error())
		}
	}

	// TODO: This works, but it feels kind of gross
	return s.All(
		s.SetCookie("user", string(gur.User.ID)),
		s.Redirect(301, "/"),
	)
}

func (s *Server) GetLogin(w http.ResponseWriter, r *http.Request) {
	scope := scopeFromContext(r.Context())
	if scope.User != nil {
		s.Redirect(301, "/")(w, r)
	}
	f := s.getLogin(r.Context())
	f(w, r)
}

func (s *Server) getLogin(ctx context.Context) http.HandlerFunc {
	return s.Render(200, s.templates.Login, nil)
}

func (s *Server) DoLogout(w http.ResponseWriter, r *http.Request) {
	f := s.doLogout(r.Context())
	f(w, r)
}

func (s *Server) doLogout(ctx context.Context) http.HandlerFunc {
	return s.All(
		s.ClearCookie("user"),
		s.Redirect(301, "/"),
	)
}
