package controller

import (
	"context"
	"io"
	"net/http"

	"github.com/haleyrc/api"
	"github.com/haleyrc/api/html"
)

type CookieJar interface {
	Clear(w http.ResponseWriter, key string)
	Set(w http.ResponseWriter, key, value string)
}

type ErrorFunc func(context.Context, error) string

type LibraryRepository interface {
	DeleteBook(ctx context.Context, id api.ID) error
	GetAuthors(ctx context.Context, filter api.AuthorsFilter) ([]api.Author, error)
	GetBook(ctx context.Context, id api.ID) (api.Book, error)
	GetBooks(ctx context.Context, filter api.BooksFilter) ([]api.Book, uint, error)
	GetGenre(ctx context.Context, id api.ID) (api.BookGenre, error)
	GetGenres(ctx context.Context) ([]api.BookGenre, error)
	SaveBook(ctx context.Context, book api.Book) error
}

type Renderer interface {
	Redirect(ctx context.Context, w http.ResponseWriter, r *http.Request, code int, path string)
	Render(ctx context.Context, w http.ResponseWriter, code int, t html.Template, data interface{})
}

type Template interface {
	Execute(ctx context.Context, w io.Writer, data interface{}) error
}

type UserRepository interface {
	GetUser(ctx context.Context, query api.UserQuery) (api.User, error)
}
