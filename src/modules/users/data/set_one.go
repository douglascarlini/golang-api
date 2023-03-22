package data

import (
	"database/sql"
	"src/modules/users/models"
	"src/shared/security"
)

func (r *UserData) SetOne(data models.User, id string) (int64, error) {

	var res sql.Result
	var hashed string
	var result int64
	var err error

	if hashed, err = security.PwdHash(data.Password); err != nil {
		return result, err
	}

	if res, err = r.DB.Exec(sql_UpdateByID, data.Name, data.Username, hashed, data.RoleID, id); err != nil {
		return 0, err
	}

	if result, err = res.RowsAffected(); err != nil {
		return 0, err
	}

	return result, nil

}
