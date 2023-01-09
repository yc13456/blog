package models

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	Name string `gorm:"name" json:"name" form:"name"`
	Password string `gorm:"password" json:"password" form:"password"`
	Email string `gorm:"email" json:"email" form:"email"`
	Phone string `gorm:"phone" json:"phone" form:"phone"`

	gorm.Model // 嵌入gorm.Model的字段
}

func (User) TableName() string {
	return "user"
}

func (c *User)Register(name,password,email,phone string){
	db.Create(&User{Name: name,Password: password,Email: email,Phone: phone})
}

func (c * User)Login(name,password string) error{
	var user User
	db.First(&user, "name = ? ", name)
	if user.Password == ""{
		return errors.New("user no found")
	}
	if c.Password == user.Password{
		return nil
	} else {
		return errors.New("password is error")
	}
}