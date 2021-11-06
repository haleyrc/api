package service

import (
	"context"
	"fmt"
	"time"

	"github.com/haleyrc/api"
	"github.com/haleyrc/api/errors"
)

type BookService struct {
	DB Database
}

type GetAuthorRequest struct {
	ID api.ID
}

type GetAuthorResponse struct {
	Author api.Author
}

func (s BookService) GetAuthor(ctx context.Context, req GetAuthorRequest) (*GetAuthorResponse, error) {
	if req.ID == "" {
		return nil, fmt.Errorf("get author failed: %w",
			errors.BadRequest{Message: "ID is required."})
	}

	var resp GetAuthorResponse
	err := s.DB.RunInTransaction(ctx, func(ctx context.Context, tx Tx) error {
		var err error
		resp.Author, err = tx.GetAuthor(ctx, req.ID)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("get author failed: %w", err)
	}

	return &resp, nil
}

type GetAuthorsRequest struct {
	// Filters
	IDs []api.ID

	// Pagination
	Offset uint
	Limit  uint
}

type GetAuthorsResponse struct {
	Authors []api.Author
	Count   uint64
	Offset  uint
	Limit   uint
}

func (s BookService) GetAuthors(ctx context.Context, req GetAuthorsRequest) (*GetAuthorsResponse, error) {
	var resp GetAuthorsResponse
	err := s.DB.RunInTransaction(ctx, func(ctx context.Context, tx Tx) error {
		var err error
		resp.Authors, err = tx.GetAuthors(ctx, api.AuthorsFilter{
			IDs: req.IDs,
		})
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("get authors failed: %w", err)
	}

	return &resp, nil
}

type SaveAuthorRequest struct {
	ID   api.MaybeID
	Name string
}

type SaveAuthorResponse struct {
	Author api.Author
}

func (s BookService) SaveAuthor(ctx context.Context, req SaveAuthorRequest) (*SaveAuthorResponse, error) {
	if req.Name == "" {
		return nil, fmt.Errorf("save author failed: %w",
			errors.BadRequest{Message: "Author name can't be blank."})
	}

	author := api.Author{
		ID:   req.ID.Value,
		Name: req.Name,
	}
	if author.ID == "" {
		author.ID = api.NewID()
	}
	err := s.DB.RunInTransaction(ctx, func(ctx context.Context, tx Tx) error {
		return tx.SaveAuthor(ctx, author)
	})
	if err != nil {
		return nil, fmt.Errorf("save author failed: %w", err)
	}

	return &SaveAuthorResponse{Author: author}, nil
}

type DeleteAuthorRequest struct {
	ID api.ID
}

type DeleteAuthorResponse struct{}

func (s BookService) DeleteAuthor(ctx context.Context, req DeleteAuthorRequest) (*DeleteAuthorResponse, error) {
	if req.ID == "" {
		return nil, fmt.Errorf("delete author failed: %w",
			errors.BadRequest{Message: "ID is required."})
	}

	err := s.DB.RunInTransaction(ctx, func(ctx context.Context, tx Tx) error {
		return tx.DeleteAuthor(ctx, req.ID)
	})
	if err != nil {
		return nil, fmt.Errorf("delete author failed: %w", err)
	}

	return &DeleteAuthorResponse{}, nil
}

type GetBookRequest struct {
	ID api.ID
}

type GetBookResponse struct {
	Book api.Book
}

func (s BookService) GetBook(ctx context.Context, req GetBookRequest) (*GetBookResponse, error) {
	if req.ID == "" {
		return nil, fmt.Errorf("get book failed: %w",
			errors.BadRequest{Message: "ID is required."})
	}

	var resp GetBookResponse
	err := s.DB.RunInTransaction(ctx, func(ctx context.Context, tx Tx) error {
		var err error
		resp.Book, err = tx.GetBook(ctx, req.ID)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("get book failed: %w", err)
	}

	return &resp, nil
}

type GetBooksRequest struct {
	// Filters
	Author api.MaybeID

	// Pagination
	Offset uint
	Limit  uint
}

type GetBooksResponse struct {
	Books  []api.Book
	Count  uint
	Offset uint
	Limit  uint
}

func (s BookService) GetBooks(ctx context.Context, req GetBooksRequest) (*GetBooksResponse, error) {
	offset := req.Offset
	limit := clampUint(1, req.Limit, MaxResults)

	var books []api.Book
	var count uint
	err := s.DB.RunInTransaction(ctx, func(ctx context.Context, tx Tx) error {
		var err error
		books, count, err = tx.GetBooks(ctx, api.BooksFilter{
			Author: req.Author,
		})
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("get books failed: %w", err)
	}

	resp := GetBooksResponse{
		Books:  books,
		Count:  count,
		Offset: offset,
		Limit:  limit,
	}

	return &resp, nil
}

type SaveBookRequest struct {
	// Required

	Author api.ID
	Genre  api.ID
	Format api.BookFormat
	Title  string

	// Optional

	ID api.MaybeID

	Anthology api.MaybeID
	ISBN10    api.MaybeString
	ISBN13    api.MaybeString
	Published api.MaybeInt
	Publisher api.MaybeID
	Subtitle  api.MaybeString
	Type      api.MaybeBookType
	Volume    api.MaybeInt
}

type SaveBookResponse struct {
	Book api.Book
}

func (s BookService) SaveBook(ctx context.Context, req SaveBookRequest) (*SaveBookResponse, error) {
	if req.Author == "" {
		return nil, fmt.Errorf("save book failed: %w",
			errors.BadRequest{Message: "Author is required."})
	}
	if req.Genre == "" {
		return nil, fmt.Errorf("save book failed: %w",
			errors.BadRequest{Message: "Genre is required."})
	}
	if req.Format == "" {
		return nil, fmt.Errorf("save book failed: %w",
			errors.BadRequest{Message: "Book format is required."})
	}
	if !req.Format.Valid() {
		return nil, fmt.Errorf("save book failed: %w",
			errors.BadRequest{Message: fmt.Sprintf("Invalid book format: %s.", req.Format)})
	}
	if req.Title == "" {
		return nil, fmt.Errorf("save book failed: %w",
			errors.BadRequest{Message: "Book title is required."})
	}

	resp := SaveBookResponse{
		Book: api.Book{
			// Required
			ID:      req.ID.Value,
			Format:  req.Format,
			Genre:   req.Genre,
			Authors: []api.ID{req.Author},
			Title:   req.Title,
			// Optional
			Anthology: req.Anthology,
			ISBN10:    req.ISBN10,
			ISBN13:    req.ISBN13,
			Published: req.Published,
			Publisher: req.Publisher,
			Subtitle:  req.Subtitle,
			Type:      req.Type,
			Volume:    req.Volume,
		},
	}
	if resp.Book.ID == "" {
		resp.Book.ID = api.NewID()
	}

	err := s.DB.RunInTransaction(ctx, func(ctx context.Context, tx Tx) error {
		if err := tx.SaveBook(ctx, resp.Book); err != nil {
			return err
		}
		if err := tx.AddAuthorToBook(ctx, resp.Book.ID, req.Author); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("save book failed: %w", err)
	}

	return &resp, nil
}

type DeleteBookRequest struct {
	ID api.ID
}

type DeleteBookResponse struct{}

func (s BookService) DeleteBook(ctx context.Context, req DeleteBookRequest) (*DeleteBookResponse, error) {
	if req.ID == "" {
		return nil, fmt.Errorf("delete book failed: %w",
			errors.BadRequest{Message: "ID is required."})
	}

	err := s.DB.RunInTransaction(ctx, func(ctx context.Context, tx Tx) error {
		return tx.DeleteBook(ctx, req.ID)
	})
	if err != nil {
		return nil, fmt.Errorf("delete book failed: %w", err)
	}

	return &DeleteBookResponse{}, nil
}

type RateBookRequest struct {
	Book   api.ID
	Rating api.Rating
	User   api.ID
}

type RateBookResponse struct{}

func (s BookService) RateBook(ctx context.Context, req RateBookRequest) (*RateBookResponse, error) {
	if req.Book == "" {
		return nil, fmt.Errorf("rate book failed: %w",
			errors.BadRequest{Message: "Book is required."})
	}
	if req.Rating > 5 {
		return nil, fmt.Errorf("rate book failed: %w",
			errors.BadRequest{Message: "Rating must be an number between 0 and 5."})
	}

	err := s.DB.RunInTransaction(ctx, func(ctx context.Context, tx Tx) error {
		return tx.RateBook(ctx, req.User, req.Book, req.Rating)
	})
	if err != nil {
		return nil, fmt.Errorf("rate book failed: %w", err)
	}

	return &RateBookResponse{}, nil
}

type StartBookRequest struct {
	Book api.ID
	User api.ID
}

type StartBookResponse struct{}

func (s BookService) StartBook(ctx context.Context, req StartBookRequest) (*StartBookResponse, error) {
	if req.Book == "" {
		return nil, fmt.Errorf("start book failed: %w",
			errors.BadRequest{Message: "Book is required."})
	}

	err := s.DB.RunInTransaction(ctx, func(ctx context.Context, tx Tx) error {
		return tx.StartBook(ctx, req.User, req.Book, time.Now())
	})
	if err != nil {
		return nil, fmt.Errorf("start book failed: %w", err)
	}

	return &StartBookResponse{}, nil
}

type FinishBookRequest struct {
	Book api.ID
	User api.ID
}

type FinishBookResponse struct{}

func (s BookService) FinishBook(ctx context.Context, req FinishBookRequest) (*FinishBookResponse, error) {
	if req.Book == "" {
		return nil, fmt.Errorf("start book failed: %w",
			errors.BadRequest{Message: "Book is required."})
	}

	err := s.DB.RunInTransaction(ctx, func(ctx context.Context, tx Tx) error {
		return tx.FinishBook(ctx, req.User, req.Book, time.Now())
	})
	if err != nil {
		return nil, fmt.Errorf("start book failed: %w", err)
	}

	return &FinishBookResponse{}, nil
}

type GetBookGenreRequest struct {
	ID api.ID
}

type GetBookGenreResponse struct {
	Genre api.BookGenre
}

func (s BookService) GetBookGenre(ctx context.Context, req GetBookGenreRequest) (*GetBookGenreResponse, error) {
	if req.ID == "" {
		return nil, fmt.Errorf("get book genre failed: %w",
			errors.BadRequest{Message: "ID is required."})
	}

	var resp GetBookGenreResponse
	err := s.DB.RunInTransaction(ctx, func(ctx context.Context, tx Tx) error {
		var err error
		resp.Genre, err = tx.GetBookGenre(ctx, req.ID)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("get book genre failed: %w", err)
	}

	return &resp, nil
}

// We don't do pagination for book genres because the assumption is that this
// list won't grow without bound.
type GetBookGenresRequest struct {
	// Filters
	// TODO
}

type GetBookGenresResponse struct {
	Genres []api.BookGenre
}

func (s BookService) GetBookGenres(ctx context.Context, req GetBookGenresRequest) (*GetBookGenresResponse, error) {
	var resp GetBookGenresResponse
	err := s.DB.RunInTransaction(ctx, func(ctx context.Context, tx Tx) error {
		var err error
		resp.Genres, err = tx.GetBookGenres(ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("get book genres failed: %w", err)
	}

	return &resp, nil
}

type SaveBookGenreRequest struct {
	ID   api.ID
	Name string
}

type SaveBookGenreResponse struct {
	Genre api.BookGenre
}

func (s BookService) SaveBookGenre(ctx context.Context, req SaveBookGenreRequest) (*SaveBookGenreResponse, error) {
	if req.Name == "" {
		return nil, fmt.Errorf("save book genre failed: %w",
			errors.BadRequest{Message: "Genre name can't be blank."})
	}

	genre := api.BookGenre{
		ID:   req.ID,
		Name: req.Name,
	}
	if genre.ID == "" {
		genre.ID = api.NewID()
	}
	err := s.DB.RunInTransaction(ctx, func(ctx context.Context, tx Tx) error {
		return tx.SaveBookGenre(ctx, genre)
	})
	if err != nil {
		return nil, fmt.Errorf("save book genre failed: %w", err)
	}

	return &SaveBookGenreResponse{Genre: genre}, nil
}

type DeleteBookGenreRequest struct {
	ID api.ID
}

type DeleteBookGenreResponse struct{}

func (s BookService) DeleteBookGenre(ctx context.Context, req DeleteBookGenreRequest) (*DeleteBookGenreResponse, error) {
	if req.ID == "" {
		return nil, fmt.Errorf("delete book genre failed: %w",
			errors.BadRequest{Message: "ID is required."})
	}

	err := s.DB.RunInTransaction(ctx, func(ctx context.Context, tx Tx) error {
		return tx.DeleteBookGenre(ctx, req.ID)
	})
	if err != nil {
		return nil, fmt.Errorf("delete book genre failed: %w", err)
	}

	return &DeleteBookGenreResponse{}, nil
}
