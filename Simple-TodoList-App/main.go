package main

import (
	"fmt"
	"net/http"
)

var goLang = "Study Go Lang"
var german = "Learn German"
var sport = "Play Basketball"
var tasks = []string{goLang, german, sport}

func main() {

	// Basic Web API

	http.HandleFunc("/", home_page)
	http.HandleFunc("/show-tasks", show_tasks)

	http.ListenAndServe(":8080", nil)

}

func home_page(writer http.ResponseWriter, request *http.Request) {
	greeting := "####### Welcome to Todolist App! #######"
	fmt.Fprintln(writer, greeting)
}

func show_tasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range tasks {
		fmt.Fprintln(writer, task)
	}

}
