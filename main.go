package main

import (
	"finalpoint/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	//dao.Db.AutoMigrate(&model.BookFirstMsg{})
	//var bk model.BookFirstMsg
	//dao.Db.First(&bk)
	//fmt.Printf("%+v", bk)
	r := gin.Default()
	controller.Router(r)
	r.Run()
}
