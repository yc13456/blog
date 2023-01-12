package cmd

import (
	"errors"
	"fmt"
	"gin-blog/models"
)

type BlogCli struct {
	MysqlDB models.Driver
	RedisDB models.Driver
}

func NewBlogCli() *BlogCli {
	return &BlogCli{
		MysqlDB: models.NewDriver("mysql"),
		RedisDB: models.NewDriver("redis"),
	}
}

func (c *BlogCli) InitBlog() error {
	// Init Database
	err := c.MysqlDB.Connect()
	if err != nil{
		return errors.New(fmt.Sprintf("init mysql database error %v",err))
	}

	err = c.RedisDB.Connect()
	if err != nil{
		return errors.New(fmt.Sprintf("init redis database error %v",err))
	}
	return nil
}
