package api

import "time"

type MovieFormat string

const (
	DVD    MovieFormat = "dvd"
	BluRay MovieFormat = "blu-ray"
	VHS    MovieFormat = "vhs"
)

type Actor struct {
	ID        ID
	Name      string
	NumMovies int
}

type ActorWithMovies struct {
	Actor
	Movies []ID
}

type Director struct {
	ID        ID
	Name      string
	NumMovies int
}

type DirectorWithMovies struct {
	Movies []ID
}

type MovieGenre struct {
	ID        ID
	Name      string
	NumMovies int
}

type MovieGenreWithMovies struct {
	MovieGenre
	Movies []ID
}

type Movie struct {
	Formats    []MovieFormat
	Director   ID
	MovieGenre ID
	ID         ID
	Owned      bool
	Released   int
	Title      string
	NumWatches int
}

type MovieWithActors struct {
	Movie
	Actors []ID
}

type MovieWithWatchers struct {
	Movie
	Watchers []ID
}

type MovieRating struct {
	ID     ID
	Movie  ID
	Rating Rating
	User   ID
}

type MovieViewing struct {
	Date  time.Time
	ID    ID
	Movie ID
	Users []ID
}

type GetActorRequest struct {
	ID ID
}

type GetActorResponse struct {
	Actor ActorWithMovies
}

type GetActorsRequest struct {
	// Filters
	// TODO

	// Pagination
	Offset uint
	Limit  uint
}

type GetActorsResponse struct {
	Actors []Actor
	Count  uint64
	Offset uint
	Limit  uint
}

type SaveActorRequest struct {
	ID   ID
	Name string
}

type SaveActorResponse struct {
	Actor Actor
}

type DeleteActorRequest struct {
	ID ID
}

type DeleteActorResponse struct{}

type GetDirectorRequest struct {
	ID ID
}

type GetDirectorResponse struct {
	Director DirectorWithMovies
}

type GetDirectorsRequest struct {
	// Filters
	// TODO

	// Pagination
	Offset uint
	Limit  uint
}

type GetDirectorsResponse struct {
	Directors []Director
	Count     uint64
	Offset    uint
	Limit     uint
}

type SaveDirectorRequest struct {
	ID   ID
	Name string
}

type SaveDirectorResponse struct {
	Director Director
}

type DeleteDirectorRequest struct {
	ID ID
}

type DeleteDirectorResponse struct{}

type GetMovieGenreRequest struct {
	ID ID
}

type GetMovieGenreResponse struct {
	Genre MovieGenreWithMovies
}

// We don't paginate here because we don't expect this list to grow without
// bound.
type GetMovieGenresRequest struct {
	// Filters
	// TODO
}

type GetMovieGenresResponse struct {
	Genres []MovieGenre
}

type SaveMovieGenreRequest struct {
	ID   ID
	Name string
}

type SaveMovieGenreResponse struct {
	Genre MovieGenre
}

type DeleteMovieGenreRequest struct {
	ID ID
}

type DeleteMovieGenreResponse struct{}

type GetMovieRequest struct {
	ID ID
}

type GetMovieResponse struct {
	Movie MovieWithActors
}

type GetMoviesRequest struct {
	// Filters
	// TODO

	// Pagination
	Limit  uint
	Offset uint
}

type GetMoviesResponse struct {
	Movies []Movie
	Count  uint64
	Limit  uint
	Offset uint
}

type SaveMovieRequest struct {
	ID ID

	Director    ID
	NewDirector *struct {
		Name string
	}

	Genre    ID
	NewGenre *struct {
		Name string
	}

	Released int
	Title    string
}

type SaveMovieResponse struct {
	Movie Movie
}

type DeleteMovieRequest struct {
	ID ID
}

type DeleteMovieResponse struct{}

type RateMovieRequest struct {
	Movie  ID
	Rating Rating
	User   ID
}

type RateMovieResponse struct{}

type WatchMovieRequest struct {
	Date  time.Time
	Movie ID
	Users []ID
}

type WatchMovieResponse struct{}
