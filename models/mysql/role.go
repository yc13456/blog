package mysql

type Role struct {
	ID			string `json:"id" gorm:"id"`
	Key			string `json:"key" gorm:"key"`
	Timestamp 	string `json:"timestamp"`
}

func (r *Role) TableName() string {
	return "role"
}

