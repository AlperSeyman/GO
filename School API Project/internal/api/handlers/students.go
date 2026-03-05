package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"restapi/internal/model"
	"restapi/internal/repository/sqlconnect"
	"strconv"
)

type StudentResponse struct {
	Status string          `json:"status"`
	Count  int             `json:"count"`
	Data   []model.Student `json:"data"`
}

// get method --> /students
func getAllStudents(w http.ResponseWriter, r *http.Request) {

	var students []model.Student
	var err error
	students, err = sqlconnect.GetStudentsDbHandler(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := StudentResponse{
		Status: "success",
		Count:  len(students),
		Data:   students,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

// get method --> students/{students_id}
func getStudentsById(w http.ResponseWriter, idStr string) {

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID: ID must be integer", http.StatusBadRequest)
		return
	}
	student, err := sqlconnect.GetStudentByIdHandler(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func GetStudentsHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("student_id")
	if idStr == "" {
		getAllStudents(w, r)
	} else {
		getStudentsById(w, idStr)
	}
}

// post method --> create a new student
func AddStudentsHandler(w http.ResponseWriter, r *http.Request) {

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error sending request", http.StatusInternalServerError)
		return
	}

	var rawStudents []map[string]any
	err = json.Unmarshal(bodyBytes, &rawStudents)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	allowedFields := sqlconnect.GetAllowedFields(model.Student{})
	for _, student := range rawStudents {
		for key := range student {
			_, ok := allowedFields[key]
			if !ok {
				http.Error(w, "Unacceptable field found. Only use allowed field", http.StatusBadRequest)
				return
			}
		}
	}

	var newStudents []model.Student
	err = json.Unmarshal(bodyBytes, &newStudents)
	if err != nil {
		http.Error(w, "Invalid Request body", http.StatusBadRequest)
		return
	}

	for _, student := range newStudents {
		val := reflect.ValueOf(student)
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			if field.Kind() == reflect.String && field.String() == "" {
				http.Error(w, "All fields are required", http.StatusBadRequest)
				return
			}
		}
	}

	addedStudents, err := sqlconnect.AddStudentDbHandler(newStudents)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := StudentResponse{
		Status: "success",
		Count:  len(addedStudents),
		Data:   addedStudents,
	}

	json.NewEncoder(w).Encode(response)
}

// put method --> /students/{student_id}
func UpdateStudentsHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("student_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Student ID", http.StatusBadRequest)
		return
	}

	var updatedStudent model.Student
	err = json.NewDecoder(r.Body).Decode(&updatedStudent)
	if err != nil {
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	updatedStudentFromDb, err := sqlconnect.UpdateStudentsDbHandler(id, updatedStudent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedStudentFromDb)

}

// patch method --> /students
func PatchStudentsHandler(w http.ResponseWriter, r *http.Request) {

	var updatedStudent []map[string]any
	err := json.NewDecoder(r.Body).Decode(&updatedStudent)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	err = sqlconnect.PatchTeachersDbHandler(updatedStudent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// patch method --> /teachers/{teacher_id}
func PatchOneStudentsHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("student_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Student ID", http.StatusBadRequest)
		return
	}

	var updatedStudent map[string]any
	err = json.NewDecoder(r.Body).Decode(&updatedStudent)
	if err != nil {
		http.Error(w, "Ivalid Request Payload", http.StatusBadRequest)
		return
	}

	existingStudent, err := sqlconnect.PatchOneStudentsDbHandler(id, updatedStudent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingStudent)
}

// delete method --> /students
func DeleteStudentsHandler(w http.ResponseWriter, r *http.Request) {

	var ids []int

	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	deletedIDs, err := sqlconnect.DeleteStudentsDbHandler(ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Status    string `json:"status"`
		DeleteIDs []int  `json:"deleted_id"`
	}{
		Status:    "Students successfully deleted",
		DeleteIDs: deletedIDs,
	}
	json.NewEncoder(w).Encode(response)
}

// delete method --> /students/{student_id}
func DeleteOneStudentsHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("student_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
