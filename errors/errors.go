package errors

import "fmt"

type ErrorKind string

const (
	KindBadRequest       ErrorKind = "bad-request"
	KindInternal         ErrorKind = "internal"
	KindResourceNotFound ErrorKind = "resource-not-found"
)

func Kind(err error) ErrorKind {
	switch err.(type) {
	case BadRequest:
		return KindBadRequest
	default:
		return KindInternal
	}
}

type BadRequest struct {
	Message string
}

func (err BadRequest) Error() string {
	return fmt.Sprint("bad request:", err.Message)
}

func (br BadRequest) Public() string {
	return br.Message
}
