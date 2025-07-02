package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AlperSeyman/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	fmt.Println("Server is getting started...")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000....")
}
