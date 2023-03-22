package data

import (
	"database/sql"
	"src/modules/roles/models"
	db1 "src/shared/database/postgres"

	_ "github.com/lib/pq"
)

const (
	sql_Count      = "SELECT COUNT(id) FROM roles"
	sql_DeleteByID = "DELETE FROM roles WHERE id = $1"
	sql_GetAll     = "SELECT id, name, label FROM roles"
	sql_UpdateByID = "UPDATE roles SET label = $1 WHERE id = $2"
	sql_GetByID    = "SELECT id, name, label FROM roles WHERE id = $1"
	sql_Insert     = "INSERT INTO roles (name, label) VALUES ($1, $2) RETURNING id"
)

type RoleData struct {
	DB *sql.DB
}

func NewRoleData() *RoleData {
	return &RoleData{DB: db1.Connect()}
}

func Scan(rows *sql.Rows, item *models.Role) error {
	return rows.Scan(&item.ID, &item.Name, &item.Label)
}

func MakeDTO(item *models.Role) models.RoleDTO {

	return models.RoleDTO{
		ID:    item.ID,
		Name:  item.Name,
		Label: item.Label,
	}

}
