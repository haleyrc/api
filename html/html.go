package html

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/haleyrc/api/html/controller/mw"
)

type Template interface {
	Execute(context.Context, io.Writer, interface{}) error
}

type Renderer struct {
	ReportError func(ctx context.Context, err error) string
}

func (rend Renderer) Redirect(ctx context.Context, w http.ResponseWriter, r *http.Request, code int, path string) {
	http.Redirect(w, r, path, code)
}

func (rend Renderer) Render(ctx context.Context, w http.ResponseWriter, code int, t Template, data interface{}) {
	w.WriteHeader(code)
	user := mw.UserFromContext(ctx)
	err := t.Execute(ctx, w, map[string]interface{}{
		"User": user,
		"Data": data,
	})
	if err != nil {
		rend.ReportError(ctx, fmt.Errorf("render failed: %w", err))
	}
}
