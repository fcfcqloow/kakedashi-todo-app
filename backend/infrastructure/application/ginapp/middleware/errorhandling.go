package middleware

import (
	domainerr "github.com/fcfcqloow/first-todo-list/backend/domain/errors"
	"github.com/fcfcqloow/go-advance/log"
	"github.com/gin-gonic/gin"
)

func ErrorHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			log.Warn(c.Errors)
			for _, err := range c.Errors {
				switch {
				case domainerr.IsNotFoundError(err):
					c.JSON(404, err.Error())
					return
				}
			}

			c.JSON(500, c.Errors.Last().Err.Error())
		}
	}
}
