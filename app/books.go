package app

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/haleyrc/api"
	"github.com/haleyrc/api/pages"
	"github.com/haleyrc/api/service"
)

func (s *Server) GetBook(w http.ResponseWriter, r *http.Request) {
	f := s.getBook(r.Context(), chi.URLParam(r, "id"))
	f(w, r)
}

func (s *Server) getBook(ctx context.Context, id string) http.HandlerFunc {
	var data pages.BookGet

	gbr, err := s.Books.GetBook(ctx, service.GetBookRequest{
		ID: api.ID(id),
	})
	if err != nil {
		return s.RenderError(404, err.Error())
	}
	data.Book = &gbr.Book

	return s.Render(200, s.templates.GetBook, data)
}

func (s *Server) GetBooks(w http.ResponseWriter, r *http.Request) {
	f := s.getBooks(r.Context())
	f(w, r)
}

func (s *Server) getBooks(ctx context.Context) http.HandlerFunc {
	var data pages.BooksGet

	gbr, err := s.Books.GetBooks(ctx, service.GetBooksRequest{})
	if err != nil {
		return s.RenderError(404, err.Error())
	}
	data.Books = gbr.Books
	data.TotalBooks = gbr.Count

	return s.Render(200, s.templates.GetBooks, data)
}
