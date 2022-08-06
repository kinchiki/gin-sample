package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kinchiki/gin-sample/api/controller"
	"github.com/kinchiki/gin-sample/api/middleware"
)

func main() {
	engine := gin.Default()

	engine.Use(middleware.RecordUaAndTime)

	bookEngine := engine.Group("/book")
	{
		v1 := bookEngine.Group("/v1")
		{
			v1.POST("/add", controller.BookAdd)
			v1.GET("/list", controller.BookList)
			v1.PUT("/update", controller.BookUpdate)
			v1.DELETE("/delete/:id", controller.BookDelete)
		}
	}
	engine.Run(":3000")
}
