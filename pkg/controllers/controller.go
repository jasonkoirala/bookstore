package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jasonkoirala/bookstore/pkg/models"
	"github.com/jasonkoirala/bookstore/pkg/utils"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Get Book By Id...")

	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while Parsing...")
	}
	bookDetails, _ := models.GetBookByID(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseRequestBody(r, book)
	b := book.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while Parsing...")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Update Book...")
	params := mux.Vars(r)

	var newBook = &models.Book{}
	utils.ParseRequestBody(r, newBook)
	bookId := params["bookId"]
	fmt.Println(bookId)
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while Parsing...")
	}
	bookDetail, db := models.GetBookByID(ID)
	fmt.Println(bookDetail.Publication)
	if newBook.Name != "" {
		bookDetail.Name = newBook.Name
	}
	if newBook.Author != "" {
		bookDetail.Author = newBook.Author
	}
	if newBook.Publication != "" {
		bookDetail.Publication = newBook.Publication
	}
	db.Save(&bookDetail)
	res, _ := json.Marshal(bookDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
