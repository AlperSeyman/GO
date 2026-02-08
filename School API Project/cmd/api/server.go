package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"restapi/internal/api/routers"
)

func main() {

	port := ":3000"

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
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}
