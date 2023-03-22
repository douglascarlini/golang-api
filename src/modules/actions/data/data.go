package data

import (
	"database/sql"
	"src/modules/actions/models"
	db1 "src/shared/database/postgres"

	_ "github.com/lib/pq"
)

const (
	sql_Count      = "SELECT COUNT(id) FROM actions"
	sql_DeleteByID = "DELETE FROM actions WHERE id = $1"
	sql_GetAll     = "SELECT id, name, label FROM actions"
	sql_UpdateByID = "UPDATE actions SET label = $1 WHERE id = $2"
	sql_GetByID    = "SELECT id, name, label FROM actions WHERE id = $1"
	sql_Insert     = "INSERT INTO actions (name, label) VALUES ($1, $2) RETURNING id"
)

type ActionData struct {
	DB *sql.DB
}

func NewActionData() *ActionData {
	return &ActionData{DB: db1.Connect()}
}

func Scan(rows *sql.Rows, item *models.Action) error {
	return rows.Scan(&item.ID, &item.Name, &item.Label)
}

func MakeDTO(item *models.Action) models.ActionDTO {

	return models.ActionDTO{
		ID:    item.ID,
		Name:  item.Name,
		Label: item.Label,
	}

}
