package models

import (
	"gin-blog/models/mysql"
	"gin-blog/models/redis"
)

type Driver interface {
	Connect() error
}

func NewDriver(name string)(driver Driver){
	switch name {
	case "mysql":
		driver = mysql.NewMysqlDB()
	case "redis":
		driver = redis.NewRedisDB()
	}
	return
}
