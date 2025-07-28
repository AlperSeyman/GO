package main

import (
	"log"
	"net/http"

	"github.com/AlperSeyman/bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	err := http.ListenAndServe(":5555", r)
	if err != nil {
		log.Fatal(err)
	}
}
