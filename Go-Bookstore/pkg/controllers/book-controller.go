package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/AlperSeyman/bookstore/pkg/config"
	"github.com/AlperSeyman/bookstore/pkg/models"
	"github.com/AlperSeyman/bookstore/pkg/utils"
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
		return
	}

	book, _ := models.GetByIdBook(uint(bookId))
	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var book models.Book
	err := utils.ParseBody(r, &book)
	if err != nil {
		log.Fatal(err)
		return
	}
	newBook := book.CreateBook()
	json.NewEncoder(w).Encode(newBook)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Fatal(err)
		return
	}

	book, err := models.DeleteByIdBook(uint(id))
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Fatal(err)
		return
	}

	currentBook, err := models.GetByIdBook(uint(id))
	if err != nil {
		log.Fatal(err)
		return
	}

	var updateBook models.Book

	err = utils.ParseBody(r, &updateBook)
	if err != nil {
		log.Fatal(err)
	}

	if updateBook.Author != "" {
		currentBook.Author = updateBook.Author
	}

	if updateBook.Name != "" {
		currentBook.Name = updateBook.Name
	}

	if updateBook.Publication != "" {
		currentBook.Publication = updateBook.Publication
	}

	config.DB.Save(&currentBook)
	json.NewEncoder(w).Encode(currentBook)

}
