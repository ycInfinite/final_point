package controller

import (
	"finalpoint/dao"
	"finalpoint/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func SelectAll(c *gin.Context) {
	booklist := dao.SelectAll()
	c.JSON(200, booklist)
}

func Page(c *gin.Context) {
	pageID := c.Param("pageid")

	idStr := pageID[4:5]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fmt.Println("转换失败： ", err)
		return
	}
	begin := (id - 1) * 20
	pagesize := 20
	bookList := dao.Page(uint(begin), pagesize)

	model.Success(c, 200, "ok", bookList)
}
