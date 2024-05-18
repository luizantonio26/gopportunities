package router

import (
	"github.com/gin-gonic/gin"
	"github.com/luizantonio26/gopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	//initialize handler
	handler.InitializerHandler()

	v1 := router.Group("/api/v1/")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
	}
}
