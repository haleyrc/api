package api

import "github.com/pborman/uuid"

type ID string

func NewID() ID {
	return ID(uuid.New())
}

type Rating uint8
