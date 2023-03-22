package data

import (
	"src/modules/roles/models"

	"github.com/google/uuid"
)

func (r *RoleData) AddOne(data models.Role) (uuid.UUID, error) {

	var result uuid.UUID
	var err error

	if err = r.DB.QueryRow(sql_Insert, data.Name, data.Label).Scan(&result); err != nil {
		return result, err
	}

	return result, nil

}
