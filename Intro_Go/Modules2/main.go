package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Home Page</h1>"))
}

func serveAbout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to About Page</h1>"))
}

func main() {
	fmt.Println("MOD in golang")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/about", serveAbout).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}
