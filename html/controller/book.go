package controller

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"github.com/haleyrc/api"
	"github.com/haleyrc/api/html/controller/mw"
	"github.com/haleyrc/api/html/template"
)

type Template interface {
	Execute(ctx context.Context, w io.Writer, data interface{}) error
}

type LibraryRepository interface {
	DeleteBook(ctx context.Context, id api.ID) error
	GetAuthors(ctx context.Context, filter api.AuthorsFilter) ([]api.Author, error)
	GetBook(ctx context.Context, id api.ID) (api.Book, error)
	GetBooks(ctx context.Context, filter api.BooksFilter) ([]api.Book, uint, error)
	GetGenre(ctx context.Context, id api.ID) (api.BookGenre, error)
	GetGenres(ctx context.Context) ([]api.BookGenre, error)
	SaveBook(ctx context.Context, book api.Book) error
}

func NewLibraryController(repo LibraryRepository) *LibraryController {
	c := &LibraryController{
		LibraryRepo: repo,
		BookPage:    template.New("threepane", "library/library", "library/getbook"),
		BooksPage:   template.New("threepane", "library/library", "library/getbooks"),
		NewBookPage: template.New("threepane", "library/library", "library/newbook"),
	}
	return c
}

type LibraryController struct {
	LibraryRepo LibraryRepository

	BookPage    Template
	BooksPage   Template
	NewBookPage Template
}

func (c *LibraryController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	err := c.LibraryRepo.DeleteBook(ctx, api.ID(id))
	if err != nil {
		c.reportError(ctx, err)
		c.Render(ctx, w, 500, c.BooksPage, GetBooksData{
			Error: "Oops, something went wrong!",
		})
		return
	}

	c.Redirect(ctx, w, r, 301, "/library/books")
}

type GetBooksData struct {
	Error      string
	Books      []api.Book
	TotalBooks uint
}

func (c *LibraryController) GetBooks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	books, count, err := c.LibraryRepo.GetBooks(ctx, api.BooksFilter{})
	if err != nil {
		c.reportError(ctx, err)
		c.Render(ctx, w, 500, c.BooksPage, GetBooksData{
			Error: "Oops, something went wrong!",
		})
		return
	}

	c.Render(ctx, w, 200, c.BooksPage, GetBooksData{
		Books:      books,
		TotalBooks: count,
	})
}

type GetBookData struct {
	Error   string
	Authors []api.Author
	Book    api.Book
	Genre   api.BookGenre
}

func (c *LibraryController) GetBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	book, err := c.LibraryRepo.GetBook(ctx, api.ID(id))
	if err != nil {
		c.reportError(ctx, err)
		c.Render(ctx, w, 500, c.BookPage, GetBookData{
			Error: "Oops, something went wrong!",
		})
		return
	}

	authors, err := c.LibraryRepo.GetAuthors(ctx, api.AuthorsFilter{
		IDs: book.Authors,
	})
	if err != nil {
		c.reportError(ctx, err)
		c.Render(ctx, w, 500, c.BookPage, GetBookData{
			Error: "Oops, something went wrong!",
		})
		return
	}

	genre, err := c.LibraryRepo.GetGenre(ctx, book.Genre)
	if err != nil {
		c.reportError(ctx, err)
		c.Render(ctx, w, 500, c.BookPage, GetBookData{
			Error: "Oops, something went wrong!",
		})
		return
	}

	c.Render(ctx, w, 200, c.BookPage, GetBookData{
		Authors: authors,
		Book:    book,
		Genre:   genre,
	})
}

type NewBookData struct {
	Error   string
	Authors []api.Author
	Book    api.Book
	Formats []api.BookFormat
	Genres  []api.BookGenre
	Types   []api.BookType
}

func (c *LibraryController) NewBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	authors, err := c.LibraryRepo.GetAuthors(ctx, api.AuthorsFilter{})
	if err != nil {
		c.reportError(ctx, err)
		c.Render(ctx, w, 500, c.NewBookPage, NewBookData{
			Error: "Oops, something went wrong!",
		})
		return
	}

	genres, err := c.LibraryRepo.GetGenres(ctx)
	if err != nil {
		c.reportError(ctx, err)
		c.Render(ctx, w, 500, c.NewBookPage, NewBookData{
			Error: "Oops, something went wrong!",
		})
		return
	}

	c.Render(ctx, w, 200, c.NewBookPage, NewBookData{
		Authors: authors,
		Book: api.Book{
			Title: "Lord Of The Rings",
		},
		Formats: []api.BookFormat{api.Hardcover, api.Paperback, api.PDF},
		Genres:  genres,
		Types:   []api.BookType{api.ComicBook, api.Novel, api.Reference},
	})
}

// TODO: Add multiple errors to template and data
func (c *LibraryController) SaveBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// TODO: Do real sanitization
	// TODO: validation
	authorID := api.ID(strings.TrimSpace(r.PostFormValue("authorID")))
	// TODO: render error
	bookType, _ := parseBookType(r.PostFormValue("type"))
	format := api.BookFormat(strings.TrimSpace(r.PostFormValue("format")))
	genreID := api.ID(strings.TrimSpace(r.PostFormValue("genreID")))
	isbn10 := strings.TrimSpace(r.PostFormValue("isbn10"))
	isbn13 := strings.TrimSpace(r.PostFormValue("isbn13"))
	// TODO: render error
	published, _ := parseInt(r.PostFormValue("published"))
	subtitle := strings.TrimSpace(r.PostFormValue("subtitle"))
	title := strings.TrimSpace(r.PostFormValue("title"))
	// TODO: render error
	volume, _ := parseInt(r.PostFormValue("volume"))

	book := api.Book{
		ID:      api.NewID(),
		Genre:   genreID,
		Authors: []api.ID{authorID},
		Format:  format,
		Title:   title,
		// Anthology: ...,
		ISBN10: api.MaybeString{
			Value: isbn10,
			Valid: isbn10 != "",
		},
		ISBN13: api.MaybeString{
			Value: isbn13,
			Valid: isbn13 != "",
		},
		Published: published,
		// Publisher: ...,
		Subtitle: api.MaybeString{
			Value: subtitle,
			Valid: subtitle != "",
		},
		Type:   bookType,
		Volume: volume,
	}

	// TODO: If there are already errors, just render the page with the book
	// object

	if err := c.LibraryRepo.SaveBook(ctx, book); err != nil {
		c.reportError(ctx, err)
		c.Render(ctx, w, 500, c.NewBookPage, NewBookData{
			Book: book,
		})
	}

	c.Redirect(ctx, w, r, 301, "/library/books")
}

func (c *LibraryController) Redirect(ctx context.Context, w http.ResponseWriter, r *http.Request, code int, path string) {
	http.Redirect(w, r, path, code)
}

func (c *LibraryController) Render(ctx context.Context, w http.ResponseWriter, code int, t Template, data interface{}) {
	w.WriteHeader(code)
	user := mw.UserFromContext(ctx)
	err := t.Execute(ctx, w, map[string]interface{}{
		"User": user,
		"Data": data,
	})
	if err != nil {
		c.reportError(ctx, fmt.Errorf("render failed: %w", err))
	}
}

func (c *LibraryController) reportError(ctx context.Context, err error) {
	fmt.Println("[ERROR]", err)
}

func parseInt(s string) (api.MaybeInt, error) {
	s = strings.TrimSpace(s)
	// Blank string means the client sent nothing, therefore the Maybe is a No
	if s == "" {
		return api.MaybeInt{}, nil
	}
	// If we can't parse the value, then the client sent something janky
	i, err := strconv.Atoi(s)
	if err != nil {
		return api.MaybeInt{}, fmt.Errorf("parseInt: %w", err)
	}
	// If we can parse the integer, then it was provided and is valid
	return api.MaybeInt{
		Valid: true,
		Value: i,
	}, nil
}

func parseBookType(s string) (api.MaybeBookType, error) {
	s = strings.TrimSpace(s)
	// Blank string means we got nothing, which is the value for the "None"
	// dropdown option
	if s == "" {
		return api.MaybeBookType{}, nil
	}
	bookType := api.BookType(s)
	if !bookType.Valid() {
		return api.MaybeBookType{}, fmt.Errorf("parseBookType: invalid: %s", s)
	}
	return api.MaybeBookType{
		Valid: true,
		Value: bookType,
	}, nil
}
