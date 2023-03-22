package modules

import (
	"src/modules/modules/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func ModuleRoutes(router *gin.RouterGroup) {

	group := router.Group("/modules")
	{
		group.DELETE("/:id", routes.DelOne())
		group.PUT("/:id", routes.PutOne())
		group.GET("/:id", routes.GetOne())
		group.POST("", routes.AddOne())
		group.GET("", routes.GetAll())
	}

}
