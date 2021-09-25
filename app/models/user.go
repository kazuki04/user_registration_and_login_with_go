package models

import (
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

type User struct {
	Name  string
	Email string
	gorm.Model
}

func (u *User) create() {
	u = &User{Name: u.Name, Email: u.Email}
	result := db.Create(u)
}
