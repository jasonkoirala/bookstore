package routes

import (
	"github.com/gorilla/mux"
	"github.com/jasonkoirala/"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.createBook).Methods("POST")
	router.HandleFunc("/book/", controllers.getBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.getBookById).Methods("GET")
	router.HandleFunc("book/{bookId}", controllers.updateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.deleteBook).Methods("DELETE")

}
