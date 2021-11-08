package template

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"path/filepath"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func New(layout string, files ...string) *Template {
	basePath := filepath.Join("templates", "base.tmpl")

	layoutName := layout + ".tmpl"
	layoutPath := filepath.Join("templates", "layouts", layoutName)

	paths := []string{basePath, layoutPath}
	for _, file := range files {
		name := file + ".tmpl"
		path := filepath.Join("templates", "pages", name)
		paths = append(paths, path)
	}

	funcs := template.FuncMap{
		"dump": func(data interface{}) string {
			return spew.Sdump(data)
		},
		"current_year": func() string {
			return fmt.Sprint(time.Now().Year())
		},
	}

	// This is pretty ugly
	t := template.Must(template.New("base.tmpl").Funcs(funcs).ParseFiles(paths...))
	return &Template{
		name:   files[len(files)-1],
		layout: layoutName,
		t:      t,
	}
}

type Template struct {
	name   string
	layout string
	t      *template.Template
}

func (t *Template) Execute(ctx context.Context, w io.Writer, data interface{}) error {
	if err := t.t.ExecuteTemplate(w, "base.tmpl", data); err != nil {
		return fmt.Errorf("execute failed: %s: %w", t.name, err)
	}
	return nil
}
