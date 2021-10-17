package models

type User struct {
	ID     uint    `gorm:"column:id;primary_key"`
	Name   string  `gorm:"column:name"`
	Email  string  `gorm:"column:email"`
	Quotes []Quote `gorm:"association_foreignkey:UserID"`
}

func (t *User) TableName() string {
	return "users"
}
