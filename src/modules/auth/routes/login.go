package routes

import (
	appModels "src/modules/apps/models"
	authModels "src/modules/auth/models"
	userData "src/modules/users/data"
	userModels "src/modules/users/models"
	"src/shared/logging"
	"src/shared/models"
	"src/shared/response"
	"src/shared/security"
	"src/shared/utils"

	"github.com/gin-gonic/gin"
)

func Login() gin.HandlerFunc {

	return func(c *gin.Context) {

		var user userModels.User
		var post authModels.Auth
		var app appModels.AppDTO
		var err error

		if app, err = utils.CheckAppKey(c); err != nil {
			response.SendError(400, err.Error(), c)
			return
		}

		if err := c.BindJSON(&post); err != nil {
			response.SendError(400, err.Error(), c)
			return
		}

		repo := userData.NewUserData()

		if user, err = repo.GetByUsername(post.Username); err != nil {
			response.SendError(404, err.Error(), c)
			return
		}

		logging.LogSetUserID(user.ID.String(), c)

		if !security.PwdTest(post.Password, user.Password) {
			response.SendError(400, "duplicated", c)
			return
		}

		payload := &models.JwtPayload{
			UserID: user.ID.String(),
			AppID:  app.ID.String(),
		}

		var token string

		if token, err = security.JwtCreate(payload); err != nil {
			response.SendError(500, err.Error(), c)
			return
		}

		response.SendToken(token, c)

	}

}
