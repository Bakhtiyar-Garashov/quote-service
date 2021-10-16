package models

type User struct {
	ID     uint    `gorm:"primary_key" json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Quotes []Quote `json:"quotes"`
}
