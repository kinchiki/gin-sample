package main

import (
	"github.com/gin-gonic/gin"
  "net/http"
)

type Person struct {
	Age int `form:"age" validate:"required,gt=10,max=100"`
	Name string `form:"name" validate:"required"`
	Address string `form:"address" validate:"required"`
}

func main() {
	serve := gin.Default()
	serve.GET("/test", func(c *gin.Context) {
		firstName := c.Query("first_name")
		lastName := c.DefaultQuery("last_name", "default_name")
		c.JSON(http.StatusOK, gin.H {
			"firstName": firstName,
			"lastName": lastName,
		})
	})

	serve.POST("/test", func(c *gin.Context) {
		firstName := c.PostForm("first_name")
		lastName := c.DefaultPostForm("last_name", "default_last_name")
		c.JSON(http.StatusOK, gin.H {
			"firstName": firstName,
			"lastName": lastName,
		})
	})

	serve.GET("/testing", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.JSON(500, gin.H{"msg:": err.Error()})
			return
		}
		c.JSON(200, gin.H{"person": person})
	})

	serve.Run(":8081")
}
