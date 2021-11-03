package service

import (
	"context"
	"fmt"
	"time"

	"github.com/haleyrc/api"
	"github.com/haleyrc/api/errors"
)

type Tx interface {
	GetAuthor(ctx context.Context, id api.ID) (api.Author, error)
	GetAuthors(ctx context.Context, offset, limit uint) ([]api.Author, error)
	SaveAuthor(ctx context.Context, author api.Author) error
	DeleteAuthor(ctx context.Context, id api.ID) error

	GetBook(ctx context.Context, id api.ID) (api.Book, error)
	GetBooks(ctx context.Context, offset, limit uint) ([]api.Book, error)
	SaveBook(ctx context.Context, book api.Book) error
	DeleteBook(ctx context.Context, id api.ID) error
	AddAuthorToBook(ctx context.Context, book, author api.ID) error
	RateBook(ctx context.Context, user, book api.ID, rating api.Rating) error
	StartBook(ctx context.Context, user, book api.ID, timestamp time.Time) error
	FinishBook(ctx context.Context, user, book api.ID, timestamp time.Time) error

	GetBookGenre(ctx context.Context, id api.ID) (api.BookGenre, error)
	GetBookGenres(ctx context.Context) ([]api.BookGenre, error)
	SaveBookGenre(ctx context.Context, genre api.BookGenre) error
	DeleteBookGenre(ctx context.Context, id api.ID) error
}

type BookService struct {
	DB interface {
		RunInTransaction(context.Context, func(context.Context, Tx) error) error
	}
}

type GetAuthorRequest struct {
	ID api.ID
}

type GetAuthorResponse struct {
	Author api.Author
}

func (s BookService) GetAuthor(ec api.ExecutionContext, req GetAuthorRequest) (*GetAuthorResponse, error) {
	if req.ID == "" {
		return nil, fmt.Errorf("get author failed: %w",
			errors.BadRequest{Message: "ID is required."})
	}

	var resp GetAuthorResponse
	err := s.DB.RunInTransaction(ec.Ctx, func(ctx context.Context, tx Tx) error {
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
	// TODO

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

func (s BookService) GetAuthors(ec api.ExecutionContext, req GetAuthorsRequest) (*GetAuthorsResponse, error) {
	var resp GetAuthorsResponse
	err := s.DB.RunInTransaction(ec.Ctx, func(ctx context.Context, tx Tx) error {
		offset := req.Offset
		limit := clampUint(1, req.Limit, MaxResults)

		var err error
		resp.Authors, err = tx.GetAuthors(ctx, offset, limit)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("get authors failed: %w", err)
	}

	return &resp, nil
}

type SaveAuthorRequest struct {
	ID   api.ID
	Name string
}

type SaveAuthorResponse struct {
	Author api.Author
}

func (s BookService) SaveAuthor(ec api.ExecutionContext, req SaveAuthorRequest) (*SaveAuthorResponse, error) {
	if req.Name == "" {
		return nil, fmt.Errorf("save author failed: %w",
			errors.BadRequest{Message: "Author name can't be blank."})
	}

	author := api.Author{
		ID:   req.ID,
		Name: req.Name,
	}
	if author.ID == "" {
		author.ID = api.NewID()
	}
	err := s.DB.RunInTransaction(ec.Ctx, func(ctx context.Context, tx Tx) error {
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

func (s BookService) DeleteAuthor(ec api.ExecutionContext, req DeleteAuthorRequest) (*DeleteAuthorResponse, error) {
	if req.ID == "" {
		return nil, fmt.Errorf("delete author failed: %w",
			errors.BadRequest{Message: "ID is required."})
	}

	err := s.DB.RunInTransaction(ec.Ctx, func(ctx context.Context, tx Tx) error {
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

func (s BookService) GetBook(ec api.ExecutionContext, req GetBookRequest) (*GetBookResponse, error) {
	if req.ID == "" {
		return nil, fmt.Errorf("get book failed: %w",
			errors.BadRequest{Message: "ID is required."})
	}

	var resp GetBookResponse
	err := s.DB.RunInTransaction(ec.Ctx, func(ctx context.Context, tx Tx) error {
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
	// TODO

	// Pagination
	Offset uint
	Limit  uint
}

type GetBooksResponse struct {
	Books  []api.Book
	Count  uint64
	Offset uint
	Limit  uint
}

func (s BookService) GetBooks(ec api.ExecutionContext, req GetBooksRequest) (*GetBooksResponse, error) {
	var resp GetBooksResponse
	err := s.DB.RunInTransaction(ec.Ctx, func(ctx context.Context, tx Tx) error {
		offset := req.Offset
		limit := clampUint(1, req.Limit, MaxResults)

		var err error
		resp.Books, err = tx.GetBooks(ctx, offset, limit)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("get books failed: %w", err)
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

func (s BookService) SaveBook(ec api.ExecutionContext, req SaveBookRequest) (*SaveBookResponse, error) {
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
			ID:     req.ID.Value,
			Format: req.Format,
			Genre:  req.Genre,
			Title:  req.Title,
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

	err := s.DB.RunInTransaction(ec.Ctx, func(ctx context.Context, tx Tx) error {
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

func (s BookService) DeleteBook(ec api.ExecutionContext, req DeleteBookRequest) (*DeleteBookResponse, error) {
	if req.ID == "" {
		return nil, fmt.Errorf("delete book failed: %w",
			errors.BadRequest{Message: "ID is required."})
	}

	err := s.DB.RunInTransaction(ec.Ctx, func(ctx context.Context, tx Tx) error {
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
}

type RateBookResponse struct{}

func (s BookService) RateBook(ec api.ExecutionContext, req RateBookRequest) (*RateBookResponse, error) {
	if req.Book == "" {
		return nil, fmt.Errorf("rate book failed: %w",
			errors.BadRequest{Message: "Book is required."})
	}
	if req.Rating > 5 {
		return nil, fmt.Errorf("rate book failed: %w",
			errors.BadRequest{Message: "Rating must be an number between 0 and 5."})
	}

	err := s.DB.RunInTransaction(ec.Ctx, func(ctx context.Context, tx Tx) error {
		return tx.RateBook(ctx, ec.User.ID, req.Book, req.Rating)
	})
	if err != nil {
		return nil, fmt.Errorf("rate book failed: %w", err)
	}

	return &RateBookResponse{}, nil
}

type StartBookRequest struct {
	Book api.ID
}

type StartBookResponse struct{}

func (s BookService) StartBook(ec api.ExecutionContext, req StartBookRequest) (*StartBookResponse, error) {
	if req.Book == "" {
		return nil, fmt.Errorf("start book failed: %w",
			errors.BadRequest{Message: "Book is required."})
	}

	err := s.DB.RunInTransaction(ec.Ctx, func(ctx context.Context, tx Tx) error {
		return tx.StartBook(ctx, ec.User.ID, req.Book, time.Now())
	})
	if err != nil {
		return nil, fmt.Errorf("start book failed: %w", err)
	}

	return &StartBookResponse{}, nil
}

type FinishBookRequest struct {
	Book api.ID
}

type FinishBookResponse struct{}

func (s BookService) FinishBook(ec api.ExecutionContext, req FinishBookRequest) (*FinishBookResponse, error) {
	if req.Book == "" {
		return nil, fmt.Errorf("start book failed: %w",
			errors.BadRequest{Message: "Book is required."})
	}

	err := s.DB.RunInTransaction(ec.Ctx, func(ctx context.Context, tx Tx) error {
		return tx.FinishBook(ctx, ec.User.ID, req.Book, time.Now())
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

func (s BookService) GetBookGenre(ec api.ExecutionContext, req GetBookGenreRequest) (*GetBookGenreResponse, error) {
	if req.ID == "" {
		return nil, fmt.Errorf("get book genre failed: %w",
			errors.BadRequest{Message: "ID is required."})
	}

	var resp GetBookGenreResponse
	err := s.DB.RunInTransaction(ec.Ctx, func(ctx context.Context, tx Tx) error {
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

func (s BookService) GetBookGenres(ec api.ExecutionContext, req GetBookGenresRequest) (*GetBookGenresResponse, error) {
	var resp GetBookGenresResponse
	err := s.DB.RunInTransaction(ec.Ctx, func(ctx context.Context, tx Tx) error {
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

func (s BookService) SaveBookGenre(ec api.ExecutionContext, req SaveBookGenreRequest) (*SaveBookGenreResponse, error) {
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
	err := s.DB.RunInTransaction(ec.Ctx, func(ctx context.Context, tx Tx) error {
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

func (s BookService) DeleteBookGenre(ec api.ExecutionContext, req DeleteBookGenreRequest) (*DeleteBookGenreResponse, error) {
	if req.ID == "" {
		return nil, fmt.Errorf("delete book genre failed: %w",
			errors.BadRequest{Message: "ID is required."})
	}

	err := s.DB.RunInTransaction(ec.Ctx, func(ctx context.Context, tx Tx) error {
		return tx.DeleteBookGenre(ctx, req.ID)
	})
	if err != nil {
		return nil, fmt.Errorf("delete book genre failed: %w", err)
	}

	return &DeleteBookGenreResponse{}, nil
}
