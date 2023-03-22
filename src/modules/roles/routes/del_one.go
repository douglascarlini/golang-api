package routes

import (
	"src/modules/roles/data"
	"src/modules/roles/utils"
	"src/shared/response"

	"github.com/gin-gonic/gin"
)

func DelOne() gin.HandlerFunc {

	return func(c *gin.Context) {

		var err error
		id := c.Param("id")
		repo := data.NewRoleData()

		if utils.CheckID(id, repo, c) {

			if _, err = repo.DelOne(id); err != nil {
				response.SendError(500, err.Error(), c)
				return
			}

			response.SendOK(c)

		}

	}

}
