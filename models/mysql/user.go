package mysql

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	Name 		string `gorm:"name" json:"name" form:"name"  binding:"lte=20,required"`
	Password 	string `gorm:"password" json:"password" form:"password" binding:"lte=20,required"`
	Email 		string `gorm:"email" json:"email" form:"email" binding:"lte=50,required"`
	Phone 		string `gorm:"phone" json:"phone" form:"phone" binding:"lte=20,required"`

	Token 		string `gorm:"token" json:"token"`
	gorm.Model // 嵌入gorm.Model的字段
}

func (User) TableName() string {
	return "user"
}

func (c *User) Register (name,password,email,phone string){
	db.Create(&User{Name: name,Password: password,Email: email,Phone: phone})
}

func (c * User) Login (name,password string) error{
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