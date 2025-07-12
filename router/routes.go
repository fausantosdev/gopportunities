package router // Everything within the same package is accessible to the entire package

import (
	"github.com/fausantosdev/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	{
		v1 := router.Group("/api/v1")

		v1.POST("/oppening", handler.CreateOpeningHandler)

		v1.GET("/oppening", handler.ReadOpeningHandler)

		v1.PUT("/oppening", handler.UpdateOpeningHandler)

		v1.DELETE("/oppening", handler.DeleteOpeningHandler)
	}
}