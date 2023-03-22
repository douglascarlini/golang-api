package doc

import (
	"encoding/json"
	"os"
	"src/shared/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func DocRoutes(router *gin.RouterGroup) {

	group := router.Group("/docs")
	{
		group.GET("", func(c *gin.Context) {

			text, _ := os.ReadFile("swagger.json")
			var data models.SwaggerJson

			if err := json.Unmarshal(text, &data); err != nil {
				c.AbortWithStatus(500)
				return
			}

			url := os.Getenv("API_URL")
			url = strings.ReplaceAll(url, "http://", "")
			url = strings.ReplaceAll(url, "https://", "")
			data.Host = url

			c.IndentedJSON(200, data)
		})

	}

}
