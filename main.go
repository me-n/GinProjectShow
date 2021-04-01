package main

import (
	"GinProjectShow/common"
	"GinProjectShow/route"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	//配置初始化
	InitConfig()
	//数据库连接初始化
	common.InitDB()
	r := gin.Default()
	//开启路由
	r = route.CollectRouter(r)
	//获取配置中服务器port
	port := viper.GetString("server.port")
	if port != "" {
		r.Run(":" + port)
	} else {
		//如未能成功读取配置端口，默认启动8080
		r.Run()
	}
}
//配置文件初始化
func InitConfig() {
	getPath, err := os.Getwd() //获取项目目录路径
	if err != nil {
		panic(err)
	}
	viper.SetConfigName("apps")              //配置文件名（不带.yml后缀）
	viper.SetConfigType("yml")               //配置文件类型
	viper.AddConfigPath(getPath + "/config") //配置执行路径
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

/*
本项目由MN(MengLin)编写，基于gin框架，实现注册、登陆、信息获取等功能的后端源代码0.01初级版本，
实现mysql用户注册、登陆查询、身份校验等，服务器token发放 验证，中间件、接口实现、路由、名字自动生成等。
mysql文件可导入数据库，更改config包中配置信息，进行测试
*/
