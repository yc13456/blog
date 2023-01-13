package mysql

import (
	"errors"
	"gorm.io/gorm"
)

const(
	SuperUser = "super"
	NormalUser = "normal"
)

type Role struct {
	Name 		string	`json:"name" gorm:"name" binding:"lte=20,required"`
	Key			string `json:"key" gorm:"key" binding:"lte=30,required"`
	Timestamp 	string `json:"timestamp"`

	gorm.Model
}

func (r *Role) TableName() string {
	return "role"
}

func (r *Role) GetKey() string{
	return r.Key
}

func (r *Role) QueryRole(name string) error{
	db.First(&r, "name = ? ", name)
	if r.Name==""{
		return errors.New("query role null")
	}
	return nil
}