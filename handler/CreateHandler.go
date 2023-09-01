package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Create Task",
	})
}
