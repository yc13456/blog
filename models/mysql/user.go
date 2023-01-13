package mysql

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	Name 		string `gorm:"name" json:"name" form:"name"  binding:"lte=20,required"`
	Password 	string `gorm:"password" json:"password" form:"password" binding:"lte=20,required"`
	Email 		string `gorm:"email" json:"email" form:"email" binding:"lte=50,required"`
	Phone 		string `gorm:"phone" json:"phone" form:"phone" binding:"lte=20,required"`

	RoleID 		int `gorm:"roleID,ForeignKey:role" json:"roleID" json:"role_id"`
	Token 		string `gorm:"token" json:"token"`
	gorm.Model // 嵌入gorm.Model的字段
}

func (User) TableName() string {
	return "user"
}

func (c *User) Register (name,password,email,phone,token string) error{
	var u User

	if db.First(&u,"name = ?",name,phone, email); u.Name != ""{
		return errors.New(fmt.Sprintf("register error,name: %s is exist",name))
	}

	if db.First(&u,"email = ? ",email); u.Email != ""{
		return errors.New(fmt.Sprintf("register error,email: %s is exist",u.Email))
	}

	if db.First(&u,"phone = ? ",phone); u.Phone != ""{
		return errors.New(fmt.Sprintf("register error,phone: %s is exist",u.Email))
	}

	db.Create(&User{Name: name,Password: password,Email: email,Phone: phone,Token: token})
	return nil
}

func (c * User) Login (name,password,token string) error{
	var user User

	if db.First(&user, "token = ? ", token); user.Name != ""{
		fmt.Println("token is expired")
		return nil
	}

	db.First(&user, "name = ? ", name)
	if user.Password == ""{
		return errors.New("user no found")
	}
	if password == user.Password{
		return nil
	} else {
		return errors.New("password is error")
	}
}