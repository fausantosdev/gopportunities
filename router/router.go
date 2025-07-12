package router

import "github.com/gin-gonic/gin"

func Initialize() {// Names that start with an uppercase letter are automatically exported, meaning they are visible outside the package in which they were defined.
	// Initializes the router with Gin default configurations.
	router := gin.Default()
	initializeRoutes(router)

	router.Run(":8080") // Run server on 0.0.0.0:8080
}