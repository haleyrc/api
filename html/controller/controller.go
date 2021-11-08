package controller

import (
	"context"
	"io"
	"net/http"

	"github.com/haleyrc/api"
	"github.com/haleyrc/api/html"
	"github.com/haleyrc/api/library"
)

type CookieJar interface {
	Clear(w http.ResponseWriter, key string)
	Set(w http.ResponseWriter, key, value string)
}

type ErrorFunc func(context.Context, error) string

type LibraryRepository interface {
	DeleteBook(ctx context.Context, id api.ID) error
	GetAuthors(ctx context.Context, filter library.AuthorsFilter) ([]library.Author, error)
	GetBook(ctx context.Context, id api.ID) (library.Book, error)
	GetBooks(ctx context.Context, filter library.BooksFilter) ([]library.Book, uint, error)
	GetGenre(ctx context.Context, id api.ID) (library.Genre, error)
	GetGenres(ctx context.Context) ([]library.Genre, error)
	SaveBook(ctx context.Context, book library.Book) error
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
