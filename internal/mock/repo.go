package mock

import (
	"context"
	"fmt"
	"time"

	"github.com/haleyrc/api"
	"github.com/haleyrc/api/service"
)

type Repository struct {
	Books      []api.Book
	BookGenres []api.BookGenre
	Users      []api.User
}

func (r *Repository) RunInTransaction(ctx context.Context, f func(ctx context.Context, tx service.Tx) error) error {
	return f(ctx, r)
}

func (r *Repository) GetAuthor(ctx context.Context, id api.ID) (api.Author, error) {
	return api.Author{}, fmt.Errorf("not implemented")
}

func (r *Repository) GetAuthors(ctx context.Context, offset, limit uint) ([]api.Author, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *Repository) SaveAuthor(ctx context.Context, author api.Author) error {
	return fmt.Errorf("not implemented")
}

func (r *Repository) DeleteAuthor(ctx context.Context, id api.ID) error {
	return fmt.Errorf("not implemented")
}

func (r *Repository) GetBook(ctx context.Context, id api.ID) (api.Book, error) {
	for _, book := range r.Books {
		if book.ID == id {
			genre, err := r.GetBookGenre(ctx, book.Genre.ID)
			if err != nil {
				return api.Book{}, fmt.Errorf("get book failed: %w", err)
			}
			book.Genre = genre
			return book, nil
		}
	}
	return api.Book{}, fmt.Errorf("get book failed: resource not found")
}

func (r *Repository) GetBooks(ctx context.Context, offset, limit uint) ([]api.Book, error) {
	return r.Books, nil
}

func (r *Repository) SaveBook(ctx context.Context, book api.Book) error {
	return fmt.Errorf("not implemented")
}

func (r *Repository) DeleteBook(ctx context.Context, id api.ID) error {
	return fmt.Errorf("not implemented")
}

func (r *Repository) AddAuthorToBook(ctx context.Context, book, author api.ID) error {
	return fmt.Errorf("not implemented")
}

func (r *Repository) RateBook(ctx context.Context, user, book api.ID, rating api.Rating) error {
	return fmt.Errorf("not implemented")
}

func (r *Repository) StartBook(ctx context.Context, user, book api.ID, timestamp time.Time) error {
	return fmt.Errorf("not implemented")
}

func (r *Repository) FinishBook(ctx context.Context, user, book api.ID, timestamp time.Time) error {
	return fmt.Errorf("not implemented")
}

func (r *Repository) GetBookGenre(ctx context.Context, id api.ID) (api.BookGenre, error) {
	for _, genre := range r.BookGenres {
		if genre.ID == id {
			return genre, nil
		}
	}
	return api.BookGenre{}, fmt.Errorf("get book genre failed: resource not found")
}

func (r *Repository) GetBookGenres(ctx context.Context) ([]api.BookGenre, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *Repository) SaveBookGenre(ctx context.Context, genre api.BookGenre) error {
	return fmt.Errorf("not implemented")
}

func (r *Repository) DeleteBookGenre(ctx context.Context, id api.ID) error {
	return fmt.Errorf("not implemented")
}

func (r *Repository) GetUserByID(ctx context.Context, id api.ID) (api.User, error) {
	for _, u := range r.Users {
		if u.ID == id {
			return u, nil
		}
	}
	return api.User{}, fmt.Errorf("get user by id failed: resource not found")
}

func (r *Repository) GetUserByName(ctx context.Context, name string) (api.User, error) {
	for _, u := range r.Users {
		if u.Name == name {
			return u, nil
		}
	}
	return api.User{}, fmt.Errorf("get user by name failed: resource not found")
}
