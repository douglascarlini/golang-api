package data

import (
	"database/sql"
	"src/modules/actions/models"
)

func (r *ActionData) GetOne(id string) (models.ActionDTO, error) {

	var dto models.ActionDTO
	var item models.Action
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
