package entity

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewId() ID {
	return ID(uuid.New())
}

func ParseId(value string) (ID, error) {
	id, err := uuid.Parse(value)

	return ID(id), err
}
