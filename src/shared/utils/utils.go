package utils

import (
	"fmt"
	"src/shared/logging"
	"src/shared/response"
	"src/shared/security"
	"strings"

	appData "src/modules/apps/data"
	appModels "src/modules/apps/models"
	sharedModels "src/shared/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CheckAppKey(c *gin.Context) (appModels.AppDTO, error) {

	var app appModels.AppDTO
	var key string
	var err error

	if key = c.Request.Header.Get("key"); key == "" {

		err = fmt.Errorf("missing key")
		response.SendError(400, err.Error(), c)
		return app, err

	}

	if _, err := uuid.Parse(key); err != nil {

		err = fmt.Errorf("missing key")
		response.SendError(400, err.Error(), c)
		return app, err

	}

	repo := appData.NewAppData()

	if app, err = repo.GetBySecret(key); err != nil {

		response.SendError(500, err.Error(), c)
		return app, err

	}

	if app.ID == uuid.Nil {

		err = fmt.Errorf("invalid app")
		response.SendError(400, err.Error(), c)
		return app, err

	}

	logging.LogSetAppID(app.ID.String(), c)

	return app, nil
}

func FindACL(need []string, c *gin.Context) bool {

	var err error
	var jwt *sharedModels.JwtPayload

	if jwt, err = security.JwtGetPayload(c); err != nil {
		response.SendError(500, err.Error(), c)
		return false
	}

	need = append(need, "#")
	acls := strings.Split(jwt.ACL, " ")

	for _, this := range acls {
		for _, that := range need {
			if this == that {
				return true
			}
		}
	}

	return false
}
