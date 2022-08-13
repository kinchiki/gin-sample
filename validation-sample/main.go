package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type Booking struct {
	CheckIn time.Time `form:"check_in" validate:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" validate:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func bookableDate(fl validator.FieldLevel) bool {
	today := time.Now()
	if date, ok := fl.Field().Interface().(time.Time); ok {
		if date.Unix() > today.Unix() {
			fmt.Println("date unix :", date.Unix())
			return true
		}
	}
	return false
}

func main() {
	serve := gin.Default()

	v := validator.New()
	v.RegisterValidation("bookabledate", bookableDate)

	serve.GET("/bookable", func(c *gin.Context) {
		var booking Booking
		if err := c.ShouldBind(&booking); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		if err := v.Struct(booking); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(),})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H {
			"message": "ok!",
			"booking": booking,
		})
	})

	serve.Run(":8080")
}
