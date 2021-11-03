package app

import (
	"fmt"
	gotemplate "html/template"
	"net/http"
)

func newTemplate(name string) template {
	t, err := gotemplate.ParseFiles("templates/layouts/base.tmpl", "templates/pages/"+name+".tmpl")
	if err != nil {
		panic(fmt.Errorf("failed to parse template %s: %v", name, err))
	}
	return template{
		name: name,
		t:    t,
	}
}

type template struct {
	name string
	t    *gotemplate.Template
}

func (t template) Render(code int, w http.ResponseWriter, pageData interface{}) {
	w.WriteHeader(code)
	if err := t.t.Execute(w, pageData); err != nil {
		panic(fmt.Errorf("failed to render template %s: %v", t.name, err))
	}
}
