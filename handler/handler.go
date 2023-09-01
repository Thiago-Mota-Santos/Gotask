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

func ShowHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "List Task",
	})
}

func UpdateHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Update Task",
	})
}

func DeleteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Delete Task",
	})
}

func ListAllHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "All Tasks",
	})
}
