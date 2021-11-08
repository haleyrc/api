package library

type Format string

func (bf Format) Valid() bool {
	switch bf {
	case Hardcover, Paperback, PDF:
		return true
	default:
		return false
	}
}

const (
	Hardcover Format = "Hardcover"
	Paperback Format = "Paperback"
	PDF       Format = "PDF"
)

type Category string

func (bt Category) Valid() bool {
	switch bt {
	case ComicBook, Reference, Novel:
		return true
	default:
		return false
	}
}

type MaybeCategory struct {
	Valid bool
	Value Category
}

const (
	ComicBook Category = "Comic Book"
	Reference Category = "Reference"
	Novel     Category = "Novel"
)
