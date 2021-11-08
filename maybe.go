package api

type MaybeID struct {
	Valid bool
	Value ID
}

type MaybeInt struct {
	Valid bool
	Value int
}

type MaybeString struct {
	Valid bool
	Value string
}
