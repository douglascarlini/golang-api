package data

import "database/sql"

func (r *DomainData) Count() (int, error) {

	var rows *sql.Rows
	var err error

	if rows, err = r.DB.Query(sql_Count); err != nil {
		return 0, err
	}

	defer rows.Close()

	total := 0
	rows.Next()
	rows.Scan(&total)

	return total, nil
}
