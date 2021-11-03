package api

import "context"

type ExecutionContext struct {
	Ctx  context.Context
	User User
}

type Service interface {
	GetActor(ExecutionContext, GetActorRequest) (*GetActorResponse, error)
	GetActors(ExecutionContext, GetActorsRequest) (*GetActorsResponse, error)
	SaveActor(ExecutionContext, SaveActorRequest) (*SaveActorResponse, error)
	DeleteActor(ExecutionContext, DeleteActorRequest) (*DeleteActorResponse, error)

	GetDirector(ExecutionContext, GetDirectorRequest) (*GetDirectorResponse, error)
	GetDirectors(ExecutionContext, GetDirectorsRequest) (*GetDirectorsResponse, error)
	SaveDirector(ExecutionContext, SaveDirectorRequest) (*SaveDirectorResponse, error)
	DeleteDirector(ExecutionContext, DeleteDirectorRequest) (*DeleteDirectorResponse, error)

	GetMovieGenre(ExecutionContext, GetMovieGenreRequest) (*GetMovieGenreResponse, error)
	GetMovieGenres(ExecutionContext, GetMovieGenresRequest) (*GetMovieGenresResponse, error)
	SaveMovieGenre(ExecutionContext, SaveMovieGenreRequest) (*SaveMovieGenreResponse, error)
	DeleteMovieGenre(ExecutionContext, DeleteMovieGenreRequest) (*DeleteMovieGenreResponse, error)

	GetMovie(ExecutionContext, GetMovieRequest) (*GetMovieResponse, error)
	GetMovies(ExecutionContext, GetMoviesRequest) (*GetMoviesResponse, error)
	SaveMovie(ExecutionContext, SaveMovieRequest) (*SaveMovieResponse, error)
	DeleteMovie(ExecutionContext, DeleteMovieRequest) (*DeleteMovieResponse, error)

	RateMovie(ExecutionContext, RateMovieRequest) (*RateMovieResponse, error)
	WatchMovie(ExecutionContext, WatchMovieRequest) (*WatchMovieResponse, error)

	GetArtist(ExecutionContext, GetArtistRequest) (*GetArtistResponse, error)
	GetArtists(ExecutionContext, GetArtistsRequest) (*GetArtistsResponse, error)
	SaveArtist(ExecutionContext, SaveArtistRequest) (*SaveArtistResponse, error)
	DeleteArtist(ExecutionContext, DeleteArtistRequest) (*DeleteArtistResponse, error)

	AssociateArtistWithAct(ExecutionContext, AssociateArtistWithActRequest) (*AssociateArtistWithActResponse, error)

	GetAct(ExecutionContext, GetActRequest) (*GetActResponse, error)
	GetActs(ExecutionContext, GetActsRequest) (*GetActsResponse, error)
	SaveAct(ExecutionContext, SaveActRequest) (*SaveActResponse, error)
	DeleteAct(ExecutionContext, DeleteActRequest) (*DeleteActResponse, error)

	GetAlbum(ExecutionContext, GetAlbumRequest) (*GetAlbumResponse, error)
	GetAlbums(ExecutionContext, GetAlbumsRequest) (*GetAlbumsResponse, error)
	SaveAlbum(ExecutionContext, SaveAlbumRequest) (*SaveAlbumResponse, error)
	DeleteAlbum(ExecutionContext, DeleteAlbumRequest) (*DeleteAlbumResponse, error)

	GetMusicGenre(ExecutionContext, GetMusicGenreRequest) (*GetMusicGenreResponse, error)
	GetMusicGenres(ExecutionContext, GetMusicGenresRequest) (*GetMusicGenresResponse, error)
	SaveMusicGenres(ExecutionContext, SaveMusicGenresRequest) (*SaveMusicGenresResponse, error)
	DeleteMusicGenres(ExecutionContext, DeleteMusicGenresRequest) (*DeleteMusicGenresResponse, error)

	GetSong(ExecutionContext, GetSongRequest) (*GetSongResponse, error)
	GetSongs(ExecutionContext, GetSongsRequest) (*GetSongsResponse, error)
	SaveSong(ExecutionContext, SaveSongRequest) (*SaveSongResponse, error)
	DeleteSong(ExecutionContext, DeleteSongRequest) (*DeleteSongResponse, error)

	RateSong(ExecutionContext, RateSongRequest) (*RateSongResponse, error)

	GetUser(ExecutionContext, GetUserRequest) (*GetUserResponse, error)
	GetUsers(ExecutionContext, GetUsersRequest) (*GetUsersResponse, error)
	SaveUser(ExecutionContext, SaveUserRequest) (*SaveUserResponse, error)
	DeleteUser(ExecutionContext, DeleteUserRequest) (*DeleteUserResponse, error)
}
