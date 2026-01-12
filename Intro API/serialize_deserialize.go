package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	user1 := User{
		Name:  "Alice",
		Email: "alice@example.com",
	}

	jsonData1, err := json.Marshal(user1) // convert to JSON type
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData1))

	var user2 User
	err = json.Unmarshal(jsonData1, &user2) // convert to Go Struct type
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user2)

	jsonData2 := `
		{"name":"John", "email":"john@example.com"}
	`
	reader := strings.NewReader(jsonData2) // receive json
	decoder := json.NewDecoder(reader)     // create decoder to convert json to struct type

	var user3 User
	err = decoder.Decode(&user3) // convert to Go struct type
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user3)

	var jsonData3 bytes.Buffer
	encoder := json.NewEncoder(&jsonData3) // create encoder to convert struct to json type
	err = encoder.Encode(user1)            // convert to json type
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(jsonData3.String())
}
