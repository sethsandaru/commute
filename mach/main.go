package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"mach/routes"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Application can't be started")
		os.Exit(1)
	}

	// Create Gin Instance
	router := gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("app/views/*")

	// Registering the routes
	routes.InitializeRoutes(router)

	// start serving the application
	router.Run(":" + os.Getenv("APPLICATION_PORT"))
}