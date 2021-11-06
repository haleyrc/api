package pages

import "github.com/haleyrc/api"

type AuthorGet struct {
	Author api.Author
	Books  []api.Book
}

type AuthorsGet struct {
	Authors      []api.Author
	TotalAuthors uint64
}

type BookGet struct {
	Book api.Book
}

type BookNew struct {
	Authors []api.Author
	Genres  []api.BookGenre
	Formats []api.BookFormat
}

type BooksGet struct {
	Books      []api.Book
	TotalBooks uint
}
