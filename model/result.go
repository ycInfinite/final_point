package model

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
