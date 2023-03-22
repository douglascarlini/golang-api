package routes

import (
	"src/modules/users/data"
	"src/modules/users/models"
	"src/modules/users/utils"
	"src/shared/response"

	"github.com/gin-gonic/gin"
)

func PutOne() gin.HandlerFunc {
	return func(c *gin.Context) {

		var err error
		id := c.Param("id")
		var user models.User
		repo := data.NewUserData()

		if err = c.BindJSON(&user); err != nil {
			response.SendError(400, err.Error(), c)
			return
		}

		if utils.CheckUserID(id, repo, c) {

			if _, err = repo.SetOne(user, id); err != nil {
				response.SendError(500, err.Error(), c)
				return
			}

			response.SendOK(c)

		}

	}
}
