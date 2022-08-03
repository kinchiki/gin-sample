package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"time"
)

func RecordUaAndTime(c *gin.Context) {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err.Error())
	}

	oldTime := time.Now()
	ua := c.GetHeader("User-Agent")
	c.Next()
	logger.Info("incoming request",
		zap.String("Path", c.Request.URL.Path),
		zap.String("User-Agent:", ua),
		zap.Int("Status", c.Writer.Status()),
		zap.Duration("Elapsed", time.Now().Sub(oldTime)),
	)
}
