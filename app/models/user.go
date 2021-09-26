package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) Save() *User {
	u = &User{Name: u.Name, Email: u.Email}
	DbConnection.Create(u)
	return u
}
