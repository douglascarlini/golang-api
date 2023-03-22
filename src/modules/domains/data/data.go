package data

import (
	"database/sql"
	"src/modules/domains/models"
	db1 "src/shared/database/postgres"

	_ "github.com/lib/pq"
)

const (
	sql_Count      = "SELECT COUNT(id) FROM domains"
	sql_DeleteByID = "DELETE FROM domains WHERE id = $1"
	sql_GetAll     = "SELECT id, name, label FROM domains"
	sql_UpdateByID = "UPDATE domains SET label = $1 WHERE id = $2"
	sql_GetByID    = "SELECT id, name, label FROM domains WHERE id = $1"
	sql_Insert     = "INSERT INTO domains (name, label) VALUES ($1, $2) RETURNING id"
)

type DomainData struct {
	DB *sql.DB
}

func NewDomainData() *DomainData {
	return &DomainData{DB: db1.Connect()}
}

func Scan(rows *sql.Rows, item *models.Domain) error {
	return rows.Scan(&item.ID, &item.Name, &item.Label)
}

func MakeDTO(item *models.Domain) models.DomainDTO {

	return models.DomainDTO{
		ID:    item.ID,
		Name:  item.Name,
		Label: item.Label,
	}

}
