package routes

import (
	"src/modules/domains/data"
	"src/modules/domains/models"
	"src/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetOne() gin.HandlerFunc {

	return func(c *gin.Context) {

		var err error
		id := c.Param("id")
		var result models.DomainDTO
		repo := data.NewDomainData()

		if result, err = repo.GetOne(id); err != nil {
			response.SendError(500, err.Error(), c)
			return
		}

		if result.ID == uuid.Nil {
			response.SendError(404, "not found", c)
			return
		}

		response.SendData(&result, c)

	}

}
