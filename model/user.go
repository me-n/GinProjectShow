package model

import "gorm.io/gorm"
//定义用户结构体
type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Phone    string `gorm:"varchar(11);not null;unique"`
	Password string `gorm:"size:255;not null"`
}
