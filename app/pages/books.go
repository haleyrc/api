package pages

import (
	"io"
	"text/template"
)

type GetBooks struct {
	Page

	tmpl *template.Template
}

func (p GetBooks) Render(w io.Writer, code int) error {
	var err error
	p.once.Do(func() {
		p.tmpl = template.New("")
	})
}
