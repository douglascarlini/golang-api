package data

import (
	"database/sql"
	"src/modules/apps/models"
)

func (r *AppData) SetOne(data models.App, id string) (int64, error) {

	var res sql.Result
	var result int64
	var err error

	if res, err = r.DB.Exec(sql_UpdateByID, data.Label, id); err != nil {
		return 0, err
	}

	if result, err = res.RowsAffected(); err != nil {
		return 0, err
	}

	return result, nil

}
