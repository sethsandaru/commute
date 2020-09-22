package routes

import (
	"github.com/gin-gonic/gin"
	controllers "mach/app/controllers"
)

var machRouter *gin.Engine

// Registering Route Method
func InitializeRoutes(router *gin.Engine) {
	// publish the router so we can registering it within the `routes` package
	machRouter = router

	// Index Route
	machRouter.GET("/", controllers.ShowIndexPage)
}