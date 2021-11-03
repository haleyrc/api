package app

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/go-chi/chi/v5"

	"github.com/haleyrc/api/service"
)

type Server struct {
	Logger *log.Logger

	initialized bool
	once        sync.Once
	router      *chi.Mux
	secure      bool
	templates   struct {
		GetBook  template
		GetBooks template
		Login    template
	}

	Books *service.BookService
	Users *service.UserService
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.once.Do(s.init)
	s.router.ServeHTTP(w, r)
}

func (s *Server) init() {
	if s.Logger == nil {
		s.Logger = log.New(os.Stderr, "[server] ", log.LstdFlags)
	}
	s.Logger.Println("initializing...")

	s.Logger.Println("Setting up routes...")
	s.router = chi.NewRouter()
	s.router.Use(s.withScope)

	s.router.Get("/", s.booksPage())

	s.router.Route("/auth", func(r chi.Router) {
		r.Get("/logout", s.logout())
		r.Get("/login", s.loginPage())
		r.Post("/login", s.login())
	})

	s.router.Route("/books", func(r chi.Router) {
		r.Get("/", s.booksPage())
		r.Get("/{id}", s.bookPage())
	})

	s.Logger.Println("Initializing templates...")
	s.templates.GetBook = newTemplate("getbook")
	s.templates.GetBooks = newTemplate("getbooks")
	s.templates.Login = newTemplate("login")

	s.initialized = true
}
