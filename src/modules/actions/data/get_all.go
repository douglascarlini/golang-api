package data

import (
	"database/sql"
	"src/modules/actions/models"
)

func (r *ActionData) GetAll() ([]models.ActionDTO, error) {

	var result []models.ActionDTO
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

		var item models.Action

		if err = Scan(rows, &item); err != nil {
			return nil, err
		}

		dto := MakeDTO(&item)
		result = append(result, dto)
	}

	return result, nil

}
