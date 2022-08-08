package main

import (
	"github.com/gin-gonic/gin"
  "net/http"
)

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
	serve.Run(":8081")
}
