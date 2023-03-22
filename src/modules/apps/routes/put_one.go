package routes

import (
	"src/modules/apps/data"
	"src/modules/apps/models"
	"src/modules/apps/utils"
	"src/shared/response"

	"github.com/gin-gonic/gin"
)

func PutOne() gin.HandlerFunc {
	return func(c *gin.Context) {

		var err error
		id := c.Param("id")
		var app models.App
		repo := data.NewAppData()

		if err = c.BindJSON(&app); err != nil {
			response.SendError(400, err.Error(), c)
			return
		}

		if utils.CheckID(id, repo, c) {

			if _, err = repo.SetOne(app, id); err != nil {
				response.SendError(500, err.Error(), c)
				return
			}

			response.SendOK(c)

		}

	}
}
