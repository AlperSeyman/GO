package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"restapi/internal/model"
	"restapi/internal/repository/sqlconnect"
	"restapi/pkg/utils"
	"strconv"
)

type ExecResponse struct {
	Status string       `json:"status"`
	Count  int          `json:"count"`
	Data   []model.Exec `json:"data"`
}

// get method --> /execs
func getAllExecs(w http.ResponseWriter, r *http.Request) {

	var execs []model.Exec
	var err error
	execs, err = sqlconnect.GetExecsDbHandler(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := ExecResponse{
		Status: "success",
		Count:  len(execs),
		Data:   execs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// get method /execs/{exec_id}
func getExecById(w http.ResponseWriter, idStr string) {

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID: ID must be integer", http.StatusBadRequest)
		return
	}
	exec, err := sqlconnect.GetExecByIdHandler(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exec)
}

func GetExecsHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("exec_id")
	if idStr == "" {
		getAllExecs(w, r)
	} else {
		getExecById(w, idStr)
	}
}

// post method --> create a new exec
func AddExecsHandler(w http.ResponseWriter, r *http.Request) {

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var rawExecs []map[string]any
	err = json.Unmarshal(bodyBytes, &rawExecs)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	allowFields := sqlconnect.GetAllowedFields(model.Exec{})
	for _, exec := range rawExecs {
		for key := range exec {
			_, ok := allowFields[key]
			if !ok {
				http.Error(w, "Invalid field in JSON", http.StatusBadRequest)
				return
			}
		}
	}

	var newExecs []model.Exec
	err = json.Unmarshal(bodyBytes, &newExecs)
	if err != nil {
		http.Error(w, "Invalid Request body", http.StatusBadRequest)
		return
	}

	for _, exec := range newExecs {
		val := reflect.ValueOf(exec)
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			if field.Kind() == reflect.String && field.String() == "" {
				http.Error(w, "All fields are required", http.StatusBadRequest)
				return
			}
		}
	}

	addExecs, err := sqlconnect.AddExecDbHandler(newExecs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := ExecResponse{
		Status: "success",
		Count:  len(addExecs),
		Data:   addExecs,
	}

	json.NewEncoder(w).Encode(response)
}

// patch method --> /execs
func PatchExecsHandler(w http.ResponseWriter, r *http.Request) {

	var updatedExec []map[string]any
	err := json.NewDecoder(r.Body).Decode(&updatedExec)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = sqlconnect.PatchExecsDbHandler(updatedExec)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// patch method --> /execs/{exec_id}
func PatchOneExecsHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("exec_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Exec ID", http.StatusBadRequest)
		return
	}

	var updatedExec map[string]any
	err = json.NewDecoder(r.Body).Decode(&updatedExec)
	if err != nil {
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	existingExec, err := sqlconnect.PatchOneExecsDbHandler(id, updatedExec)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingExec)

}

// delete method --> /execs
func DeleteExecsHandler(w http.ResponseWriter, r *http.Request) {

	var ids []int

	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	deletedIDs, err := sqlconnect.DeleteExecsDbHandler(ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Status    string `json:"status"`
		DeleteIDs []int  `json:"deleted_id"`
	}{
		Status:    "Execs successfully deleted",
		DeleteIDs: deletedIDs,
	}
	json.NewEncoder(w).Encode(response)
}

// delete method --> /execs/{exec_id}
func DeleteOneExecsHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("exec_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Exec ID", http.StatusBadRequest)
		return
	}

	err = sqlconnect.DeleteOneExecsDbHandler(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Status string `json:"status"`
		ID     int    `json:"id"`
	}{
		Status: "Exec successfully deleted",
		ID:     id,
	}
	json.NewEncoder(w).Encode(response)
}

func LoginExecsHandler(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	var req model.Exec
	// Data Validation
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// Search for user if user actually exists
	exec, err := sqlconnect.LoginExecsDbHandler(req.Username)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// is user active
	if exec.InactiveStatus {
		http.Error(w, "Account is inactive", http.StatusForbidden)
		return
	}

	// verify password
	match, err := utils.VerifyPassword(req.Password, exec.Password)
	if err != nil {
		http.Error(w, "Error verifying password", http.StatusInternalServerError)
		return
	}
	if !match {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	// generate token

	// send token as a response or as a cookie
}
