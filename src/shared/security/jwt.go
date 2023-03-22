package security

import (
	"fmt"
	"log"
	"os"
	"src/shared/logging"
	"src/shared/models"
	"src/shared/response"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		if _, ok := c.Get("public"); ok {
			c.Next()
			return
		}

		var authHeader string

		if authHeader = c.Request.Header.Get("Authorization"); authHeader == "" {
			response.SendError(401, " missing token", c)
			return
		}

		obj, err := jwt.ParseWithClaims(authHeader[len("Bearer "):], &models.JwtPayload{}, func(obj *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			response.SendError(401, err.Error(), c)
			return
		}

		if payload, ok := obj.Claims.(*models.JwtPayload); ok && obj.Valid {

			if payload.AppID == "" {
				response.SendError(401, "app key", c)
				return
			}

			if payload.UserID == "" {
				response.SendError(401, "user id", c)
				return
			}

			logging.LogSetAccess(payload.AppID, payload.UserID, c)
			c.Set("jwt", payload)
			c.Next()
			return

		}

		response.SendError(401, "", c)

	}
}

func JwtCreate(payload *models.JwtPayload) (string, error) {

	var secret = []byte(os.Getenv("JWT_SECRET"))
	payload.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	obj := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := obj.SignedString(secret)
	if err != nil {
		log.Print(err.Error())
		return "", err
	}

	return token, nil
}

func JwtGetPayload(c *gin.Context) (*models.JwtPayload, error) {

	if payload, ok := c.Get("jwt"); ok {
		return payload.(*models.JwtPayload), nil
	}

	return &models.JwtPayload{}, fmt.Errorf("unauthorized")

}
