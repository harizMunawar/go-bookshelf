package main

import (
	"fmt"
	"net/http"

	"go_bookshelf/handler"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	go router.HandleFunc("/books", handler.CreateBook).Methods("POST")
	go router.HandleFunc("/books", handler.GetAllBooks).Methods("GET")
	go router.HandleFunc("/books/{bookId}", handler.GetBook).Methods("GET")
	go router.HandleFunc("/books/{bookId}", handler.UpdateBook).Methods("PUT")
	go router.HandleFunc("/books/{bookId}", handler.DeleteBook).Methods("DELETE")
	fmt.Println("Server is up and running at localhost:8000")
	http.ListenAndServe(":8000", router)
}
