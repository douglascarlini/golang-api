package data

import (
	"database/sql"
	"src/modules/roles/models"
)

func (r *RoleData) GetOne(id string) (models.RoleDTO, error) {

	var dto models.RoleDTO
	var item models.Role
	var rows *sql.Rows
	var err error

	if rows, err = r.DB.Query(sql_GetByID, id); err != nil {
		return dto, err
	}

	defer rows.Close()

	if err = rows.Err(); err != nil {
		return dto, err
	}

	if !rows.Next() {
		return dto, nil
	}

	if err = Scan(rows, &item); err != nil {
		return dto, err
	}

	dto = MakeDTO(&item)
	return dto, nil

}
