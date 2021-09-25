package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var DbConnection *gorm.DB

func init() {
	var err error
	dsn := "root:@/user_registration_and_login_with_go"
	DbConnection, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DbConnection.AutoMigrate(&User{})
}
