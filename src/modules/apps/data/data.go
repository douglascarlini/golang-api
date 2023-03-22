package data

import (
	"database/sql"
	"src/modules/apps/models"
	database "src/shared/database/postgres"

	_ "github.com/lib/pq"
)

const (
	sql_Count       = "SELECT COUNT(id) FROM apps"
	sql_DeleteByID  = "DELETE FROM apps WHERE id = $1"
	sql_UpdateByID  = "UPDATE apps SET label = $1 WHERE id = $2"
	sql_GetAll      = "SELECT id, name, label, secret FROM apps"
	sql_GetByID     = "SELECT id, name, label, secret FROM apps WHERE id = $1"
	sql_GetBySecret = "SELECT id, name, label, secret FROM apps WHERE secret = $1"
	sql_RenewSecret = "UPDATE apps SET secret = uuid_generate_v4() WHERE id = $1"
	sql_Insert      = "INSERT INTO apps (name, label) VALUES ($1, $2) RETURNING id"
)

type AppData struct {
	DB *sql.DB
}

func NewAppData() *AppData {
	return &AppData{DB: database.Connect()}
}

func Scan(rows *sql.Rows, item *models.App) error {
	return rows.Scan(&item.ID, &item.Name, &item.Label, &item.Secret)
}

func MakeDTO(item *models.App) models.AppDTO {

	return models.AppDTO{
		ID:    item.ID,
		Name:  item.Name,
		Label: item.Label,
	}

}
