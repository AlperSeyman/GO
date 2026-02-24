package handlers

import (
	"encoding/json"
	"net/http"
	"restapi/internal/model"
	"restapi/internal/repository/sqlconnect"
	"strconv"
)

type Response struct {
	Status string          `json:"status"`
	Count  int             `json:"count"`
	Data   []model.Teacher `json:"data"`
}

// get method --> /teachers
func getAllTeachers(w http.ResponseWriter, r *http.Request) {

	var teachers []model.Teacher
	teachers, err := sqlconnect.GetTeachersDbHandler(teachers, r)
	if err != nil {
		return
	}
	response := Response{
		Status: "success",
		Count:  len(teachers),
		Data:   teachers,
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

// get method --> /teachers/{teacher_id}
func getTeacherById(w http.ResponseWriter, r *http.Request, idStr string) {

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID: ID must be intger", http.StatusBadRequest)
		return
	}
	teacher, err := sqlconnect.GetTeacherByIdHandler(id)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teacher)
}

func GetTeachersHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("teacher_id")
	if idStr == "" {
		getAllTeachers(w, r)
	} else {
		getTeacherById(w, r, idStr)
	}
}

// post method --> create a new teacher
func AddTeachersHandler(w http.ResponseWriter, r *http.Request) {

	var newTeachers []model.Teacher
	err := json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	addedTeachers, err := sqlconnect.AddTeacherDbHandler(newTeachers)
	if err != nil {
		return
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

// put method --> /teachers/{teacher_id}
func UpdateTeachersHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("teacher_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Teacher ID", http.StatusBadRequest)
		return
	}

	var updatedTeacher model.Teacher
	err = json.NewDecoder(r.Body).Decode(&updatedTeacher)
	if err != nil {
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	updatedTeacherFromDb, err := sqlconnect.UpdatedTeachersDbHandler(id, updatedTeacher)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTeacherFromDb)
}

// patch method --> /teachers
func PatchTeachersHandler(w http.ResponseWriter, r *http.Request) {

	var updatedTeachers []map[string]any
	err := json.NewDecoder(r.Body).Decode(&updatedTeachers)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = sqlconnect.PatchTeachersDbHandler(updatedTeachers)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// patch method --> /teachers/{teacher_id}
func PatchOneTeachersHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("teacher_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Teacher ID", http.StatusBadRequest)
		return
	}

	var updatedTeacher map[string]any
	err = json.NewDecoder(r.Body).Decode(&updatedTeacher)
	if err != nil {
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	existingTeacher, err := sqlconnect.PatchOneTeachersDbHandler(id, updatedTeacher)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingTeacher)
}

// delete method --> /teachers
func DeleteTeachersHandler(w http.ResponseWriter, r *http.Request) {

	var ids []int

	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	deletedIDs, err := sqlconnect.DeleteTeachersDbHandler(ids)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Status     string `json:"status"`
		DeletedIDs []int  `json:"deleted_id"`
	}{
		Status:     "Teachers successfully deleted",
		DeletedIDs: deletedIDs,
	}
	json.NewEncoder(w).Encode(response)
}

// delete method --> /teachers/{teacher_id}
func DeleteOneTeachersHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("teacher_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Teacher ID", http.StatusBadRequest)
		return
	}

	err = sqlconnect.DeleteOneTeachersDbHandler(id)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Status string `json:"status"`
		ID     int    `json:"id"`
	}{
		Status: "Teacher successfully deleted",
		ID:     id,
	}
	json.NewEncoder(w).Encode(response)
}
