package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Logger *log.Logger

	initialized bool
	once        sync.Once
	router      *chi.Mux
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.once.Do(s.init)
	if !s.initialized {
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprint(w, http.StatusText(http.StatusNotImplemented))
	}
	s.router.ServeHTTP(w, r)
}

func (s *Server) init() {
	if s.Logger == nil {
		s.Logger = log.New(os.Stderr, "[server] ", log.LstdFlags)
	}
	s.Logger.Println("initializing...")

	s.router = chi.NewRouter()
	s.router.Get("/books", s.getBooks)

	s.initialized = true
}
