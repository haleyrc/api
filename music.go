package api

type Artist struct {
	ID       ID
	Name     string
	NumActs  int
	NumSongs int

	// The number of songs they are featured on.
	FeaturedOn int
}

type ArtistWithActs struct {
	Artist
	AssociatedActs []Act
}

type Act struct {
	ID         ID
	Name       string
	NumArtists int
	NumAlbums  int
}

type ActWithArtistsAndAlbums struct {
	Act
	Artists []Artist
	Albums  []Album
}

type Album struct {
	Act        ID
	ID         ID
	MusicGenre ID
	Name       string
	Released   int
	NumSongs   int
}

type AlbumWithSongs struct {
	Album
	Songs []ID
}

type MusicGenre struct {
	ID        ID
	Name      string
	NumAlbums int
}

type MusicGenreWithAlbums struct {
	MusicGenre
	Albums []ID
}

type Song struct {
	Album     ID
	ID        ID
	Featuring []ID
	Name      string
}

type SongRating struct {
	Song   ID
	Rating Rating
	User   ID
}

type GetArtistRequest struct {
	ID ID
}

type GetArtistResponse struct {
	Artist ArtistWithActs
}

type GetArtistsRequest struct {
	// Filters
	// TODO

	// Pagination
	Limit  uint
	Offset uint
}

type GetArtistsResponse struct {
	Artists []Artist
	Count   uint64
	Limit   uint
	Offset  uint
}

type SaveArtistRequest struct {
	ID   ID
	Name string
}

type SaveArtistResponse struct {
	Artist Artist
}

type DeleteArtistRequest struct {
	ID ID
}

type DeleteArtistResponse struct{}

type AssociateArtistWithActRequest struct {
	Artist ID
	Act    ID
}

type AssociateArtistWithActResponse struct{}

type GetActRequest struct {
	ID ID
}

type GetActResponse struct {
	Act ActWithArtistsAndAlbums
}

type GetActsRequest struct {
	// Filters
	// TODO

	// Pagination
	Limit  uint
	Offset uint
}

type GetActsResponse struct {
	Acts   []Act
	Count  uint64
	Limit  uint
	Offset uint
}

type SaveActRequest struct {
	ID   ID
	Name string
}

type SaveActResponse struct {
	Act Act
}

type DeleteActRequest struct {
	ID ID
}

type DeleteActResponse struct{}

type GetAlbumRequest struct {
	ID ID
}

type GetAlbumResponse struct {
	Album AlbumWithSongs
}

type GetAlbumsRequest struct {
	// Filters
	// TODO

	// Pagination
	Limit  uint
	Offset uint
}

type GetAlbumsResponse struct {
	Albums []Album
	Count  uint64
	Limit  uint
	Offset uint
}

type SaveAlbumRequest struct {
	ID ID

	Act    ID
	NewAct *struct {
		Name string
	}

	Genre    ID
	NewGenre *struct {
		Name string
	}

	Name     string
	Released int
}

type SaveAlbumResponse struct {
}

type DeleteAlbumRequest struct {
}

type DeleteAlbumResponse struct {
}

type GetMusicGenreRequest struct {
}

type GetMusicGenreResponse struct {
}

type GetMusicGenresRequest struct {
}

type GetMusicGenresResponse struct {
}

type SaveMusicGenresRequest struct {
}

type SaveMusicGenresResponse struct {
}

type DeleteMusicGenresRequest struct {
}

type DeleteMusicGenresResponse struct {
}

type GetSongRequest struct {
}

type GetSongResponse struct {
}

type GetSongsRequest struct {
}

type GetSongsResponse struct {
}

type SaveSongRequest struct {
}

type SaveSongResponse struct {
}

type DeleteSongRequest struct {
}

type DeleteSongResponse struct {
}

type RateSongRequest struct {
}

type RateSongResponse struct {
}

type GetUserRequest struct {
}

type GetUserResponse struct {
}

type GetUsersRequest struct {
}

type GetUsersResponse struct {
}

type SaveUserRequest struct {
}

type SaveUserResponse struct {
}

type DeleteUserRequest struct {
}

type DeleteUserResponse struct {
}
