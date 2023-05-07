package middleware

import (
	"github.com/fcfcqloow/go-advance/log"
	"github.com/gin-gonic/gin"
)

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info(c.Request.Method, c.Request.URL.String())
		c.Next()
	}
}
