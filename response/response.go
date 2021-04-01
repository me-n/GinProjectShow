package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
//服务端响应结构体
func SResponse(c *gin.Context, httpStatus int, msg string, data interface{}) {
	c.JSON(httpStatus, gin.H{
		"code": httpStatus,
		"msg":  msg,
		"data": data,
	})
}
//响应成功
func RespSuccess(c *gin.Context, msg string, data interface{}) {
	SResponse(c, http.StatusOK, msg, data)
}
//响应失败
func RespFail(c *gin.Context, httpStatus int, msg string, data interface{}) {
	SResponse(c, httpStatus, msg, data)
}
