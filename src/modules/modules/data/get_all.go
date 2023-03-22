package data

import (
	"database/sql"
	"src/modules/modules/models"
)

func (r *ModuleData) GetAll() ([]models.ModuleDTO, error) {

	var result []models.ModuleDTO
	var rows *sql.Rows
	var err error

	if rows, err = r.DB.Query(sql_GetAll); err != nil {
		return nil, err
	}

	defer rows.Close()

	if err = rows.Err(); err != nil {
		return nil, err
	}

	for rows.Next() {

		var item models.Module

		if err = Scan(rows, &item); err != nil {
			return nil, err
		}

		dto := MakeDTO(&item)
		result = append(result, dto)
	}

	return result, nil

}
