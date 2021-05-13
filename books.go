package api

type Author struct {
	ID       ID
	Name     string
	NumBooks int
}

type AuthorWithBooks struct {
	Author
	Books []ID
}

type Book struct {
	Author    ID
	BookGenre ID
	ID        ID
	ISBN10    string
	ISBN13    string
	Published Year
	Title     string
}

type BookRating struct {
	Book   ID
	ID     ID
	Rating ThumbRating
	User   ID
}

type BookReading struct {
	Book     ID
	ID       ID
	Started  Time
	Finished Time
	User     ID
}

type BookGenre struct {
	ID       ID
	Name     string
	NumBooks int
}

type BookGenreWithBooks struct {
	BookGenre
	Books []ID
}
