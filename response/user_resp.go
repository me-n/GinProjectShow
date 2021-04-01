package response

import "GinProjectShow/model"
//用于路由中"/api/info"获取用户信息的结构体
type UserRespDto struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

//响应用户信息，数据传输对象
func RespUserDto(u model.User) UserRespDto {
	return UserRespDto{
		Name:  u.Name,
		Phone: u.Phone,
	}
}
