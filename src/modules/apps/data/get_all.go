package data

import (
	"database/sql"
	"src/modules/apps/models"
)

func (r *AppData) GetAll() ([]models.AppDTO, error) {

	var result []models.AppDTO
	var rows *sql.Rows
	var err error

	if rows, err = r.DB.Query(sql_GetAll); err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var item models.App

		if err = Scan(rows, &item); err != nil {
			return nil, err
		}

		dto := MakeDTO(&item)
		result = append(result, dto)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil

}
