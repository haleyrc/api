package main

import (
	"log"
	"net/http"

	"github.com/haleyrc/api"
	"github.com/haleyrc/api/app"
	"github.com/haleyrc/api/internal/mock"
	"github.com/haleyrc/api/service"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	bookGenres := []api.BookGenre{
		{ID: api.NewID(), Name: "Adventure"},
	}
	repo := &mock.Repository{
		Books: []api.Book{
			{ID: api.NewID(), Genre: api.BookGenre{ID: bookGenres[0].ID}, Title: "Moby Dick", Format: api.Hardcover},
		},
		BookGenres: bookGenres,
		Users: []api.User{
			{ID: api.NewID(), Name: "ryan"},
		},
	}
	var server app.Server
	server.Books = &service.BookService{DB: repo}
	server.Users = &service.UserService{DB: repo}
	log.Println("listening on :8080...")
	return http.ListenAndServe(":8080", &server)
}
