package routes

import (
	"src/modules/modules/data"
	"src/modules/modules/models"
	"src/modules/modules/utils"
	"src/shared/response"

	"github.com/gin-gonic/gin"
)

func PutOne() gin.HandlerFunc {

	return func(c *gin.Context) {

		var err error
		id := c.Param("id")
		var app models.Module
		repo := data.NewModuleData()

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
