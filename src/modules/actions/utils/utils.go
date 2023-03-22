package utils

import (
	"src/modules/actions/data"
	"src/modules/actions/models"
	"src/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CheckID(id string, repo *data.ActionData, c *gin.Context) bool {

	var result models.ActionDTO
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
