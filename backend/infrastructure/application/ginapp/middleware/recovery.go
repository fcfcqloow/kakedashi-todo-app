package middleware

import (
	"fmt"
	"runtime"

	"github.com/fcfcqloow/go-advance/log"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				result := ""
				for depth := 0; ; depth++ {
					_, file, line, ok := runtime.Caller(depth)
					if !ok {
						break
					}
					result += fmt.Sprintf("======> %d: %v:%d\n", depth, file, line)
				}
				c.AbortWithStatusJSON(500, gin.H{"message": "Internal Server Error"})
				log.Error(result)
			}
		}()

		c.Next()
	}
}
