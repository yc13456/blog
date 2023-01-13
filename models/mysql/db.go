package mysql

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"gin-blog/utils"
	"gorm.io/driver/mysql"

)

var db *gorm.DB


type mysqlDb struct {
	sql utils.SqlParam
}

func NewMysqlDB() *mysqlDb{
	return &mysqlDb{
		sql: utils.SqlParam{
			UserName: "root",
			PassWord: "123456",
			Host:     "192.168.1.106",
			Port:     3306,
			Dbname :  "blog",
			Timeout:  "10s",
		},
	}
}

func (c *mysqlDb) Connect() error{
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		c.sql.UserName, c.sql.PassWord, c.sql.Host, c.sql.Port, c.sql.Dbname, c.sql.Timeout)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New("连接数据库失败")
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Role{})
	return nil
}
