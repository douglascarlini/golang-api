package data

import (
	"database/sql"
	"src/modules/modules/models"
	db1 "src/shared/database/postgres"

	_ "github.com/lib/pq"
)

const (
	sql_Count      = "SELECT COUNT(id) FROM modules"
	sql_DeleteByID = "DELETE FROM modules WHERE id = $1"
	sql_GetAll     = "SELECT id, name, label FROM modules"
	sql_UpdateByID = "UPDATE modules SET label = $1 WHERE id = $2"
	sql_GetByID    = "SELECT id, name, label FROM modules WHERE id = $1"
	sql_Insert     = "INSERT INTO modules (name, label) VALUES ($1, $2) RETURNING id"
)

type ModuleData struct {
	DB *sql.DB
}

func NewModuleData() *ModuleData {
	return &ModuleData{DB: db1.Connect()}
}

func Scan(rows *sql.Rows, item *models.Module) error {
	return rows.Scan(&item.ID, &item.Name, &item.Label)
}

func MakeDTO(item *models.Module) models.ModuleDTO {

	return models.ModuleDTO{
		ID:    item.ID,
		Name:  item.Name,
		Label: item.Label,
	}

}
