package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Update Task",
	})
}
