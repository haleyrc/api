package app

import (
	"fmt"
	gotemplate "html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
)

func dump(data interface{}) string {
	return spew.Sdump(data)
}

func newTemplate(name string) template {
	funcs := gotemplate.FuncMap{
		"dump": dump,
	}

	filename := name + ".tmpl"
	path := filepath.Join("templates", "pages", filename)
	t, err := gotemplate.New(filename).Funcs(funcs).ParseFiles(path, "templates/layouts/threepane.tmpl", "templates/base.tmpl")
	if err != nil {
		fmt.Printf("failed to parse template %s: %v\n", name, err)
		os.Exit(1)
	}

	return template{name: filename, t: t}
}

type template struct {
	name string
	t    *gotemplate.Template
}

func (t template) Render(code int, w http.ResponseWriter, pageData interface{}) {
	w.WriteHeader(code)
	if err := t.t.ExecuteTemplate(w, "base.tmpl", pageData); err != nil {
		fmt.Printf("failed to render template %s: %v\n", t.name, err)
	}
}
