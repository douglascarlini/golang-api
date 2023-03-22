package data

import (
	"database/sql"
	"src/modules/users/models"
)

func (r *UserData) GetAll() ([]models.UserDTO, error) {

	var result []models.UserDTO
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

		var item models.User

		if err = Scan(rows, &item); err != nil {
			return nil, err
		}

		dto := MakeDTO(&item)
		result = append(result, dto)
	}

	return result, nil

}
