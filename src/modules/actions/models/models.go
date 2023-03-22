package models

import (
	"github.com/google/uuid"
)

type Action struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Label string    `json:"label"`
}

type ActionDTO struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Label string    `json:"label"`
}
