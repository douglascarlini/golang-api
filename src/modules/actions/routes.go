package actions

import (
	"src/modules/actions/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func ActionRoutes(router *gin.RouterGroup) {

	group := router.Group("/actions")
	{
		group.DELETE("/:id", routes.DelOne())
		group.PUT("/:id", routes.PutOne())
		group.GET("/:id", routes.GetOne())
		group.POST("", routes.AddOne())
		group.GET("", routes.GetAll())
	}

}
