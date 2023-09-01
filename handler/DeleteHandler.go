package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Delete Task",
	})
}
