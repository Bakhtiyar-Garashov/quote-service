package models

type User struct {
	ID     uint    `gorm:"column:id;primary_key"`
	Name   string  `gorm:"column:name"`
	Email  string  `gorm:"column:email"`
	Quotes []Quote `gotm:"column:quotes"`
}

func (t *User) TableName() string {
	return "users"
}
