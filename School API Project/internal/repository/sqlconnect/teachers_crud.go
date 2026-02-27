package sqlconnect

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"restapi/internal/model"
	"restapi/pkg/utils"
	"strings"
)

func getAllowedColumns() map[string]bool {

	allowed := make(map[string]bool)
	teacherType := reflect.TypeOf(model.Teacher{})

	for i := 0; i < teacherType.NumField(); i++ {
		dbTag := teacherType.Field(i).Tag.Get("db")
		if dbTag != "" && dbTag != "-" {
			columnName := strings.Split(dbTag, ",")[0]
			allowed[columnName] = true
		}
	}
	return allowed
}

func queryFunc(r *http.Request, query string, args []any) (string, []any) {

	allowedColumns := getAllowedColumns()

	for param, values := range r.URL.Query() {
		if allowedColumns[param] {
			value := values[0]
			if value != "" {
				query += " AND" + param + " = ?"
				args = append(args, value)
			}
		}
	}
	return query, args
}

func isValidSortOrder(order string) bool {
	cleanOrder := strings.ToLower(strings.TrimSpace(order))
	return cleanOrder == "asc" || order == "desc"
}

func isValidSortField(field string) bool {

	allowedColumnd := getAllowedColumns()
	cleanField := strings.ToLower(strings.TrimSpace(field))
	return allowedColumnd[cleanField]

}

func sortFunc(r *http.Request, query string) string {

	sortParams := r.URL.Query()["sortby"]
	if len(sortParams) == 0 {
		return query
	}

	var validSorts []string
	for _, param := range sortParams {
		parts := strings.Split(param, ":")
		if len(parts) != 2 {
			continue
		}
		field := parts[0]
		order := parts[1]

		if isValidSortField(field) && isValidSortOrder(order) {
			validSorts = append(validSorts, field+" "+strings.ToLower(order))
		}
	}
	if len(validSorts) > 0 {
		query += " ORDER BY" + strings.Join(validSorts, ", ")
	}
	return query

}

func generateInsertQuery(model any) string {

	modelType := reflect.TypeOf(model)

	if modelType.Kind() == reflect.Slice {
		modelType = modelType.Elem()
	}
	var columns, placeholders string
	for i := 0; i < modelType.NumField(); i++ {
		dbTag := modelType.Field(i).Tag.Get("db")
		fmt.Println("DB Tag", dbTag)
		dbTag = strings.TrimSuffix(dbTag, ",omitempty")
		if dbTag != "" && dbTag != "id" { // skip the ID field if it's auto increment
			if columns != "" {
				columns += ", "
				placeholders += ", "
			}
			columns += dbTag
			placeholders += "?"
		}
	}
	return fmt.Sprintf("INSERT INTO teachers (%s) VALUES (%s)", columns, placeholders)

}

func getStructValues(model any) []any {

	modelValue := reflect.ValueOf(model)
	modelType := modelValue.Type()
	values := []interface{}{}
	for i := 0; i < modelType.NumField(); i++ {
		dbTag := modelType.Field(i).Tag.Get("db")
		if dbTag != "" && dbTag != "id,omitempty" {
			values = append(values, modelValue.Field(i).Interface())
		}
	}
	log.Println(values...)
	return values
}

func GetTeachersDbHandler(teachers []model.Teacher, r *http.Request) ([]model.Teacher, error) {

	db, err := ConnectDB()
	if err != nil {
		//http.Error(w, "Error connecting ta database", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Error connecting ta database")
	}
	defer db.Close()

	query := "SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE 1=1"
	var args []any

	// filtering
	query, args = queryFunc(r, query, args)

	// sorting
	query = sortFunc(r, query)

	rows, err := db.Query(query, args...)
	if err != nil {
		// http.Error(w, "Database query error", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Database query error")
	}
	// teacherList := make([]model.Teacher, 0)
	for rows.Next() {
		var teacher model.Teacher
		err = rows.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
		if err != nil {
			// http.Error(w, "Error scanning database results", http.StatusInternalServerError)
			return nil, utils.ErrorHandler(err, "Error scanning database results")
		}
		teachers = append(teachers, teacher)
	}
	return teachers, nil
}

func GetTeacherByIdHandler(id int) (model.Teacher, error) {

	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return model.Teacher{}, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

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
	if err != nil {
		if err == sql.ErrNoRows {
			// http.Error(w, "Teacher not found", http.StatusNotFound)
			return model.Teacher{}, utils.ErrorHandler(err, "Teacher not found")
		}
		// http.Error(w, "Database query error", http.StatusInternalServerError)
		return model.Teacher{}, utils.ErrorHandler(err, "Database query error")
	}
	return teacher, nil
}

func AddTeacherDbHandler(newTeachers []model.Teacher) ([]model.Teacher, error) {

	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := generateInsertQuery(model.Teacher{})

	stmt, err := db.Prepare(query)
	if err != nil {
		// http.Error(w, "Error in preparing SQL query", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Error in preparing SQL query")
	}
	defer stmt.Close()

	addedTeachers := make([]model.Teacher, len(newTeachers))
	for i, newTeacher := range newTeachers {
		structValues := getStructValues(newTeacher)
		result, err := stmt.Exec(structValues...)
		if err != nil {
			// http.Error(w, "Error inserting data into database", http.StatusInternalServerError)
			return nil, utils.ErrorHandler(err, "Error inserting data into database")
		}
		lastID, err := result.LastInsertId()
		if err != nil {
			// http.Error(w, "Error getting last insert ID", http.StatusInternalServerError)
			return nil, utils.ErrorHandler(err, "Error getting last insert ID")
		}
		newTeacher.ID = int(lastID)
		addedTeachers[i] = newTeacher
	}
	return addedTeachers, nil
}

func UpdatedTeachersDbHandler(id int, updatedTeacher model.Teacher) (model.Teacher, error) {

	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Eror connecting to database", http.StatusInternalServerError)
		return model.Teacher{}, utils.ErrorHandler(err, "Eror connecting to database")
	}
	defer db.Close()

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
			// http.Error(w, "Teacher not found", http.StatusNotFound)
			return model.Teacher{}, utils.ErrorHandler(err, "Teacher not found")
		}
		// http.Error(w, "Unable to retrieve data", http.StatusInternalServerError)
		return model.Teacher{}, utils.ErrorHandler(err, "Unable to retrieve data")
	}

	updatedTeacher.ID = existingTeacher.ID

	query = "UPDATE teachers SET first_name=?, last_name=?, email=?, class=?, subject=? WHERE id=?"

	_, err = db.Exec(query, updatedTeacher.FirstName, updatedTeacher.LastName, updatedTeacher.Email, updatedTeacher.Class, updatedTeacher.Subject, updatedTeacher.ID)
	if err != nil {
		// http.Error(w, "Error updating teacher", http.StatusInternalServerError)
		return model.Teacher{}, utils.ErrorHandler(err, "Error updating teacher")
	}

	return updatedTeacher, nil
}

func PatchTeachersDbHandler(updatedTeachers []map[string]any) error {

	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Errro connecting to dabase", http.StatusInternalServerError)
		return utils.ErrorHandler(err, "Errro connecting to dabase")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Print(err)
		// http.Error(w, "Error starting transaction", http.StatusInternalServerError)
		return utils.ErrorHandler(err, "Error starting transaction")
	}

	for _, updateTeacher := range updatedTeachers {
		idFloat, ok := updateTeacher["id"].(float64)
		if !ok {
			// http.Error(w, "Invalid teacher ID in update", http.StatusBadRequest)
			return utils.ErrorHandler(err, "Invalid teacher ID in update")
		}
		id := int(idFloat)

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
			tx.Rollback()
			if err == sql.ErrNoRows {
				// http.Error(w, "Teacher not found", http.StatusNotFound)
				return utils.ErrorHandler(err, "Teacher not found")
			}
			// http.Error(w, "Unable to retrieve data", http.StatusInternalServerError)
			return nil
		}

		// appyling update using reflect

		teacherValue := reflect.ValueOf(&existingTeacher).Elem()
		teacherType := teacherValue.Type()

		for column, value := range updateTeacher {
			if column == "id" {
				continue // skip updating the id field.
			}
			for i := 0; i < teacherValue.NumField(); i++ {
				field := teacherType.Field(i)
				if field.Tag.Get("json") == column+",omitempty" {
					fieldValue := teacherValue.Field(i) //  old value
					if fieldValue.CanSet() {
						val := reflect.ValueOf(value) // new value
						if val.Type().ConvertibleTo(fieldValue.Type()) {
							fieldValue.Set(val.Convert(fieldValue.Type()))
						} else {
							tx.Rollback()
							log.Printf("Cannot convert %v to %v", val.Type(), fieldValue.Type())
							return utils.ErrorHandler(err, "Error updating data")
						}
					}
					break
				}
			}
		}

		query = "UPDATE teachers SET first_name=?, last_name=?, email=?, class=?, subject=? WHERE id = ?"
		_, err = tx.Exec(query,
			existingTeacher.FirstName,
			existingTeacher.LastName,
			existingTeacher.Email,
			existingTeacher.Class,
			existingTeacher.Subject,
			id,
		)
		if err != nil {
			tx.Rollback()
			// http.Error(w, "Error updating teacher", http.StatusInternalServerError)
			return utils.ErrorHandler(err, "Error updating teacher")
		}
	}
	err = tx.Commit()
	if err != nil {
		// http.Error(w, "Error committing transaction", http.StatusInternalServerError)
		return utils.ErrorHandler(err, "Error committing transaction")
	}

	return nil
}

func PatchOneTeachersDbHandler(id int, updatedTeacher map[string]any) (model.Teacher, error) {

	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return model.Teacher{}, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

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
			// http.Error(w, "Teacher not found", http.StatusNotFound)
			return model.Teacher{}, utils.ErrorHandler(err, "Teacher not found")
		}
		// http.Error(w, "Unable to retrieve data", http.StatusInternalServerError)
		return model.Teacher{}, utils.ErrorHandler(err, "Unable to retrieve data")
	}

	// apply updates using reflect
	teacherValue := reflect.ValueOf(&existingTeacher).Elem()
	teacherType := teacherValue.Type()
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
		// http.Error(w, "Error updating teacher", http.StatusInternalServerError)
		return model.Teacher{}, utils.ErrorHandler(err, "Error updating teacher")
	}
	return existingTeacher, nil
}

func DeleteOneTeachersDbHandler(id int) error {

	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	query := "DELETE FROM teachers WHERE id=?"
	result, err := db.Exec(query, id)
	if err != nil {
		// http.Error(w, "Error deleting teacher", http.StatusInternalServerError)
		return utils.ErrorHandler(err, "Error deleting teacher")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// http.Error(w, "Error retrieving delete result", http.StatusInternalServerError)
		return utils.ErrorHandler(err, "Error retrieving delete result")
	}
	if rowsAffected == 0 {
		// http.Error(w, "Teacher not found", http.StatusNotFound)
		return utils.ErrorHandler(err, "Teacher not found")
	}

	return nil
}

func DeleteTeachersDbHandler(ids []int) ([]int, error) {

	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		// http.Error(w, "Error starting transaction", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Error starting transaction")
	}

	query := "DELETE FROM teachers WHERE id = ?"
	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		// http.Error(w, "Error preparing delete statement", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Error preparing delete statement")
	}
	defer stmt.Close()

	var deletedIDs []int

	for _, id := range ids {
		result, err := stmt.Exec(id)
		if err != nil {
			tx.Rollback()
			// http.Error(w, "Error deleting teacher", http.StatusInternalServerError)
			return nil, utils.ErrorHandler(err, "Error deleting teacher")
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			tx.Rollback()
			// http.Error(w, "Error retrieving deleted result", http.StatusInternalServerError)
			return nil, utils.ErrorHandler(err, "Error retrieving deleted result")
		}
		// if teacher was deleted then add the ID to deletedIDs slice
		if rowsAffected > 0 {
			deletedIDs = append(deletedIDs, id)
		}
		if rowsAffected < 1 {
			tx.Rollback()
			// http.Error(w, fmt.Sprintf("ID %d does not exist", id), http.StatusInternalServerError)
			return nil, utils.ErrorHandler(err, fmt.Sprintf("ID %d does not exist", id))
		}
	}
	// commit
	err = tx.Commit()
	if err != nil {
		// http.Error(w, "Error committing transcaction", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Error committing transcaction")
	}
	if len(deletedIDs) < 1 {
		// http.Error(w, "IDs do not exist", http.StatusBadRequest)
		return nil, utils.ErrorHandler(err, "IDs do not exist")
	}
	return deletedIDs, nil

}
