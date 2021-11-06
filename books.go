package api

import (
	"time"
)

type BookFormat string

func (bf BookFormat) Valid() bool {
	switch bf {
	case Hardcover, Paperback, PDF:
		return true
	default:
		return false
	}
}

const (
	Hardcover BookFormat = "Hardcover"
	Paperback BookFormat = "Paperback"
	PDF       BookFormat = "PDF"
)

type BookType string

func (bt BookType) Valid() bool {
	switch bt {
	case ComicBook, Reference, Novel:
		return true
	default:
		return false
	}
}

const (
	ComicBook BookType = "Comic Book"
	Reference BookType = "Reference"
	Novel     BookType = "Novel"
)

type Anthology struct {
	ID ID

	// Required

	Name string
}

type AuthorsFilter struct {
	IDs []ID
}

type Author struct {
	ID ID

	// Required

	Name string

	// Computed

	NumBooks int
}

type BooksFilter struct {
	Author MaybeID
}

type Book struct {
	ID ID

	// Foreign Keys
	Genre   ID
	Authors []ID

	// Required

	Format BookFormat
	Title  string

	// Optional

	Anthology MaybeID
	ISBN10    MaybeString
	ISBN13    MaybeString
	Published MaybeInt
	Publisher MaybeID
	Subtitle  MaybeString
	Type      MaybeBookType
	Volume    MaybeInt
}

func (b Book) HasAuthorWithID(id ID) bool {
	for _, value := range b.Authors {
		if value == id {
			return true
		}
	}
	return false
}

type BookRating struct {
	ID ID

	// Foreign Keys

	Book ID
	User ID

	Rating Rating
}

type BookReading struct {
	ID ID

	// Foreign Keys

	Book ID
	User ID

	Started  time.Time
	Finished time.Time
}

type BookGenre struct {
	ID ID

	// Required

	Name string

	// Computed

	NumBooks int
}
