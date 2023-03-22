package auth

import (
	"src/modules/auth/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func AuthRoutes(router *gin.RouterGroup) {

	group := router.Group("/login")
	{
		group.POST("", routes.Login())
	}

}
