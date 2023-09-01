package router

import "github.com/gin-gonic/gin"

func Init() {
	r := gin.Default()
	r.GET("/task", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "the task",
		})
	})
	r.Run()
}
