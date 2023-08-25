package utils

import (
	"douyin/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	database, err := gorm.Open(sqlite.Open("test.DB"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB.")
	}

	DB = database
	DB.AutoMigrate(&models.User{})
}
