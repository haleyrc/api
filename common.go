package api

import "github.com/pborman/uuid"

func NewID() ID {
	return ID(uuid.New())
}

type ID string

func (id ID) In(ids []ID) bool {
	for _, v := range ids {
		if v == id {
			return true
		}
	}
	return false
}

type Rating uint8
