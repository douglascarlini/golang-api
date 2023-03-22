package routes

import (
	"src/modules/apps/data"
	"src/modules/apps/utils"
	"src/shared/response"

	"github.com/gin-gonic/gin"
)

func DelOne() gin.HandlerFunc {
	return func(c *gin.Context) {

		var err error
		id := c.Param("id")
		repo := data.NewAppData()

		if utils.CheckID(id, repo, c) {

			if _, err = repo.DelOne(id); err != nil {
				response.SendError(500, err.Error(), c)
				return
			}

			response.SendOK(c)

		}

	}
}
