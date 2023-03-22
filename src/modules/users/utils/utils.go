package utils

import (
	"src/modules/users/data"
	"src/modules/users/models"
	"src/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CheckUserID(id string, repo *data.UserData, c *gin.Context) bool {

	var result models.UserDTO
	var err error

	if result, err = repo.GetOne(id); err != nil {
		response.SendError(500, err.Error(), c)
		return false
	}

	if result.ID == uuid.Nil {
		response.SendError(404, "not found", c)
		return false
	}

	return true

}

func CheckUserData(post *models.User, c *gin.Context) bool {

	if post.Username == "" {
		response.SendError(400, "invalid username", c)
		return false
	}

	if post.Password == "" {
		response.SendError(400, "invalid password", c)
		return false
	}

	if post.RoleID == uuid.Nil {
		response.SendError(400, "invalid role_id", c)
		return false
	}

	return true

}
