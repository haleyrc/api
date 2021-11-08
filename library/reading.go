package library

import (
	"time"

	"github.com/haleyrc/api"
)

type Reading struct {
	ID       api.ID
	Book     api.ID
	User     api.ID
	Started  time.Time
	Finished time.Time
}
