package models

import (
	"github.com/google/uuid"
)

type Role struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Label string    `json:"label"`
}

type RoleDTO struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Label string    `json:"label"`
}
