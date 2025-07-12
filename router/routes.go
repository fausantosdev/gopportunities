package router // Everything within the same package is accessible to the entire package

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	{
		v1 := router.Group("/api/v1")

		v1.POST("/oppening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "It's Ok",
			})
		})

		v1.GET("/oppening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "It's Ok",
			})
		})

		v1.PUT("/oppening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "It's Ok",
			})
		})

		v1.DELETE("/oppening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "It's Ok",
			})
		})
	}
}