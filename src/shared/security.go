package shared

import (
	"src/shared/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowHeaders = []string{"*"}
	return cors.New(config)
}

func PublicEndpoints() []models.PublicEndpoint {
	return []models.PublicEndpoint{
		{Path: "/v1/login", Method: "POST"},
		{Path: "/v1/users", Method: "POST"},
		{Path: "/v1/docs", Method: "GET"},
	}
}
