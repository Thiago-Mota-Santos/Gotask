package router

import (
	"github.com/Thiago-Mota-Santos/Gotask/handler"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	api := router.Group("/api/")
	{
		api.GET("/task", handler.ShowHandler)
		api.POST("/task", handler.CreateHandler)
		api.PUT("/task", handler.UpdateHandler)
		api.DELETE("/task", handler.DeleteHandler)
		api.GET("/tasks", handler.ListAllHandler)
	}
}
