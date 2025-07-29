package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/AlperSeyman/bookstore/pkg/models"
	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	books := models.GetAllBooks() // Query all books from the DB
	json.NewEncoder(w).Encode(books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["bookId"]
	bookId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	book, _ := models.GetByIdBook(uint(bookId))
	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var newBook models.Book
	_ = json.NewDecoder(r.Body).Decode(&newBook)

	newBook.CreateBook()
}
