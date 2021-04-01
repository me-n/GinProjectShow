package controller

import (
	"GinProjectShow/common"
	"GinProjectShow/model"
	"GinProjectShow/response"
	"GinProjectShow/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

//注册
func Register(c *gin.Context) {
	var reqUser model.User
	c.Bind(&reqUser)
	name := reqUser.Name
	phone := reqUser.Phone
	password := reqUser.Password
	//对手机号、密码长度判断
	if len(phone) != 11 {
		response.RespFail(c, http.StatusUnprocessableEntity, "手机号必须11位", nil)
		fmt.Println(phone, len(phone))
		return
	}
	if len(password) < 6 || len(password) > 18 {
		response.RespFail(c, http.StatusUnprocessableEntity, "密码为6-18位", nil)
		return
	}
	//用户名是否为空，为空为其生成随机名字
	if len(name) == 0 {
		name = util.CreateName()
	}
	if isPhoneExist(common.DB, phone) {
		response.RespFail(c, http.StatusUnprocessableEntity, "手机号已存在", nil)
		return
	}
	//创建用户
	//对密码进行加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.RespFail(c, http.StatusInternalServerError, "加密错误", err)
		return
	}
	user := model.User{
		Name:     name,
		Password: string(hashPassword),
		Phone:    phone,
	}
	//将用户存入mysql中users表
	common.DB.Create(&user)
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.RespFail(c, http.StatusInternalServerError, "用户数据持久化错误", err)
		return
	}
	response.RespSuccess(c, "注册成功", gin.H{"token": token})
}

//登陆
func Login(c *gin.Context) {
	var reqUser model.User
	c.Bind(&reqUser)
	phone := reqUser.Phone
	password := reqUser.Password
	if len(phone) != 11 {
		response.RespFail(c, http.StatusUnprocessableEntity, "手机号必须11位", nil)
		return
	}
	if len(password) < 6 || len(password) > 18 {
		response.RespFail(c, http.StatusUnprocessableEntity, "密码为6-18位", nil)
		return
	}
	//依据手机查询注册用户信息
	var user model.User
	common.DB.Where("phone=?", phone).First(&user)
	if user.ID == 0 {
		response.RespFail(c, http.StatusUnprocessableEntity, "用户不存在", nil)
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.RespFail(c, http.StatusUnprocessableEntity, "密码错误", err)
		return
	}
	//为用户下发token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.RespFail(c, http.StatusInternalServerError, "token下发失败", err)
		return
	}
	response.RespSuccess(c, "登陆成功", gin.H{"name": user.Name,
		"token": token})
}

//获取用户信息
func GetUserInfo(c *gin.Context) {
	user, _ := c.Get("user") //返回值为空接口 key对应中间件设置的key 2@1
	response.RespSuccess(c, "响应成功", gin.H{
		"user": response.RespUserDto(user.(model.User))}) //将user断言称为model.User
}

//验证手机号是否存在
func isPhoneExist(db *gorm.DB, phone string) bool {
	var u model.User
	db.Where("phone=?", phone).First(&u)
	if u.ID != 0 {
		return true
	}
	return false
}
