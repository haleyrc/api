package library

import "github.com/haleyrc/api"

type AuthorsFilter struct {
	IDs []api.ID
}

type Author struct {
	ID   api.ID
	Name string
}
