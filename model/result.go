package model

import "github.com/gin-gonic/gin"

func Success(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
