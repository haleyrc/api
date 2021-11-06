package main

import (
	"context"
	"log"
	"net/http"

	"github.com/haleyrc/api"
	"github.com/haleyrc/api/app"
	"github.com/haleyrc/api/html/controller"
	"github.com/haleyrc/api/internal/mock"
	"github.com/haleyrc/api/service"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	ctx := context.Background()

	repo := &mock.Repository{}
	seed(ctx, repo)

	var server app.Server
	server.Books = &service.BookService{DB: repo}
	server.Users = &service.UserService{DB: repo}
	server.Library = controller.NewLibraryController(repo)

	log.Println("listening on :8080...")
	return http.ListenAndServe(":8080", &server)
}

func seed(ctx context.Context, repo *mock.Repository) {
	bookGenres := []api.BookGenre{
		{ID: api.NewID(), Name: "Adventure"},
	}
	for _, genre := range bookGenres {
		repo.SaveBookGenre(ctx, genre)
	}

	authors := []api.Author{
		{ID: api.NewID(), Name: "Herman Melville"},
	}
	for _, author := range authors {
		repo.SaveAuthor(ctx, author)
	}

	books := []api.Book{{
		ID:      api.NewID(),
		Authors: []api.ID{authors[0].ID},
		Genre:   bookGenres[0].ID,
		Title:   "Moby Dick",
		Format:  api.Hardcover,
	}}
	for _, book := range books {
		repo.SaveBook(ctx, book)
		repo.AddAuthorToBook(ctx, book.ID, authors[0].ID)
	}

	users := []api.User{
		{ID: api.NewID(), Name: "ryan"},
	}
	for _, user := range users {
		repo.SaveUser(ctx, user)
	}
}
