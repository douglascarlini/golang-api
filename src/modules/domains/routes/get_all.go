package routes

import (
	"src/modules/domains/data"
	"src/modules/domains/models"
	"src/shared/response"

	"github.com/gin-gonic/gin"
)

func GetAll() gin.HandlerFunc {

	return func(c *gin.Context) {

		var total int
		var err error
		repo := data.NewDomainData()
		var result []models.DomainDTO

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
