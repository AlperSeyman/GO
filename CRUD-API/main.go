package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	for index, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	for _, movie := range movies {
		if movie.ID == id {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var newMovie Movie

	json.NewDecoder(r.Body).Decode(&newMovie)

	// Generate a new ID
	newID := strconv.Itoa(rand.Intn(100))
	newMovie.ID = newID

	movies = append(movies, newMovie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	for index, movie := range movies {
		if movie.ID == id {

			movies = append(movies[:index], movies[index+1:]...)
			var updateMovie Movie
			err := json.NewDecoder(r.Body).Decode(&updateMovie)
			if err != nil {
				log.Fatal(err)
			}

			updateMovie.ID = movie.ID

			if updateMovie.Title == "" {
				updateMovie.Title = movie.Title
			}

			if updateMovie.Isbn == "" {
				updateMovie.Isbn = movie.Isbn
			}

			if updateMovie.Director == nil {
				updateMovie.Director = &Director{}
			}

			if updateMovie.Director.Firstname == "" {
				updateMovie.Director.Firstname = movie.Director.Firstname
			}

			if updateMovie.Director.Lastname == "" {
				updateMovie.Director.Lastname = movie.Director.Lastname
			}

			movies = append(movies, updateMovie)
			json.NewEncoder(w).Encode(updateMovie)
			return

		}
	}
}

func main() {

	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "1234", Title: "Movie One", Director: &Director{Firstname: "First", Lastname: "Director"}})
	movies = append(movies, Movie{ID: "2", Isbn: "5678", Title: "Movie Two", Director: &Director{Firstname: "Second", Lastname: "Director"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PATCH")
	r.HandleFunc("/movies{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
