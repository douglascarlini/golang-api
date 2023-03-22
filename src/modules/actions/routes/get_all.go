package routes

import (
	"src/modules/actions/data"
	"src/modules/actions/models"
	"src/shared/response"

	"github.com/gin-gonic/gin"
)

func GetAll() gin.HandlerFunc {

	return func(c *gin.Context) {

		var total int
		var err error
		repo := data.NewActionData()
		var result []models.ActionDTO

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
