package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`       // convert to json type
	Title    string `json:"title"`    // convert to json type
	Author   string `json:"author"`   // convert to json type
	Quantity int    `json:"quantity"` // convert to json type
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, b := range books {
		if b.ID == id {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
		return
	}

	for i, b := range books {
		if b.ID == id {
			if books[i].Quantity > 0 {
				books[i].Quantity--
				c.IndentedJSON(http.StatusOK, books[i])
				return
			} else {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book is not avaliable"})
				return
			}
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func postBook(c *gin.Context) { // create book
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	for _, b := range books {
		if b.ID == newBook.ID {
			c.JSON(http.StatusConflict, gin.H{"error": "Book with this ID already exists"})
			return
		}
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.PATCH("/checkout", checkoutBook)
	router.POST("/books", postBook)
	router.Run("localhost:8080")
}
