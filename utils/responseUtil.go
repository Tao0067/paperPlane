package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": -1,
		"msg":  msg,
	})
}

func SucceedData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}
func SucceedMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  msg,
	})
}
