package database

import (
	"fmt"

	"github.com/AlperSeyman/fiber-crm-basic/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBconn *gorm.DB

func InitDatabase() {

	var err error

	DBconn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection opened to database")
	DBconn.AutoMigrate(&models.Lead{})
	fmt.Println("Database migrate")
}
