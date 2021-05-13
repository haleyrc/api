package api

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
	Director   ID
	MovieGenre ID
	ID         ID
	Released   Year
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
	Rating ThumbRating
	User   ID
}

type MovieViewing struct {
	Date  Time
	ID    ID
	Movie ID
	Users []ID
}
