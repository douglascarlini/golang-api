package apps

import (
	"src/modules/apps/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func AppRoutes(router *gin.RouterGroup) {

	group := router.Group("/apps")
	{
		group.PATCH("/:id/renew", routes.NewKey())
		group.DELETE("/:id", routes.DelOne())
		group.PUT("/:id", routes.PutOne())
		group.GET("/:id", routes.GetOne())
		group.POST("", routes.AddOne())
		group.GET("", routes.GetAll())
	}

}
