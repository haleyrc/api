package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/haleyrc/api"
	"github.com/haleyrc/api/html"
	"github.com/haleyrc/api/html/controller"
	"github.com/haleyrc/api/html/controller/mw"
	"github.com/haleyrc/api/html/cookies"
	"github.com/haleyrc/api/internal/mock"
	"github.com/haleyrc/api/log"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	ctx := context.Background()
	log.SetLevel(log.DebugLevel)
	log.SetHandler(log.CLIHandler)

	repo := &mock.Repository{}
	seed(ctx, repo)

	cookies := cookies.Jar{}
	errorFunc := func(ctx context.Context, err error) string {
		id := api.NewID()
		log.WithError(err).WithField("id", id).Error("unexpected error")
		return string(id)
	}
	renderer := html.Renderer{
		ReportError: errorFunc,
	}

	router := chi.NewRouter()
	router.Use(mw.Authenticate(cookies, repo, errorFunc))

	publicC := controller.NewPublicController(errorFunc, renderer)
	router.Get("/", publicC.Index)

	router.Route("/auth", func(r chi.Router) {
		authC := controller.NewAuthController(cookies, errorFunc, renderer, repo)
		r.Get("/logout", authC.Logout)
		r.Get("/login", authC.Login)
		r.Post("/login", authC.Authenticate)
	})

	router.Route("/library/books", func(r chi.Router) {
		libraryC := controller.NewLibraryController(errorFunc, renderer, repo)
		r.Use(mw.RequireAuthenticated("/auth/login"))
		r.Get("/new", libraryC.NewBook)
		r.Post("/{id}/delete", libraryC.DeleteBook)
		r.Get("/{id}", libraryC.GetBook)
		r.Post("/", libraryC.SaveBook)
		r.Get("/", libraryC.GetBooks)
	})

	log.WithField("port", ":8080").Debug("listening")
	return http.ListenAndServe(":8080", router)
}

func seed(ctx context.Context, repo *mock.Repository) {
	bookGenres := []api.BookGenre{
		{ID: api.NewID(), Name: "Adventure"},
	}
	for _, genre := range bookGenres {
		repo.SaveBookGenre(ctx, genre)
	}

	authors := []api.Author{
		{ID: "9fccbac5-562c-4af3-9a95-1556b9d54e10", Name: "Herman Melville"},
	}
	for _, author := range authors {
		repo.SaveAuthor(ctx, author)
	}

	books := []api.Book{{
		ID:      "44e3befa-b9ea-4d95-a4b3-dbbce9771e49",
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
		{ID: "c27d921c-e648-46c7-bc6e-96e9224d6e68", Name: "ryan"},
	}
	for _, user := range users {
		repo.SaveUser(ctx, user)
	}
}
