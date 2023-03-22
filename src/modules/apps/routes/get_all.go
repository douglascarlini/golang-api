package routes

import (
	"src/modules/apps/data"
	"src/modules/apps/models"
	"src/shared/response"

	"github.com/gin-gonic/gin"
)

func GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		var total int
		var err error
		var result []models.AppDTO
		repo := data.NewAppData()

		if result, err = repo.GetAll(); err != nil {
			response.SendError(500, err.Error(), c)
			return
		}

		if total, err = repo.Count(); err != nil {
			response.SendError(500, err.Error(), c)
			return
		}

		response.SendList(total, &result, c)

	}
}
