package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"restapi/internal/model"
	"restapi/internal/repository/sqlconnect"
	"strconv"
	"strings"
)

type Response struct {
	Status string          `json:"status"`
	Count  int             `json:"count"`
	Data   []model.Teacher `json:"data"`
}

// get methed --> /teachers/
func getAllTeachers(w http.ResponseWriter, r *http.Request) {

	db, err := sqlconnect.ConnectDB()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Queries
	// firstName := r.URL.Query().Get("first-name")
	// lastName := r.URL.Query().Get("last-name")

	query := "SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE 1=1"
	var args []any

	// filtering
	query, args = queryFunc(r, query, args)

	// sorting
	query = sortFunc(r, query)

	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Database Query Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	teacherList := make([]model.Teacher, 0)
	for rows.Next() {
		var teacher model.Teacher
		err = rows.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
		if err != nil {
			http.Error(w, "Error scanning database results", http.StatusInternalServerError)
			return
		}
		teacherList = append(teacherList, teacher)

	}

	response := Response{
		Status: "success",
		Count:  len(teacherList),
		Data:   teacherList,
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

// get method --> /teachers/{teacher_id}
func getTeacherById(w http.ResponseWriter, r *http.Request, idStr string) {

	db, err := sqlconnect.ConnectDB()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if idStr == "" {
		getAllTeachers(w, r)
	} else {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID: ID must be a number", http.StatusBadRequest)
			return
		}

		var teacher model.Teacher
		query := "SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?"
		err = db.QueryRow(query, id).Scan(
			&teacher.ID,
			&teacher.FirstName,
			&teacher.LastName,
			&teacher.Email,
			&teacher.Class,
			&teacher.Subject,
		)
		if err == sql.ErrNoRows {
			http.Error(w, "Teacher not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, "Database query error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
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

// post method --> create a new teacher
func addTeachersHandler(w http.ResponseWriter, r *http.Request) {

	db, err := sqlconnect.ConnectDB()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var newTeachers []model.Teacher
	err = json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("INSERT INTO teachers (first_name, last_name, email, class, subject) VALUES (?,?,?,?,?)")
	if err != nil {
		http.Error(w, "Error in preparing SQL query", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	addedTeachers := make([]model.Teacher, len(newTeachers))
	for i, newTeacher := range newTeachers {
		result, err := stmt.Exec(newTeacher.FirstName, newTeacher.LastName, newTeacher.Email, newTeacher.Class, newTeacher.Subject)
		if err != nil {
			http.Error(w, "Error inserting data into database", http.StatusInternalServerError)
			return
		}
		lastID, err := result.LastInsertId()
		if err != nil {
			http.Error(w, "Error getting last insert ID", http.StatusInternalServerError)
			return
		}
		newTeacher.ID = int(lastID)
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

// put method --> /teachers/{teacher_id}
func updateTeachersHandler(w http.ResponseWriter, r *http.Request) {

	db, err := sqlconnect.ConnectDB()
	if err != nil {
		http.Error(w, "Eror connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	idStr := strings.TrimPrefix(r.URL.Path, "/teachers/")
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

	var existingTeacher model.Teacher
	query := "SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?"
	err = db.QueryRow(query, id).Scan(
		&existingTeacher.ID,
		&existingTeacher.FirstName,
		&existingTeacher.LastName,
		&existingTeacher.Email,
		&existingTeacher.Class,
		&existingTeacher.Subject,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Teacher not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Unable to retrieve data", http.StatusInternalServerError)
		return
	}

	updatedTeacher.ID = existingTeacher.ID

	query = "UPDATE teachers SET first_name=?, last_name=?, email=?, class=?, subject=? WHERE id=?"

	_, err = db.Exec(query, updatedTeacher.FirstName, updatedTeacher.LastName, updatedTeacher.Email, updatedTeacher.Class, updatedTeacher.Subject, updatedTeacher.ID)
	if err != nil {
		http.Error(w, "Error updating teacher", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTeacher)
}

// patch method --> /teachers/{teacher_id}
func patchUpdateTeachersHandler(w http.ResponseWriter, r *http.Request) {

	db, err := sqlconnect.ConnectDB()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	idStr := strings.TrimPrefix(r.URL.Path, "/teachers/")
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

	var existingTeacher model.Teacher
	query := "SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?"
	err = db.QueryRow(query, id).Scan(
		&existingTeacher.ID,
		&existingTeacher.FirstName,
		&existingTeacher.LastName,
		&existingTeacher.Email,
		&existingTeacher.Class,
		&existingTeacher.Subject,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Teacher not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Unable to retrieve data", http.StatusInternalServerError)
		return
	}

	// apply update
	// for field, value := range updatedTeacher {

	// 	strValue, ok := value.(string)
	// 	if !ok {
	// 		http.Error(w, "Invalid data type: expected text", http.StatusBadRequest)
	// 		return
	// 	}

	// 	switch field {
	// 	case "first_name":
	// 		existingTeacher.FirstName = strValue
	// 	case "last_name":
	// 		existingTeacher.LastName = strValue
	// 	case "email":
	// 		existingTeacher.Email = strValue
	// 	case "class":
	// 		existingTeacher.Class = strValue
	// 	case "subject":
	// 		existingTeacher.Subject = strValue
	// 	}
	// }

	// apply updates using reflect
	teacherValue := reflect.ValueOf(&existingTeacher).Elem() // --> {100 Alice Brown alice@example.com 6A World History}
	teacherType := teacherValue.Type()                       // --> model.Teacher
	// teacherValue.Field(0)  --> 100
	// teacherType.Field(0))  --> {ID  int json:"id,omitempty" 0 [0] false}
	// teacherType.NumField()  --> 6
	// field.Tag.Get("json") --> teacher's struct  example: "first_name,omitempty"
	// teacherField := teacherValue.Field(i) --> alice@example.com (old)
	// teacherField := teacherValue.Field(i).Type() --> string
	// reflect.ValueOf(value) --> aliceSmith@example.com (new)

	for column, value := range updatedTeacher {
		for i := 0; i < teacherType.NumField(); i++ {
			field := teacherType.Field(i)
			if field.Tag.Get("json") == column+",omitempty" {
				if teacherValue.Field(i).CanSet() { // checking that user is allowed to update
					teacherField := teacherValue.Field(i)
					teacherField.Set(reflect.ValueOf(value).Convert(teacherValue.Field(i).Type()))

				}
			}
		}
	}

	query = "UPDATE teachers SET first_name=?, last_name=?, email=?, class=?, subject=? WHERE id=?"
	_, err = db.Exec(query,
		existingTeacher.FirstName,
		existingTeacher.LastName,
		existingTeacher.Email,
		existingTeacher.Class,
		existingTeacher.Subject,
		existingTeacher.ID,
	)
	if err != nil {
		http.Error(w, "Error updating teacher", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingTeacher)
}

// teacher handler
func TeachersHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Main Page")

	switch r.Method {
	case http.MethodGet:
		// call get method handler function
		getTeachersHandler(w, r)
	case http.MethodPost:
		// call post method handler function
		addTeachersHandler(w, r)
	case http.MethodPut:
		// call put method handler function
		updateTeachersHandler(w, r)
	case http.MethodPatch:
		// call patch method handler function
		patchUpdateTeachersHandler(w, r)
	case http.MethodDelete:
		w.Write([]byte("DELETE Method on Teachers Route Page"))
	default:
		w.Write([]byte("Teachers Route Page"))
	}
}

// Helper Functions
func queryFunc(r *http.Request, query string, args []any) (string, []any) { // add filters

	params := map[string]string{
		"first_name": "first_name",
		"last_name":  "last_name",
		"emai":       "email",
		"class":      "class",
		"subject":    "subject",
	}

	for param, dbFied := range params {
		value := r.URL.Query().Get(param)
		if value != "" {
			query += " AND " + dbFied + " = ?"
			args = append(args, value)
		}
	}
	return query, args

	// params := map[string]string{
	// 	"first_name": "first_name",
	// 	"last_name":  "last_name",
	// 	"email":      "email",
	// 	"class":      "class",
	// 	"subject":    "subject",
	// }

	// for param, dbField := range params {
	// 	value := r.URL.Query().Get(param)
	// 	if value != "" {
	// 		query += " AND " + dbField + " = ?"
	// 		args = append(args, value)
	// 	}
	// }

	// if firstName != "" {
	// 	query += " AND first_name=?"
	// 	args = append(args, firstName)
	// }
	// if lastName != "" {
	// 	query += " AND last_name=?"
	// 	args = append(args, lastName)
	// }
}

func isValidSortOrder(order string) bool {
	return order == "asc" || order == "desc"
}

func isValidSortField(field string) bool {
	validField := map[string]bool{
		"first_name": true,
		"last_name":  true,
		"email":      true,
		"class":      true,
		"subject":    true,
	}

	return validField[field]
}

func sortFunc(r *http.Request, query string) string { // add sorting
	sortParams := r.URL.Query()["sortby"]
	if len(sortParams) > 0 {
		query += " ORDER BY"
		for i, param := range sortParams {
			parts := strings.Split(param, ":")
			if len(parts) > 2 {
				continue
			}
			field, order := parts[0], parts[1]
			if !isValidSortOrder(field) || !isValidSortField(order) {
				continue
			}
			if i > 0 {
				query += " ,"
			}
			query += " " + field + " " + order
		}
	}
	return query
}
