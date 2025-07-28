package models

import (
	"github.com/AlperSeyman/bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB // Declares a private variable db to hold the GORM DB connection

type Book struct {
	gorm.Model         // adds ID, CreatedAt, UpdatedAt, DeletedAt
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()        // Connect to the DB
	db = config.GetDB()     // Get DB instance
	db.AutoMigrate(&Book{}) // Auto-create the books table
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetByIdBook(id uint) (*Book, *gorm.DB) {
	var book Book
	db := db.First(&book, id)
	return &book, db
}

// CreateBook creates a new book record in DB
func (b *Book) CreateBook() *Book {
	db.Create(*b)
	return b
}
