package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	port := ":3000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Main Page")
		w.Write([]byte("Main Page"))
	})

	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Main Page")

		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("GET Method on Teachers Route Page"))
		case http.MethodPost:
			w.Write([]byte("POST Method on Teachers Route Page"))
		case http.MethodPut:
			w.Write([]byte("PUT Method on Teachers Route Page"))
		case http.MethodDelete:
			w.Write([]byte("DELETE Method on Teachers Route Page"))
		default:
			w.Write([]byte("Teachers Route Page"))
		}

	})

	// http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
	// 	// fmt.Fprintf(w, "Main Page")
	// 	w.Write([]byte("Students Page"))
	// })

	// http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
	// 	// fmt.Fprintf(w, "Main Page")
	// 	w.Write([]byte("Execs Page"))
	// })

	fmt.Println("Server is running on port: ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}
