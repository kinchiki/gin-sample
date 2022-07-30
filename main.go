package main

import (
	"net/http"
	"io"
	"log"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	ua := ""

	engine.LoadHTMLGlob("templates/*")
	engine.Use(func(c *gin.Context) {
		ua = c.GetHeader("User-Agent")
		c.Next()
	})

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H {
			"message": "Hello World!",
			"userAgent": ua,
		})
	})

	engine.Static("/static", "./static")

	engine.GET("/image-upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "image-upload.html", gin.H {
		})
	})

	// file upload
	engine.POST("/image-upload", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("image")
		if err != nil {
			c.String(http.StatusBadRequest, "Bad request!")
			return
		}
		fileName := header.Filename
		dir, _ := os.Getwd()
		out, err := os.Create(dir+"/images/"+fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H {
			"status": "ok",
		})
	})

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message": "pong",
		})
	})
	engine.Run()
}
