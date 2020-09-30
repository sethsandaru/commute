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

	// Profile Routes
	machRouter.GET("/profile", controllers.GetProfileList)
	machRouter.GET("/profile/:id", controllers.GetSingleProfile)
	machRouter.POST("/profile", controllers.InsertProfile)
	machRouter.PUT("/profile/:id", controllers.UpdateProfile)
	machRouter.DELETE("/profile/:id", controllers.DeleteProfile)
}