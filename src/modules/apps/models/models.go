package models

import (
	"github.com/google/uuid"
)

type App struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Label  string    `json:"label"`
	Secret uuid.UUID `json:"secret"`
}

type AppDTO struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Label string    `json:"label"`
}
