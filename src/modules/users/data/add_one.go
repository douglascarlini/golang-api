package data

import (
	"src/modules/users/models"
	"src/shared/security"

	"github.com/google/uuid"
)

func (r *UserData) AddOne(data models.User) (uuid.UUID, error) {

	var result uuid.UUID
	var hashed string
	var err error

	if hashed, err = security.PwdHash(data.Password); err != nil {
		return result, err
	}

	if err = r.DB.QueryRow(sql_Insert, data.Name, data.Username, hashed, data.RoleID).Scan(&result); err != nil {
		return result, err
	}

	return result, nil

}
