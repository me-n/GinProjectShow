package route

import (
	"GinProjectShow/controller"
	"GinProjectShow/middleware"
	"github.com/gin-gonic/gin"
)
//路由设置
func CollectRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleWare(),middleware.RecoverErrorMidW())
	r.POST("/api/register",controller.Register)
	r.POST("/api/login",controller.Login)
	r.GET("/api/info",middleware.TokenMiddleWare(),controller.GetUserInfo)
	return r
}
