package models

type User struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Email  string  `gorm:"unique_index" json:"email"`
	Quotes []Quote `json:"quotes"`
}

func (t *User) TableName() string {
	return "users"
}
