package mid

import (
	"github.com/GostGrimmy/gdl/http"
	"github.com/gin-gonic/gin"
)

type ErrorHandler func(rec interface{}, c *gin.Context)

func GlobalRecover(handlers ...ErrorHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		wb := http.NewBodyWriter(c)
		c.Writer = wb
		wb.NewWriter()
		defer func(c *gin.Context) {
			rec := recover()
			if rec != nil {
				for _, handler := range handlers {
					handler(rec, c)
				}
			}
		}(c)
		c.Next()
		wb = c.Writer.(*http.BodyWriter)
		wb.Done()
	}
}
