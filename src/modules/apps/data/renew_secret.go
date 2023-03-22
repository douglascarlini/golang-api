package data

import "database/sql"

func (r *AppData) RenewSecret(id string) (int64, error) {

	var res sql.Result
	var result int64
	var err error

	if res, err = r.DB.Exec(sql_RenewSecret, id); err != nil {
		return 0, err
	}

	if result, err = res.RowsAffected(); err != nil {
		return 0, err
	}

	return result, nil

}
