package pages

import "github.com/haleyrc/api"

const (
	FlashError   = "error"
	FlashWarning = "warning"
	FlashSuccess = "success"
)

type Flash struct {
	Messages []FlashMessage
}

func (f *Flash) Error(msg string) {
	f.Messages = append(f.Messages, FlashMessage{
		Message: msg,
		Type:    FlashError,
	})
}

type FlashMessage struct {
	Message string
	Type    string
}

type Page struct {
	Flash []Flash
	User  *api.User
}

type Login struct {
	Flash
}

type Error struct {
	Page
	Error string
}
