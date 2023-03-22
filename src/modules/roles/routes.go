package roles

import (
	"src/modules/roles/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func RoleRoutes(router *gin.RouterGroup) {

	group := router.Group("/roles")
	{
		group.DELETE("/:id", routes.DelOne())
		group.PUT("/:id", routes.PutOne())
		group.GET("/:id", routes.GetOne())
		group.POST("", routes.AddOne())
		group.GET("", routes.GetAll())
	}

}
