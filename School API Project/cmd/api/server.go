package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Teacher struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Class     string `json:"class,omitempty"`
	Subject   string `json:"subject,omitempty"`
}

type Response struct {
	Status string    `json:"status"`
	Count  int       `json:"count"`
	Data   []Teacher `json:"data"`
}

var (
	teachers = make(map[int]Teacher)
	mutex    = &sync.Mutex{} // Prevents data corruption by locking shared resources. Using for 'post' method.
	nextID   = 1
)

func init() {

	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "John",
		LastName:  "Doe",
		Class:     "9A",
		Subject:   "Math",
	}
	nextID++
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Smith",
		Class:     "10A",
		Subject:   "Algebra",
	}
	nextID++
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Doe",
		Class:     "12C",
		Subject:   "Biology",
	}
	nextID++
}

func getAllTeachers(w http.ResponseWriter, r *http.Request) {

	// Queris
	firstName := r.URL.Query().Get("first-name")
	lastName := r.URL.Query().Get("last-name")

	teacherList := make([]Teacher, 0, len(teachers))
	for _, teacher := range teachers {
		if (firstName == "" || teacher.FirstName == firstName) && (lastName == "" || teacher.LastName == lastName) {
			teacherList = append(teacherList, teacher)
		}
	}

	response := Response{
		Status: "success",
		Count:  len(teacherList),
		Data:   teacherList,
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func getTeacherById(w http.ResponseWriter, r *http.Request, idStr string) {

	if idStr == "" {
		getAllTeachers(w, r)
	} else {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID: ID must be a number", http.StatusBadRequest)
			return
		}
		teacher, exists := teachers[id]
		if !exists {
			http.Error(w, "Teacher not found", http.StatusFound)
			return
		}
		json.NewEncoder(w).Encode(teacher)
	}
}

func getTeachersHandler(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	if path == "" {
		getAllTeachers(w, r)
	} else {
		getTeacherById(w, r, path)
	}
}

func addTeacherHandler(w http.ResponseWriter, r *http.Request) {

	mutex.Lock()
	defer mutex.Unlock()

	var newTeachers []Teacher

	err := json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	addedTeachers := make([]Teacher, len(newTeachers))

	for i, newTeacher := range newTeachers {
		newTeacher.ID = nextID
		teachers[nextID] = newTeacher
		addedTeachers[i] = newTeacher
		nextID++
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := Response{
		Status: "success",
		Count:  len(addedTeachers),
		Data:   addedTeachers,
	}

	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Main Page")
	w.Write([]byte("Main Page"))
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Main Page")

	switch r.Method {
	case http.MethodGet:
		// call get method handler function
		getTeachersHandler(w, r)
	case http.MethodPost:
		// call post method handler function
		addTeacherHandler(w, r)
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
		Addr:    port,
		Handler: mux,
		// Handler: middlewares.SecurityHeaders(middlewares.CORSMiddleware(mux)),
		// Handler:   middlewares.CORSMiddleware(mux),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port:", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}
