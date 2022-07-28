package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	ua := ""

	engine.Use(func(c *gin.Context) {
		ua = c.GetHeader("User-Agent")
		c.Next()
	})

	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message": "Hello World!",
			"User-Agent": ua,
		})
	})
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message": "pong",
		})
	})
	engine.Run()
}
