package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "jason"
const DB_PASSWORD = "jason123"
const DB_NAME = "bookstore_db"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var db *gorm.DB

func getDatabase() *gorm.DB {
	db = connectToDatabase()
	return db
}

func connectToDatabase() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error: %v", err)
		return nil
	}

	return db
}
