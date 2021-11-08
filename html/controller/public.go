package controller

import (
	"net/http"

	"github.com/haleyrc/api/html/template"
)

func NewPublicController(errors ErrorFunc, renderer Renderer) *PublicController {
	return &PublicController{
		IndexPage:   template.New("fullscreen", "public/index"),
		Renderer:    renderer,
		ReportError: errors,
	}
}

type PublicController struct {
	Renderer

	ReportError ErrorFunc

	IndexPage Template
}

func (c *PublicController) Index(w http.ResponseWriter, r *http.Request) {
	c.Renderer.Render(r.Context(), w, 200, c.IndexPage, nil)
}
