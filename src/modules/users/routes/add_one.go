package routes

import (
	"src/modules/users/data"
	"src/modules/users/models"
	"src/shared/response"
	"src/shared/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddOne() gin.HandlerFunc {
	return func(c *gin.Context) {

		var err error
		var post models.User
		var find models.User
		var result uuid.UUID
		repo := data.NewUserData()

		if _, err = utils.CheckAppKey(c); err != nil {
			response.SendError(400, err.Error(), c)
			return
		}

		if err := c.BindJSON(&post); err != nil {
			response.SendError(400, err.Error(), c)
			return
		}

		if find, _ = repo.GetByUsername(post.Username); find.Password != "" {
			response.SendError(409, "duplicated", c)
			return
		}

		if result, err = repo.AddOne(post); err != nil {
			response.SendError(500, err.Error(), c)
			return
		}

		response.SendID(result.String(), c)

	}
}
