package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	const serverAddr string = "127.0.0.1:3000"

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello server")
	})

	http.HandleFunc("/pizza", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pizza Menu")
	})

	http.HandleFunc("/soup", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Soup Menu")
	})

	fmt.Println("Server listening on port 3000")
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		log.Fatal("error starting server: ", err)
	}
}
