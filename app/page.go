package app

import "github.com/haleyrc/api"

const (
	Error   = "error"
	Warning = "warning"
	Success = "success"
)

type flash struct {
	Message string
	Type    string
}

func newPage(s scope) page {
	return page{User: s.User}
}

type page struct {
	Flash []flash
	User  *api.User
}

func (p *page) Error(msg string) {
	p.Flash = append(p.Flash, flash{
		Message: msg,
		Type:    Error,
	})
}
