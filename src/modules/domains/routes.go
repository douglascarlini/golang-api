package domains

import (
	"src/modules/domains/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func DomainRoutes(router *gin.RouterGroup) {

	group := router.Group("/domains")
	{
		group.DELETE("/:id", routes.DelOne())
		group.PUT("/:id", routes.PutOne())
		group.GET("/:id", routes.GetOne())
		group.POST("", routes.AddOne())
		group.GET("", routes.GetAll())
	}

}
