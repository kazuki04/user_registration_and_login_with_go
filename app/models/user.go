package models

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) create() *User {
	u = &User{Name: u.Name, Email: u.Email}
	DbConnection.Create(u)
	return u
}
