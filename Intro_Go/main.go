package main

import (
	"encoding/json" // To convert data to JSON
	"fmt"           // For printing messages
	"math/rand"
	"net/http" // Go's HTTP package
	"strconv"  // To generate a new ID

	"github.com/gorilla/mux" // Gorilla Mux for routing
)

// Course struct to represent a picnic course
type Course struct {
	ID          string  `json:"id"`          // Unique identifier
	Name        string  `json:"name"`        // Course name
	Price       int     `json:"price"`       // Course price
	Description string  `json:"description"` // Course description
	Author      *Author `json:"author"`      // Pointer to Author struct
}

// Author struct for the course creator
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake database (slice of course)
var courses []Course

// get all courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get All Courses")
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Convert the courses slice to JSON and write it to the response
	json.NewEncoder(w).Encode(courses)
}

// get one course by id
func getOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get One Course")
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Get the ID from the request(URL)
	vars := mux.Vars(r)
	id := vars["id"]

	for _, course := range courses {
		if course.ID == id {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given ID")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Create One Course")
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	var newCourse Course

	// Decode to JSON body into a Course struct
	json.NewDecoder(r.Body).Decode(&newCourse)

	// Basic 'name' validation
	if newCourse.Name == "" {
		json.NewEncoder(w).Encode("Course name is required")
		return
	}

	// Basic 'name' validation
	for _, course := range courses {
		if course.Name == newCourse.Name {
			json.NewEncoder(w).Encode("Name has already taken")
			return
		}
	}

	// Generate a new ID
	newID := strconv.Itoa(rand.Intn(100))
	newCourse.ID = newID

	// Add the new course the slice
	courses = append(courses, newCourse)

	json.NewEncoder(w).Encode(courses)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Update One Course")

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Get the ID from the request(URL)
	vars := mux.Vars(r)

	// Find the course to update
	for i, course := range courses {
		if course.ID == vars["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			var updateCourse Course
			json.NewDecoder(r.Body).Decode(&updateCourse)
			updateCourse.ID = vars["id"]
			courses = append(courses, updateCourse)
			json.NewEncoder(w).Encode(updateCourse)
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found")

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One Course")

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Get the ID from the request(URL)
	vars := mux.Vars(r)
	for i, course := range courses {
		if course.ID == vars["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			break
		}
	}
}

func main() {
	fmt.Println("Building API")

	// Initialize some fake data
	courses = []Course{
		{
			ID:          "2",
			Name:        "Picnic Planning 101",
			Price:       20,
			Description: "Learn to plan the perfect picnic!",
			Author:      &Author{Fullname: "Alice", Website: "alice.com"},
		},
		{
			ID:          "5",
			Name:        "Outdoor Cooking Basics",
			Price:       30,
			Description: "Master cooking outdoors!",
			Author:      &Author{Fullname: "Bob", Website: "bob.com"},
		},
	}

	// Set up the router (more on this next)
	r := mux.NewRouter()
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// Start the server
	http.ListenAndServe(":8080", r)

}
