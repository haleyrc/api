package mock

import (
	"context"
	"fmt"
	"time"

	"github.com/haleyrc/api"
	"github.com/haleyrc/api/library"
)

type Repository struct {
	authors     []library.Author
	books       []library.Book
	bookAuthors map[api.ID][]api.ID
	bookGenres  []library.Genre
	users       []api.User
}

// func (r *Repository) RunInTransaction(ctx context.Context, f func(ctx context.Context, tx service.Tx) error) error {
// 	return f(ctx, r)
// }

func (r *Repository) GetAuthor(ctx context.Context, id api.ID) (library.Author, error) {
	for _, author := range r.authors {
		if author.ID == id {
			return author, nil
		}
	}
	return library.Author{}, fmt.Errorf("get author failed: resource not found")
}

func (r *Repository) GetAuthors(ctx context.Context, filter library.AuthorsFilter) ([]library.Author, error) {
	authors := []library.Author{}
	for _, author := range r.authors {
		if filter.IDs != nil && len(filter.IDs) > 0 {
			if !author.ID.In(filter.IDs) {
				continue
			}
		}
		authors = append(authors, author)
	}
	return authors, nil
}

func (r *Repository) SaveAuthor(ctx context.Context, author library.Author) error {
	r.authors = append(r.authors, author)
	return nil
}

func (r *Repository) DeleteAuthor(ctx context.Context, id api.ID) error {
	return fmt.Errorf("DeleteAuthor not implemented")
}

func (r *Repository) GetBook(ctx context.Context, id api.ID) (library.Book, error) {
	for _, book := range r.books {
		if book.ID == id {
			return book, nil
		}
	}
	return library.Book{}, fmt.Errorf("get book failed: resource not found")
}

func (r *Repository) GetBooks(ctx context.Context, filter library.BooksFilter) ([]library.Book, uint, error) {
	books := []library.Book{}
	for _, book := range r.books {
		if filter.Author.Valid {
			if !book.HasAuthorWithID(filter.Author.Value) {
				continue
			}
		}
		books = append(books, book)
	}
	return books, uint(len(books)), nil
}

func (r *Repository) SaveBook(ctx context.Context, book library.Book) error {
	r.books = append(r.books, book)
	return nil
}

func (r *Repository) DeleteBook(ctx context.Context, id api.ID) error {
	idx := -1
	for i, book := range r.books {
		if book.ID == id {
			idx = i
			break
		}
	}
	if idx < 0 {
		return fmt.Errorf("delete book failed: resource not found")
	}
	r.books = append(r.books[:idx], r.books[idx+1:]...)
	return nil
}

func (r *Repository) AddAuthorToBook(ctx context.Context, book, author api.ID) error {
	if r.bookAuthors == nil {
		r.bookAuthors = make(map[api.ID][]api.ID)
	}
	r.bookAuthors[book] = append(r.bookAuthors[book], author)
	return nil
}

func (r *Repository) RateBook(ctx context.Context, user, book api.ID, rating api.Rating) error {
	return fmt.Errorf("RateBook not implemented")
}

func (r *Repository) StartBook(ctx context.Context, user, book api.ID, timestamp time.Time) error {
	return fmt.Errorf("StartBook not implemented")
}

func (r *Repository) FinishBook(ctx context.Context, user, book api.ID, timestamp time.Time) error {
	return fmt.Errorf("FinishBook not implemented")
}

func (r *Repository) GetGenre(ctx context.Context, id api.ID) (library.Genre, error) {
	for _, genre := range r.bookGenres {
		if genre.ID == id {
			return genre, nil
		}
	}
	return library.Genre{}, fmt.Errorf("get book genre failed: resource not found")
}

func (r *Repository) GetBookGenre(ctx context.Context, id api.ID) (library.Genre, error) {
	return r.GetGenre(ctx, id)
}

func (r *Repository) GetGenres(ctx context.Context) ([]library.Genre, error) {
	return r.bookGenres, nil
}

func (r *Repository) GetBookGenres(ctx context.Context) ([]library.Genre, error) {
	return r.GetGenres(ctx)
}

func (r *Repository) SaveBookGenre(ctx context.Context, genre library.Genre) error {
	r.bookGenres = append(r.bookGenres, genre)
	return nil
}

func (r *Repository) DeleteBookGenre(ctx context.Context, id api.ID) error {
	return fmt.Errorf("DeleteBookGenre not implemented")
}

func (r *Repository) GetUser(ctx context.Context, query api.UserQuery) (api.User, error) {
	if query.ID == "" && query.Name == "" {
		return api.User{}, fmt.Errorf("get user failed: either id or name must be provided")
	}
	if query.ID != "" {
		return r.getUserByID(ctx, query.ID)
	}
	return r.getUserByName(ctx, query.Name)
}

func (r *Repository) getUserByID(ctx context.Context, id api.ID) (api.User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return u, nil
		}
	}
	return api.User{}, fmt.Errorf("get user by id failed: resource not found")
}

func (r *Repository) getUserByName(ctx context.Context, name string) (api.User, error) {
	for _, u := range r.users {
		if u.Name == name {
			return u, nil
		}
	}
	return api.User{}, fmt.Errorf("get user by name failed: resource not found")
}

func (r *Repository) SaveUser(ctx context.Context, user api.User) error {
	r.users = append(r.users, user)
	return nil
}
