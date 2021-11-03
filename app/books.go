package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/haleyrc/api"
	"github.com/haleyrc/api/service"
)

func (s *Server) bookPage() http.HandlerFunc {
	type PageData struct {
		page

		Book *api.Book
	}
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		scope := scopeFromContext(ctx)
		data := PageData{page: newPage(scope)}

		id := chi.URLParam(r, "id")
		gbr, err := s.Books.GetBook(ctx, service.GetBookRequest{
			ID: api.ID(id),
		})
		if err != nil {
			data.Error(err.Error())
			s.templates.GetBook.Render(400, w, data)
			return
		}
		data.Book = &gbr.Book

		s.templates.GetBook.Render(200, w, data)
	}
}

func (s *Server) booksPage() http.HandlerFunc {
	type PageData struct {
		page

		Books      []api.Book
		TotalBooks uint64
	}
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		scope := scopeFromContext(ctx)
		data := PageData{page: newPage(scope)}

		gbr, err := s.Books.GetBooks(ctx, service.GetBooksRequest{})
		if err != nil {
			data.Error(err.Error())
			s.templates.GetBook.Render(400, w, data)
			return
		}
		data.Books = gbr.Books
		data.TotalBooks = gbr.Count

		s.templates.GetBooks.Render(200, w, data)
	}
}
