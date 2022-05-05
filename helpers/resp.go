package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespWithOK(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
		"msg":  message,
	})
}

func RespWithBadRequest(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusBadRequest,
		"data": data,
		"msg":  message,
	})
}

func RespWithSystemError(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusInternalServerError,
		"data": data,
		"msg":  message,
	})
}

func RespWithUnauthorized(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusUnauthorized,
		"data": data,
		"msg":  message,
	})
}
