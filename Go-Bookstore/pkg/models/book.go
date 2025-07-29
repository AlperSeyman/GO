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

func GetByIdBook(id uint) (*Book, error) {
	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func DeleteByIdBook(id uint) (*Book, error) {
	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}
	db.Delete(&book, id)
	return &book, nil
}

// CreateBook creates a new book record in DB
func (book *Book) CreateBook() *Book {
	db.Create(*book)
	return book
}
