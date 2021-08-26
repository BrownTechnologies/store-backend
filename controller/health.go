package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Server is running ...",
	})
}
