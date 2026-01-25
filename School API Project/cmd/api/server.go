package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"restapi/internal/api/middlewares"
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

	mux := http.NewServeMux()

	// Load the TLS cert and key
	cert := "cert.pem"
	key := "key.pem"

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/teachers/", teachersHandler)
	mux.HandleFunc("/students/", studentsHandler)
	mux.HandleFunc("/execs/", execsHandler)

	// Configure TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS13,
	}

	// Create a custom server
	server := &http.Server{
		Addr:      port,
		Handler:   middlewares.SecurityHeaders(mux),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port:", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}
