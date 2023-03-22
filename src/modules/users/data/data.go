package data

import (
	"database/sql"
	roleData "src/modules/roles/data"
	roleModels "src/modules/roles/models"
	"src/modules/users/models"
	database "src/shared/database/postgres"

	_ "github.com/lib/pq"
)

const (
	sql_Count         = "SELECT COUNT(id) FROM users"
	sql_DeleteByID    = "DELETE FROM users WHERE id = $1"
	sql_GetAll        = "SELECT id, name, username, password, role_id FROM users"
	sql_GetByID       = "SELECT id, name, username, password, role_id FROM users WHERE id = $1"
	sql_GetByUsername = "SELECT id, name, username, password, role_id FROM users WHERE username = $1"
	sql_UpdateByID    = "UPDATE users SET name = $1, username = $2, password = $3, role_id = $4 WHERE id = $5"
	sql_Insert        = "INSERT INTO users (name, username, password, role_id) VALUES ($1, $2, $3, $4) RETURNING id"
)

type UserData struct {
	DB *sql.DB
}

func NewUserData() *UserData {
	return &UserData{DB: database.Connect()}
}

func Scan(rows *sql.Rows, item *models.User) error {
	if err := rows.Scan(
		&item.ID,
		&item.Name,
		&item.Username,
		&item.Password,
		&item.RoleID,
	); err != nil {
		return err
	}
	return nil
}

func MakeDTO(item *models.User) models.UserDTO {

	role, _ := roleData.NewRoleData().GetOne(item.RoleID.String())

	return models.UserDTO{
		ID:       item.ID,
		Name:     item.Name,
		Username: item.Username,
		Role:     roleModels.RoleDTO(role),
	}

}
