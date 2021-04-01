package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
//数据库连接初始化
func InitDB() *gorm.DB {
	var err error
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")
	database := viper.GetString("datasource.database")
	sqlPath := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username, password, host, port, database, charset, loc)
	fmt.Println(sqlPath)
	DB, err = gorm.Open(mysql.Open(sqlPath), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败，" + err.Error())
	}
	return DB
}
