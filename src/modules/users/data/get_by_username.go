package data

import (
	"database/sql"
	"fmt"
	"src/modules/users/models"
)

func (r *UserData) GetByUsername(username string) (models.User, error) {

	var item models.User
	var rows *sql.Rows
	var err error

	if rows, err = r.DB.Query(sql_GetByUsername, username); err != nil {
		return item, err
	}

	defer rows.Close()

	if !rows.Next() {
		return item, fmt.Errorf("not found")
	}

	if err = rows.Err(); err != nil {
		return item, err
	}

	if err = Scan(rows, &item); err != nil {
		return item, err
	}

	return item, nil

}
