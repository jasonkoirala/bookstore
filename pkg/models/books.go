package models

import (
	"log"

	"github.com/jasonkoirala/bookstore/pkg/config"
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
	config.ConnectToDatabase()
	db = config.GetDatabase()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("id=?", id).Find(&book)
	return &book, db
}

func DeleteBook(id int64) *Book {
	log.Println(id)
	log.Println("...bout to get deleted...")
	var book Book
	db.Where("id=?", id).Delete(&book)
	return &book
}
