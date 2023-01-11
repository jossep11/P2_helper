package entities

import (
	"gorm.io/gorm"
)

type Data struct {
	gorm.Model
	Name   string `json:"name"`
	Upload string `json:"upload"`
}

type Users struct {
	gorm.Model
	UserName  string `gorm:"unique"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `gorm:"unique"`
	Password  string `gorm:"not null"`
	Age       string `json:"age"`
}
