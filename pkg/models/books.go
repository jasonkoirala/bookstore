package models

import (
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.connectToDatabase()
	db = config.getDatabase()
	db.AutoMigrate(&Book{})
}