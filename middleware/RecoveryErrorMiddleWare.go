package middleware

import (
	"GinProjectShow/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

//错误捕获
func RecoverErrorMidW() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.RespFail(c, 400, fmt.Sprint(err), nil)
				c.Abort()
				return
			}
		}()
	}
}
