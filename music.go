package api

type Artist struct {
	ID      ID
	Name    string
	NumActs int
	// TODO: Should this be all their songs from associated acts or only how
	// many they're featured on?
	NumSongs int
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

type ActWithArtists struct {
	Act
	Artists []Artist
}

type ActWithAlbums struct {
	Act
	Albums []ID
}

type Album struct {
	Act        ID
	ID         ID
	MusicGenre ID
	Name       string
	Released   Year
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
	Rating ThumbRating
	User   ID
}
