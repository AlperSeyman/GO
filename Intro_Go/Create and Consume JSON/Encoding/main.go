package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	url := "https://jsonplaceholder.typicode.com/todos/1"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {

		todoItem := Todo{}

		json.NewDecoder(response.Body).Decode(&todoItem)

		// Convert to back to JSON
		//todo, err := json.Marshal(todoItem)
		todo, err := json.MarshalIndent(todoItem, "", "\t")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(todo))
	}
}
