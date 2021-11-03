package pages

import "io"

type Page interface {
	Render(w io.Writer, code int) error
}
