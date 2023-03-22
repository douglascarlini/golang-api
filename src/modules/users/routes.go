package users

import (
	"src/modules/users/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func UserRoutes(router *gin.RouterGroup) {

	group := router.Group("/users")
	{
		group.DELETE("/:id", routes.DelOne())
		group.PUT("/:id", routes.PutOne())
		group.GET("/:id", routes.GetOne())
		group.POST("", routes.AddOne())
		group.GET("", routes.GetAll())
	}

}
