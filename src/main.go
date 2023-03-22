package main

import (
	"src/modules/actions"
	"src/modules/apps"
	"src/modules/auth"
	"src/modules/doc"
	"src/modules/domains"
	"src/modules/modules"
	"src/modules/roles"
	"src/modules/users"
	"src/shared"
	"src/shared/logging"
	"src/shared/security"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	// Init router (in release mode)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Enable CORS middleware
	router.Use(shared.Cors())

	// Add Pub middleware
	router.Use(security.PubMiddleware())

	// Add Log middleware
	router.Use(logging.LogMiddleware())

	// Add JWT security
	router.Use(security.JwtMiddleware())

	// Add routes
	group := router.Group("/v1")
	domains.DomainRoutes(group)
	modules.ModuleRoutes(group)
	actions.ActionRoutes(group)
	roles.RoleRoutes(group)
	users.UserRoutes(group)
	auth.AuthRoutes(group)
	apps.AppRoutes(group)
	doc.DocRoutes(group)

	// Run server
	router.Run(":80")

}
