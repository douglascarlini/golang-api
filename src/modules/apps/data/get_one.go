package data

import (
	"database/sql"
	"src/modules/apps/models"
)

func (r *AppData) GetOne(id string) (models.AppDTO, error) {

	var result models.AppDTO
	var item models.App
	var rows *sql.Rows
	var err error

	if rows, err = r.DB.Query(sql_GetByID, id); err != nil {
		return result, err
	}

	defer rows.Close()

	if !rows.Next() {
		return result, nil
	}

	if err = Scan(rows, &item); err != nil {
		return result, err
	}

	if err = rows.Err(); err != nil {
		return result, err
	}

	dto := MakeDTO(&item)
	return dto, nil

}
