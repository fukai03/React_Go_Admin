package common

import (
	"fmt"
	"ginEssential/model"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 初始化数据库
func InitDB() *gorm.DB {
	// driverName := "mysql"      // 数据库类型
	host := viper.GetString("datasource.host")         // 数据库地址
	port := viper.GetString("datasource.port")         // 数据库端口
	database := viper.GetString("datasource.database") // 数据库名
	username := viper.GetString("datasource.username") // 数据库用户名
	password := viper.GetString("datasource.password") // 数据库密码,填写mysql设置的密码
	charset := viper.GetString("datasource.charset")   // 编码方式
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	// 创建数据表
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

// 获取db实例
func GetDB() *gorm.DB {
	return DB
}
