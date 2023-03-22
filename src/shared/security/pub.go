package security

import (
	"src/shared"

	"github.com/gin-gonic/gin"
)

func PubMiddleware() gin.HandlerFunc {

	public := shared.PublicEndpoints()

	return func(c *gin.Context) {

		path := c.Request.URL.Path
		method := c.Request.Method

		if method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		if path == "/favicon.ico" {
			c.AbortWithStatus(404)
			return
		}

		for _, item := range public {
			if method == item.Method && item.Path == path {
				c.Set("public", true)
				c.Next()
				return
			}
		}

		c.Next()

	}

}
