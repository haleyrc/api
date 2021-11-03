package app

import (
	"net/http"

	"github.com/haleyrc/api/app/pages"
)

func (s *Server) getBooks(w http.ResponseWriter, r *http.Request) {
	var p pages.GetBooks
	if err := p.Render(w, 200); err != nil {
		s.Logger.Println(err)
	}
}
