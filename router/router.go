package router

import "github.com/gin-gonic/gin"

func Initialize() {
		// Inicializa o router com as configs defaut do gin
	router := gin.Default()
	
	// Definição de rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
		"message": "pong",
		})
	})
	router.Run(":8080") // Listen and serve on 0.0.0.0:8080
}