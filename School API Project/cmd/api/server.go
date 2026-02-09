package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"restapi/internal/api/routers"
	"restapi/internal/repository/sqlconnect"
)

func main() {

	_, err := sqlconnect.ConnectDB()
	if err != nil {
		fmt.Println("Error------: ", err)
	}

	port := os.Getenv("SERVER_PORT")

	// Load the TLS cert and key
	cert := "cert.pem"
	key := "key.pem"

	// Configure TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS13,
	}

	// Create a custom server
	server := &http.Server{
		Addr:    port,
		Handler: routers.Router(),
		// Handler:   middlewares.SecurityHeaders(middlewares.CORSMiddleware(mux)),
		// Handler:   middlewares.CORSMiddleware(mux),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port:", port)
	err = server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}
