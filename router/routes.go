package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	api := router.Group("/api/")
	{
		api.GET("/task", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "One Task",
			})
		})

		api.POST("/task", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "POST Task",
			})
		})

		api.PUT("/task", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "PUT Task",
			})
		})

		api.DELETE("/task", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Delete Task",
			})
		})

		api.GET("/tasks", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "get all Tasks",
			})
		})
	}
}
