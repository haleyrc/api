package app

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/haleyrc/api"
	"github.com/haleyrc/api/pages"
	"github.com/haleyrc/api/service"
)

func (s *Server) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	f := s.deleteAuthor(r.Context(), chi.URLParam(r, "id"))
	f(w, r)
}

func (s *Server) deleteAuthor(ctx context.Context, id string) http.HandlerFunc {
	if _, err := s.Books.DeleteAuthor(ctx, service.DeleteAuthorRequest{
		ID: api.ID(id),
	}); err != nil {
		return s.RenderError(500, err.Error())
	}
	return s.Redirect(301, "/authors")
}

func (s *Server) GetAuthor(w http.ResponseWriter, r *http.Request) {
	f := s.getAuthor(r.Context(), chi.URLParam(r, "id"))
	f(w, r)
}

func (s *Server) getAuthor(ctx context.Context, id string) http.HandlerFunc {
	var data pages.AuthorGet

	gar, err := s.Books.GetAuthor(ctx, service.GetAuthorRequest{
		ID: api.ID(id),
	})
	if err != nil {
		return s.RenderError(500, err.Error())
	}
	data.Author = gar.Author

	gbr, err := s.Books.GetBooks(ctx, service.GetBooksRequest{
		Author: api.MaybeID{Valid: true, Value: data.Author.ID},
	})
	if err != nil {
		return s.RenderError(500, err.Error())
	}
	data.Books = gbr.Books

	return s.Render(200, s.templates.GetAuthor, data)
}

func (s *Server) GetAuthors(w http.ResponseWriter, r *http.Request) {
	f := s.getAuthors(r.Context())
	f(w, r)
}

func (s *Server) getAuthors(ctx context.Context) http.HandlerFunc {
	var data pages.AuthorsGet
	gar, err := s.Books.GetAuthors(ctx, service.GetAuthorsRequest{})
	if err != nil {
		return s.RenderError(500, err.Error())
	}
	data.Authors = gar.Authors

	return s.Render(200, s.templates.GetAuthors, data)
}

func (s *Server) NewAuthor(w http.ResponseWriter, r *http.Request) {
	f := s.newAuthor(r.Context())
	f(w, r)
}

func (s *Server) newAuthor(ctx context.Context) http.HandlerFunc {
	return s.Render(200, s.templates.NewAuthor, nil)
}

func (s *Server) SaveAuthor(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		s.RenderError(400, err.Error())
	}

	name := r.Form.Get("name")

	f := s.saveAuthor(r.Context(), name)
	f(w, r)
}

func (s *Server) saveAuthor(ctx context.Context, name string) http.HandlerFunc {
	_, err := s.Books.SaveAuthor(ctx, service.SaveAuthorRequest{
		ID:   api.MaybeID{},
		Name: name,
	})
	if err != nil {
		return s.RenderError(400, err.Error())
	}

	return s.Redirect(301, "/authors")
}
