package models

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error{
	//配置MySQL连接参数
	username := "root"  //账号
	password := "123456" //密码
	host := "192.168.1.116" //数据库地址，可以是Ip或者域名
	port := 3306 //数据库端口
	Dbname := "blog" //数据库名
	timeout := "10s" //连接超时，10秒

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New("连接数据库失败")
	}
	db.AutoMigrate(&User{})
	return nil
}
