package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"github.com/haleyrc/api"
	"github.com/haleyrc/api/html/template"
	"github.com/haleyrc/api/library"
)

func NewLibraryController(
	errors ErrorFunc,
	renderer Renderer,
	library LibraryRepository,
) *LibraryController {
	c := &LibraryController{
		BookPage:    template.New("threepane", "library/library", "library/getbook"),
		BooksPage:   template.New("threepane", "library/library", "library/getbooks"),
		Library:     library,
		NewBookPage: template.New("threepane", "library/library", "library/newbook"),
		Renderer:    renderer,
		ReportError: errors,
	}
	return c
}

type LibraryController struct {
	Renderer

	Library     LibraryRepository
	ReportError ErrorFunc

	BookPage    Template
	BooksPage   Template
	NewBookPage Template
}

func (c *LibraryController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	err := c.Library.DeleteBook(ctx, api.ID(id))
	if err != nil {
		c.ReportError(ctx, err)
		c.Render(ctx, w, 500, c.BooksPage, GetBooksData{
			Error: "Oops, something went wrong!",
		})
		return
	}

	c.Redirect(ctx, w, r, 301, "/library/books")
}

type GetBooksData struct {
	Error      string
	Books      []library.Book
	TotalBooks uint
}

func (c *LibraryController) GetBooks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	books, count, err := c.Library.GetBooks(ctx, library.BooksFilter{})
	if err != nil {
		c.ReportError(ctx, err)
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
	Authors []library.Author
	Book    library.Book
	Genre   library.Genre
}

func (c *LibraryController) GetBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	book, err := c.Library.GetBook(ctx, api.ID(id))
	if err != nil {
		c.ReportError(ctx, err)
		c.Render(ctx, w, 500, c.BookPage, GetBookData{
			Error: "Oops, something went wrong!",
		})
		return
	}

	authors, err := c.Library.GetAuthors(ctx, library.AuthorsFilter{
		IDs: book.Authors,
	})
	if err != nil {
		c.ReportError(ctx, err)
		c.Render(ctx, w, 500, c.BookPage, GetBookData{
			Error: "Oops, something went wrong!",
		})
		return
	}

	genre, err := c.Library.GetGenre(ctx, book.Genre)
	if err != nil {
		c.ReportError(ctx, err)
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
	Error      string
	Authors    []library.Author
	Book       library.Book
	Categories []library.Category
	Formats    []library.Format
	Genres     []library.Genre
}

func (c *LibraryController) NewBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	authors, err := c.Library.GetAuthors(ctx, library.AuthorsFilter{})
	if err != nil {
		c.ReportError(ctx, err)
		c.Render(ctx, w, 500, c.NewBookPage, NewBookData{
			Error: "Oops, something went wrong!",
		})
		return
	}

	genres, err := c.Library.GetGenres(ctx)
	if err != nil {
		c.ReportError(ctx, err)
		c.Render(ctx, w, 500, c.NewBookPage, NewBookData{
			Error: "Oops, something went wrong!",
		})
		return
	}

	c.Render(ctx, w, 200, c.NewBookPage, NewBookData{
		Authors: authors,
		Book: library.Book{
			Title: "Lord Of The Rings",
		},
		Categories: []library.Category{library.ComicBook, library.Novel, library.Reference},
		Formats:    []library.Format{library.Hardcover, library.Paperback, library.PDF},
		Genres:     genres,
	})
}

// TODO: Add multiple errors to template and data
func (c *LibraryController) SaveBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// TODO: Do real sanitization
	// TODO: validation
	authorID := api.ID(strings.TrimSpace(r.PostFormValue("authorID")))
	// TODO: render error
	bookType, _ := parseCategory(r.PostFormValue("type"))
	format := library.Format(strings.TrimSpace(r.PostFormValue("format")))
	genreID := api.ID(strings.TrimSpace(r.PostFormValue("genreID")))
	isbn10 := strings.TrimSpace(r.PostFormValue("isbn10"))
	isbn13 := strings.TrimSpace(r.PostFormValue("isbn13"))
	// TODO: render error
	published, _ := parseInt(r.PostFormValue("published"))
	subtitle := strings.TrimSpace(r.PostFormValue("subtitle"))
	title := strings.TrimSpace(r.PostFormValue("title"))
	// TODO: render error
	volume, _ := parseInt(r.PostFormValue("volume"))

	book := library.Book{
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

	if err := c.Library.SaveBook(ctx, book); err != nil {
		c.ReportError(ctx, err)
		c.Render(ctx, w, 500, c.NewBookPage, NewBookData{
			Book: book,
		})
	}

	c.Redirect(ctx, w, r, 301, "/library/books")
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

func parseCategory(s string) (library.MaybeCategory, error) {
	s = strings.TrimSpace(s)
	// Blank string means we got nothing, which is the value for the "None"
	// dropdown option
	if s == "" {
		return library.MaybeCategory{}, nil
	}
	bookType := library.Category(s)
	if !bookType.Valid() {
		return library.MaybeCategory{}, fmt.Errorf("parseCategory: invalid: %s", s)
	}
	return library.MaybeCategory{
		Valid: true,
		Value: bookType,
	}, nil
}
