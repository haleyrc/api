package service

import (
	"context"
	"time"

	"github.com/haleyrc/api"
)

type Database interface {
	RunInTransaction(context.Context, func(context.Context, Tx) error) error
}

type Tx interface {
	GetAuthor(ctx context.Context, id api.ID) (api.Author, error)
	GetAuthors(ctx context.Context, filter api.AuthorsFilter) ([]api.Author, error)
	SaveAuthor(ctx context.Context, author api.Author) error
	DeleteAuthor(ctx context.Context, id api.ID) error

	GetBook(ctx context.Context, id api.ID) (api.Book, error)
	GetBooks(ctx context.Context, filter api.BooksFilter) ([]api.Book, uint, error)
	SaveBook(ctx context.Context, book api.Book) error
	DeleteBook(ctx context.Context, id api.ID) error
	AddAuthorToBook(ctx context.Context, book, author api.ID) error
	RateBook(ctx context.Context, user, book api.ID, rating api.Rating) error
	StartBook(ctx context.Context, user, book api.ID, timestamp time.Time) error
	FinishBook(ctx context.Context, user, book api.ID, timestamp time.Time) error

	GetBookGenre(ctx context.Context, id api.ID) (api.BookGenre, error)
	GetBookGenres(ctx context.Context) ([]api.BookGenre, error)
	SaveBookGenre(ctx context.Context, genre api.BookGenre) error
	DeleteBookGenre(ctx context.Context, id api.ID) error

	GetUserByID(ctx context.Context, id api.ID) (api.User, error)
	GetUserByName(ctx context.Context, name string) (api.User, error)
	SaveUser(ctx context.Context, user api.User) error
}
