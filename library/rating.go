package library

import "github.com/haleyrc/api"

type Rating struct {
	ID     api.ID
	Book   api.ID
	User   api.ID
	Rating api.Rating
}
