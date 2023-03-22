package data

import (
	"database/sql"
	"src/modules/users/models"
)

func (r *UserData) GetOne(id string) (models.UserDTO, error) {

	var dto models.UserDTO
	var rows *sql.Rows
	var err error

	if rows, err = r.DB.Query(sql_GetByID, id); err != nil {
		return dto, err
	}

	defer rows.Close()

	if !rows.Next() {
		return dto, nil
	}

	var item models.User

	if err = Scan(rows, &item); err != nil {
		return dto, err
	}

	if err = rows.Err(); err != nil {
		return dto, err
	}

	user := MakeDTO(&item)

	return user, nil

}
