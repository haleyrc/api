package api

type ExecutionContext struct {
	User User
}

type Service interface {
	CreateAuthor(ExecutionContext, CreateAuthorRequest) (*CreateAuthorResponse, error)
	GetAuthor(ExecutionContext, GetAuthorRequest) (*GetAuthorResponse, error)
	GetAuthors(ExecutionContext, GetAuthorsRequest) (*GetAuthorsResponse, error)
	EditAuthor(ExecutionContext, EditAuthorRequest) (*EditAuthorResponse, error)
	DeleteAuthor(ExecutionContext, DeleteAuthorRequest) (*DeleteAuthorResponse, error)

	CreateBook(ExecutionContext, CreateBookRequest) (*CreateBookResponse, error)
	GetBook(ExecutionContext, GetBookRequest) (*GetBookResponse, error)
	GetBooks(ExecutionContext, GetBooksRequest) (*GetBooksResponse, error)
	EditBook(ExecutionContext, EditBookRequest) (*EditBookResponse, error)
	DeleteBook(ExecutionContext, DeleteBookRequest) (*DeleteBookResponse, error)

	RateBook(ExecutionContext, RateBookRequest) (*RateBookResponse, error)
	StartBook(ExecutionContext, StartBookRequest) (*StartBookResponse, error)
	FinishBook(ExecutionContext, FinishBookRequest) (*FinishBookResponse, error)

	CreateBookGenre(ExecutionContext, CreateBookGenreRequest) (*CreateBookGenreResponse, error)
	GetBookGenre(ExecutionContext, GetBookGenreRequest) (*GetBookGenreResponse, error)
	GetBookGenres(ExecutionContext, GetBookGenresRequest) (*GetBookGenresResponse, error)
	EditBookGenre(ExecutionContext, EditBookGenreRequest) (*EditBookGenreResponse, error)
	DeleteBookGenre(ExecutionContext, DeleteBookGenreRequest) (*DeleteBookGenreResponse, error)

	CreateActor(ExecutionContext, CreateActorRequest) (*CreateActorResponse, error)
	GetActor(ExecutionContext, GetActorRequest) (*GetActorResponse, error)
	GetActors(ExecutionContext, GetActorsRequest) (*GetActorsResponse, error)
	EditActor(ExecutionContext, EditActorRequest) (*EditActorResponse, error)
	DeleteActor(ExecutionContext, DeleteActorRequest) (*DeleteActorResponse, error)

	CreateDirector(ExecutionContext, CreateDirectorRequest) (*CreateDirectorResponse, error)
	GetDirector(ExecutionContext, GetDirectorRequest) (*GetDirectorResponse, error)
	GetDirectors(ExecutionContext, GetDirectorsRequest) (*GetDirectorsResponse, error)
	EditDirector(ExecutionContext, EditDirectorRequest) (*EditDirectorResponse, error)
	DeleteDirector(ExecutionContext, DeleteDirectorRequest) (*DeleteDirectorResponse, error)

	CreateMovieGenre(ExecutionContext, CreateMovieGenreRequest) (*CreateMovieGenreResponse, error)
	GetMovieGenre(ExecutionContext, GetMovieGenreRequest) (*GetMovieGenreResponse, error)
	GetMovieGenres(ExecutionContext, GetMovieGenresRequest) (*GetMovieGenresResponse, error)
	EditMovieGenre(ExecutionContext, EditMovieGenreRequest) (*EditMovieGenreResponse, error)
	DeleteMovieGenre(ExecutionContext, DeleteMovieGenreRequest) (*DeleteMovieGenreResponse, error)

	CreateMovie(ExecutionContext, CreateMovieRequest) (*CreateMovieResponse, error)
	GetMovie(ExecutionContext, GetMovieRequest) (*GetMovieResponse, error)
	GetMovies(ExecutionContext, GetMoviesRequest) (*GetMoviesResponse, error)
	EditMovie(ExecutionContext, EditMovieRequest) (*EditMovieResponse, error)
	DeleteMovie(ExecutionContext, DeleteMovieRequest) (*DeleteMovieResponse, error)

	RateMovie(ExecutionContext, RateMovieRequest) (*RateMovieResponse, error)
	WatchMovie(ExecutionContext, WatchMovieRequest) (*WatchMovieResponse, error)

	CreateArtist(ExecutionContext, CreateArtistRequest) (*CreateArtistResponse, error)
	GetArtist(ExecutionContext, GetArtistRequest) (*GetArtistResponse, error)
	GetArtists(ExecutionContext, GetArtistsRequest) (*GetArtistsResponse, error)
	EditArtist(ExecutionContext, EditArtistRequest) (*EditArtistResponse, error)
	DeleteArtist(ExecutionContext, DeleteArtistRequest) (*DeleteArtistResponse, error)

	AssociateArtistWithAct(ExecutionContext, AssociateArtistWithActRequest) (*AssociateArtistWithActResponse, error)

	CreateAct(ExecutionContext, CreateActRequest) (*CreateActResponse, error)
	GetAct(ExecutionContext, GetActRequest) (*GetActResponse, error)
	GetActs(ExecutionContext, GetActsRequest) (*GetActsResponse, error)
	EditAct(ExecutionContext, EditActRequest) (*EditActResponse, error)
	DeleteAct(ExecutionContext, DeleteActRequest) (*DeleteActResponse, error)

	CreateAlbum(ExecutionContext, CreateAlbumRequest) (*CreateAlbumResponse, error)
	GetAlbum(ExecutionContext, GetAlbumRequest) (*GetAlbumResponse, error)
	GetAlbums(ExecutionContext, GetAlbumsRequest) (*GetAlbumsResponse, error)
	EditAlbum(ExecutionContext, EditAlbumRequest) (*EditAlbumResponse, error)
	DeleteAlbum(ExecutionContext, DeleteAlbumRequest) (*DeleteAlbumResponse, error)

	CreateMusicGenre(ExecutionContext, CreateMusicGenreRequest) (*CreateMusicGenreResponse, error)
	GetMusicGenre(ExecutionContext, GetMusicGenreRequest) (*GetMusicGenreResponse, error)
	GetMusicGenres(ExecutionContext, GetMusicGenresRequest) (*GetMusicGenresResponse, error)
	EditMusicGenres(ExecutionContext, EditMusicGenresRequest) (*EditMusicGenresResponse, error)
	DeleteMusicGenres(ExecutionContext, DeleteMusicGenresRequest) (*DeleteMusicGenresResponse, error)

	CreateSong(ExecutionContext, CreateSongRequest) (*CreateSongResponse, error)
	GetSong(ExecutionContext, GetSongRequest) (*GetSongResponse, error)
	GetSongs(ExecutionContext, GetSongsRequest) (*GetSongsResponse, error)
	EditSong(ExecutionContext, EditSongRequest) (*EditSongResponse, error)
	DeleteSong(ExecutionContext, DeleteSongRequest) (*DeleteSongResponse, error)

	RateSong(ExecutionContext, RateSongRequest) (*RateSongResponse, error)

	CreateUser(ExecutionContext, CreateUserRequest) (*CreateUserResponse, error)
	GetUser(ExecutionContext, GetUserRequest) (*GetUserResponse, error)
	GetUsers(ExecutionContext, GetUsersRequest) (*GetUsersResponse, error)
	EditUser(ExecutionContext, EditUserRequest) (*EditUserResponse, error)
	DeleteUser(ExecutionContext, DeleteUserRequest) (*DeleteUserResponse, error)
}
