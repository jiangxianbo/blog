// named_return1.go
package main

import (
	"blog/controller"
	"blog/dao/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	dns := "root:jiang123@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}
	// 加载静态文件
	router.Static("/static/", "./static")
	router.LoadHTMLGlob("views/*")
	router.GET("/", controller.IndexHandler)
	router.GET("/category", controller.CategoryList)
	articleGroup := router.Group("/article")
	{
		articleGroup.GET("/detail", controller.ArticleDetail)
	}
	router.Run(":8001")
}
