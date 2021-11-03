package api

import "context"

type Service interface {
	GetActor(context.Context, GetActorRequest) (*GetActorResponse, error)
	GetActors(context.Context, GetActorsRequest) (*GetActorsResponse, error)
	SaveActor(context.Context, SaveActorRequest) (*SaveActorResponse, error)
	DeleteActor(context.Context, DeleteActorRequest) (*DeleteActorResponse, error)

	GetDirector(context.Context, GetDirectorRequest) (*GetDirectorResponse, error)
	GetDirectors(context.Context, GetDirectorsRequest) (*GetDirectorsResponse, error)
	SaveDirector(context.Context, SaveDirectorRequest) (*SaveDirectorResponse, error)
	DeleteDirector(context.Context, DeleteDirectorRequest) (*DeleteDirectorResponse, error)

	GetMovieGenre(context.Context, GetMovieGenreRequest) (*GetMovieGenreResponse, error)
	GetMovieGenres(context.Context, GetMovieGenresRequest) (*GetMovieGenresResponse, error)
	SaveMovieGenre(context.Context, SaveMovieGenreRequest) (*SaveMovieGenreResponse, error)
	DeleteMovieGenre(context.Context, DeleteMovieGenreRequest) (*DeleteMovieGenreResponse, error)

	GetMovie(context.Context, GetMovieRequest) (*GetMovieResponse, error)
	GetMovies(context.Context, GetMoviesRequest) (*GetMoviesResponse, error)
	SaveMovie(context.Context, SaveMovieRequest) (*SaveMovieResponse, error)
	DeleteMovie(context.Context, DeleteMovieRequest) (*DeleteMovieResponse, error)

	RateMovie(context.Context, RateMovieRequest) (*RateMovieResponse, error)
	WatchMovie(context.Context, WatchMovieRequest) (*WatchMovieResponse, error)

	GetArtist(context.Context, GetArtistRequest) (*GetArtistResponse, error)
	GetArtists(context.Context, GetArtistsRequest) (*GetArtistsResponse, error)
	SaveArtist(context.Context, SaveArtistRequest) (*SaveArtistResponse, error)
	DeleteArtist(context.Context, DeleteArtistRequest) (*DeleteArtistResponse, error)

	AssociateArtistWithAct(context.Context, AssociateArtistWithActRequest) (*AssociateArtistWithActResponse, error)

	GetAct(context.Context, GetActRequest) (*GetActResponse, error)
	GetActs(context.Context, GetActsRequest) (*GetActsResponse, error)
	SaveAct(context.Context, SaveActRequest) (*SaveActResponse, error)
	DeleteAct(context.Context, DeleteActRequest) (*DeleteActResponse, error)

	GetAlbum(context.Context, GetAlbumRequest) (*GetAlbumResponse, error)
	GetAlbums(context.Context, GetAlbumsRequest) (*GetAlbumsResponse, error)
	SaveAlbum(context.Context, SaveAlbumRequest) (*SaveAlbumResponse, error)
	DeleteAlbum(context.Context, DeleteAlbumRequest) (*DeleteAlbumResponse, error)

	GetMusicGenre(context.Context, GetMusicGenreRequest) (*GetMusicGenreResponse, error)
	GetMusicGenres(context.Context, GetMusicGenresRequest) (*GetMusicGenresResponse, error)
	SaveMusicGenres(context.Context, SaveMusicGenresRequest) (*SaveMusicGenresResponse, error)
	DeleteMusicGenres(context.Context, DeleteMusicGenresRequest) (*DeleteMusicGenresResponse, error)

	GetSong(context.Context, GetSongRequest) (*GetSongResponse, error)
	GetSongs(context.Context, GetSongsRequest) (*GetSongsResponse, error)
	SaveSong(context.Context, SaveSongRequest) (*SaveSongResponse, error)
	DeleteSong(context.Context, DeleteSongRequest) (*DeleteSongResponse, error)

	RateSong(context.Context, RateSongRequest) (*RateSongResponse, error)

	GetUser(context.Context, GetUserRequest) (*GetUserResponse, error)
	GetUsers(context.Context, GetUsersRequest) (*GetUsersResponse, error)
	SaveUser(context.Context, SaveUserRequest) (*SaveUserResponse, error)
	DeleteUser(context.Context, DeleteUserRequest) (*DeleteUserResponse, error)
}
