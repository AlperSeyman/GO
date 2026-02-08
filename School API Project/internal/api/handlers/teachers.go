package handlers

import (
	"encoding/json"
	"net/http"
	"restapi/internal/model"
	"strconv"
	"strings"
	"sync"
)

type Response struct {
	Status string          `json:"status"`
	Count  int             `json:"count"`
	Data   []model.Teacher `json:"data"`
}

var (
	teachers = make(map[int]model.Teacher)
	mutex    = &sync.Mutex{} // Prevents data corruption by locking shared resources. Using for 'post' method.
	nextID   = 1
)

func init() {

	teachers[nextID] = model.Teacher{
		ID:        nextID,
		FirstName: "John",
		LastName:  "Doe",
		Class:     "9A",
		Subject:   "Math",
	}
	nextID++
	teachers[nextID] = model.Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Smith",
		Class:     "10A",
		Subject:   "Algebra",
	}
	nextID++
	teachers[nextID] = model.Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Doe",
		Class:     "12C",
		Subject:   "Biology",
	}
	nextID++
}

func getAllTeachers(w http.ResponseWriter, r *http.Request) {

	// Queries
	firstName := r.URL.Query().Get("first-name")
	lastName := r.URL.Query().Get("last-name")

	teacherList := make([]model.Teacher, 0, len(teachers))
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

	var newTeachers []model.Teacher

	err := json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	addedTeachers := make([]model.Teacher, len(newTeachers))
	for i, newTeacher := range newTeachers {
		newTeacher.ID = nextID
		teachers[nextID] = newTeacher
		addedTeachers[i] = newTeacher
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

func TeachersHandler(w http.ResponseWriter, r *http.Request) {
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
