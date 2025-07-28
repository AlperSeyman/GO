package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/AlperSeyman/bookstore/pkg/models"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	books := models.GetAllBooks() // Query all books from the DB
	json.NewEncoder(w).Encode(books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

}
