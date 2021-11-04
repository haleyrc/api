package app

import (
	"log"
	"net/http"
	"os"
	"sync"
	"time"

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
		Error    template
		GetBook  template
		GetBooks template
		Login    template
	}

	Books *service.BookService
	Users *service.UserService
}

func (s *Server) All(next ...http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, h := range next {
			h(w, r)
		}
	}
}

func (s *Server) Redirect(code int, path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, path, code)
	}
}

func (s *Server) Render(code int, tmpl template, data interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		scope := scopeFromContext(r.Context())
		tmpl.Render(200, w, map[string]interface{}{
			"User": scope.User,
			"Data": data,
		})
	}
}

func (s *Server) RenderError(code int, msg string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		scope := scopeFromContext(r.Context())
		s.templates.Error.Render(code, w, map[string]interface{}{
			"User":  scope.User,
			"Error": msg,
		})
	}
}

func (s *Server) ClearCookie(key string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:     key,
			Value:    "",
			MaxAge:   -1,
			Secure:   s.secure,
			HttpOnly: true,
			Path:     "/",
		})
	}
}

// Right now this uses default values for everyting but the key and value. There
// might be room down the road to change this to be more customizable, but for
// now this ticks all the boxes.
func (s *Server) SetCookie(key, value string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:     key,
			Value:    value,
			MaxAge:   int(time.Hour.Seconds()),
			Secure:   s.secure,
			HttpOnly: true,
			Path:     "/",
		})
	}
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

	s.Logger.Println("setting up routes...")
	s.router = chi.NewRouter()
	s.router.Use(s.withScope)

	s.router.Get("/", s.GetBooks)

	s.router.Route("/auth", func(r chi.Router) {
		r.Get("/logout", s.DoLogout)
		r.Get("/login", s.GetLogin)
		r.Post("/login", s.DoLogin)
	})

	s.router.Route("/books", func(r chi.Router) {
		r.Get("/", s.GetBooks)
		r.Get("/{id}", s.GetBook)
	})

	s.Logger.Println("initializing templates...")
	s.templates.Error = newTemplate("error")
	s.templates.GetBook = newTemplate("getbook")
	s.templates.GetBooks = newTemplate("getbooks")
	s.templates.Login = newTemplate("login")

	s.initialized = true
}
