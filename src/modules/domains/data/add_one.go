package data

import (
	"src/modules/domains/models"

	"github.com/google/uuid"
)

func (r *DomainData) AddOne(data models.Domain) (uuid.UUID, error) {

	var result uuid.UUID
	var err error

	if err = r.DB.QueryRow(sql_Insert, data.Name, data.Label).Scan(&result); err != nil {
		return result, err
	}

	return result, nil

}
