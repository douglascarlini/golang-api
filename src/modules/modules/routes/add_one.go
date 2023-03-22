package routes

import (
	"src/modules/modules/data"
	"src/modules/modules/models"
	"src/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddOne() gin.HandlerFunc {

	return func(c *gin.Context) {

		var err error
		var post models.Module
		var result uuid.UUID
		repo := data.NewModuleData()

		if err := c.BindJSON(&post); err != nil {
			response.SendError(400, err.Error(), c)
			return
		}

		if result, err = repo.AddOne(post); err != nil {
			response.SendError(500, err.Error(), c)
			return
		}

		response.SendID(result.String(), c)

	}

}
