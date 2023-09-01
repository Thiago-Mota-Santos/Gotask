package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListAllHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "All Tasks",
	})
}
