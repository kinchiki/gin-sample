package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/kinchiki/gin-sample/api/model"
	"github.com/kinchiki/gin-sample/api/service"
	"strconv"
)

func BookAdd(c *gin.Context) {
	book := model.Book{}
	err := c.Bind(&book)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	bookService := service.BookService{}
	err = bookService.SetBook(&book)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H {
		"status": "ok",
	})
}

func BookList(c *gin.Context) {
	bookService := service.BookService{}
	BookLists := bookService.GetBookList()
	c.JSONP(http.StatusOK, gin.H {
		"message": "ok",
		"data": BookLists,
	})
}

func BookUpdate(c *gin.Context) {
	book := model.Book{}
	err := c.Bind(&book)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	bookService := service.BookService{}
	err = bookService.UpdateBook(&book)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func BookDelete(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 0)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	bookService := service.BookService{}
	err = bookService.DeleteBook(int(id))
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H {
		"status": "ok",
	})
}
