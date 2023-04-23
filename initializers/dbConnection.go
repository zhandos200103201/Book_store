package initializers

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// simple function to connect project with database

func Connect() {
	dsn := "zhandos:SAy#wm81j5AcM$Oy@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db = database
}

// for getting db

func GetDB() *gorm.DB {
	return db
}
