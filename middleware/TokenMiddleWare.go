package middleware

import (
	"GinProjectShow/common"
	"GinProjectShow/model"
	"GinProjectShow/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//token中间件
func TokenMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := "me-n:"
		//在Headers中添加： key：Authorization；value：me-n:xxx(token值)进行测试
		tokenStr := c.GetHeader("Authorization")
		//验证token是否为空，或者前缀是否为"me-n:"
		if tokenStr == "" || !strings.HasPrefix(tokenStr, auth) {
			response.RespFail(c, http.StatusUnauthorized, "权限不足", nil)
			c.Abort()
			return
		}
		index := strings.Index(tokenStr, auth)
		//获取真实token值
		tokenStr = tokenStr[index+len(auth):]
		token, claims, err := common.ParseToken(tokenStr)
		if err != nil || !token.Valid {
			response.RespFail(c, http.StatusBadRequest, "解析错误或Token已过期", err)
			c.Abort()
			return
		}
		id := claims.UserId
		var user model.User
		common.DB.First(&user, id)
		//如未读取到值，说明token有误
		if user.ID == 0 {
			response.RespFail(c, http.StatusUnauthorized, "权限不足", nil)
			c.Abort()
			return
		}
		//将user 的值注入context中 1@2
		c.Set("user", user)
		c.Next()
	}
}
