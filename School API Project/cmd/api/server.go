package main

import (
	"fmt"
	"log"
	"net/http"
)

type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Main Page")
	w.Write([]byte("Main Page"))
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
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
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Student Page"))
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("GET Method on Students Route Page"))
	case http.MethodPost:
		w.Write([]byte("POST Method on Students Route Page"))
	case http.MethodPut:
		w.Write([]byte("PUT Method on Students Route Page"))
	case http.MethodDelete:
		w.Write([]byte("DELETE Method on Students Route Page"))
	default:
		w.Write([]byte("Students Route Page"))
	}
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Execs Page"))
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("GET Method on Execs Route Page"))
	case http.MethodPost:
		w.Write([]byte("POST Method on Execs Route Page"))
	case http.MethodPut:
		w.Write([]byte("PUT Method on Execs Route Page"))
	case http.MethodDelete:
		w.Write([]byte("DELETE Method on Execs Route Page"))
	default:
		w.Write([]byte("Execs Route Page"))
	}
}

func main() {

	port := ":3000"

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/teachers/", teachersHandler)
	http.HandleFunc("/students/", studentsHandler)
	http.HandleFunc("/execs/", execsHandler)

	fmt.Println("Server is running on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}
