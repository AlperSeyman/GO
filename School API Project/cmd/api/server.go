package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

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

			// Parse form data (necessary for x-www-form-urlencoded)

			// err := r.ParseForm()
			// if err != nil {
			// 	http.Error(w, "Error parsing form", http.StatusBadRequest)
			// 	return
			// }
			// fmt.Println("Form:", r.Form)

			// // Prepare response data
			// response := make(map[string]any)
			// for key, value := range r.Form {
			// 	response[key] = value[0]
			// }
			// fmt.Println("Processed Response Map:", response)

			// name := r.FormValue("name")
			// fmt.Println("Name: ", name)

			// RAW Body

			body, err := io.ReadAll(r.Body)
			if err != nil {
				return
			}
			defer r.Body.Close()

			fmt.Println("RAW Body:", body)
			fmt.Println("RAW Body:", string(body))

			// If we expect json data, then unmarshal it.
			var user user
			err = json.Unmarshal(body, &user)
			if err != nil {
				return
			}
			fmt.Println("Unmarshaled JSON into an instance of user struct: ", user)
			fmt.Println("Recieved user name as:", user.Name)

			// // Prepare response data
			response := make(map[string]any)
			for key, value := range r.Form {
				response[key] = value[0]
			}

			err = json.Unmarshal(body, &response)
			if err != nil {
				return
			}
			fmt.Println("Unmarshaled JSON into a map: ", response)

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

	fmt.Println("Server is running on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}
