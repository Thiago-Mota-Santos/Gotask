package router

import "github.com/gin-gonic/gin"

func Init() {
	r := gin.Default()

	Routes(r)

	r.Run()
}
