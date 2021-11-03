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

type Author struct {
	ID ID

	// Required

	Name string

	// Computed

	NumBooks int
}

type Book struct {
	ID ID

	// Foreign Keys
	Genre BookGenre

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

	// Authors      []ID
	// Illustrators []ID
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
