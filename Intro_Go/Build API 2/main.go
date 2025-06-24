package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

// controllers - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h4>Welcome to Home Page</h4>"))
}

// get all courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "applicaiton/json")
	json.NewEncoder(w).Encode(courses)
}

// get one course
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "applicaiton/json")

	// grab id from request
	vars := mux.Vars(r)

	// loop trough courses, find matching id and return the response
	for _, course := range courses {
		if vars["id"] == course.CourseId {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}

// create one course
func creatOneCourse(w http.ResponseWriter, r *http.Request) {

	var course Course

	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "applicaiton/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what about - {}
	json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// generate unique id, string
	course.CourseId = strconv.Itoa(rand.Intn(100))

	// append new course into courses
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "applicaiton/json")

	// first - grab id from request
	vars := mux.Vars(r)

	// loop, id, remove, add with my ID
	for index, course := range courses {
		if vars["id"] == course.CourseId {
			courses = append(courses[:index], courses[index+1:]...) // remove a value from slices based on index
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = vars["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "applicaiton/json")

	// first - grab id from request
	vars := mux.Vars(r)

	// loop, id, remove
	for index, course := range courses {
		if vars["id"] == course.CourseId {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}

func main() {

	courses = []Course{
		{CourseId: "2", CourseName: "Python", CoursePrice: 100, Author: &Author{Fullname: "nikola", Website: "Udemy"}},
		{CourseId: "5", CourseName: "Java", CoursePrice: 119, Author: &Author{Fullname: "moin", Website: "Online Java"}},
	}

	// routing
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", creatOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
