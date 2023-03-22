package models

import (
	"src/modules/roles/models"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	RoleID   uuid.UUID `json:"role_id"`
}

type UserDTO struct {
	ID       uuid.UUID      `json:"id"`
	Name     string         `json:"name"`
	Username string         `json:"username"`
	Role     models.RoleDTO `json:"role"`
}
