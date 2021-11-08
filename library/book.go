package library

import (
	"github.com/haleyrc/api"
)

type BooksFilter struct {
	Author api.MaybeID
}

type Book struct {
	ID        api.ID
	Genre     api.ID
	Authors   []api.ID
	Format    Format
	Title     string
	Anthology api.MaybeID
	ISBN10    api.MaybeString
	ISBN13    api.MaybeString
	Published api.MaybeInt
	Publisher api.MaybeID
	Subtitle  api.MaybeString
	Type      MaybeCategory
	Volume    api.MaybeInt
}

func (b Book) HasAuthorWithID(id api.ID) bool {
	for _, value := range b.Authors {
		if value == id {
			return true
		}
	}
	return false
}
