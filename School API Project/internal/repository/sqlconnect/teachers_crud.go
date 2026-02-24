package sqlconnect

import (
	"database/sql"
	"log"
	"net/http"
	"reflect"
	"restapi/internal/model"
	"strings"
)

func QueryFunc(r *http.Request, query string, args []any) (string, []any) {

	params := map[string]string{
		"first_name": "first_name",
		"last_name":  "last_name",
		"email":      "email",
		"class":      "class",
		"subject":    "subject",
	}

	for param, dbField := range params {
		value := r.URL.Query().Get(param)
		if value != "" {
			query += " AND" + dbField + " =?"
			args = append(args, value)
		}
	}
	return query, args

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

func SortFunc(r *http.Request, query string) string {

	sortParams := r.URL.Query()["sortby"]
	if len(sortParams) > 0 {
		query += " ORDER BY"
		for i, param := range sortParams {
			parts := strings.Split(param, ":")
			if len(parts) > 2 {
				continue
			}
			field, order := parts[0], parts[1]
			if !isValidSortOrder(field) || !isValidSortField(field) {
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

func GetTeachersDbHandler(teachers []model.Teacher, r *http.Request) ([]model.Teacher, error) {

	db, err := ConnectDB()
	if err != nil {
		//http.Error(w, "Error connecting ta database", http.StatusInternalServerError)
		return nil, err
	}
	defer db.Close()

	query := "SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE 1=1"
	var args []any

	// filtering
	query, args = QueryFunc(r, query, args)

	// sorting
	query = SortFunc(r, query)

	rows, err := db.Query(query, args...)
	if err != nil {
		// http.Error(w, "Database query error", http.StatusInternalServerError)
		return nil, err
	}
	// teacherList := make([]model.Teacher, 0)
	for rows.Next() {
		var teacher model.Teacher
		err = rows.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
		if err != nil {
			// http.Error(w, "Error scanning database results", http.StatusInternalServerError)
			return nil, err
		}
		teachers = append(teachers, teacher)
	}
	return teachers, nil
}

func GetTeacherByIdHandler(id int) (model.Teacher, error) {

	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return model.Teacher{}, err
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
			return model.Teacher{}, err
		}
		// http.Error(w, "Database query error", http.StatusInternalServerError)
		return model.Teacher{}, err
	}
	return teacher, nil
}

func AddTeacherDbHandler(newTeachers []model.Teacher) ([]model.Teacher, error) {

	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO teachers (first_name, last_name, email, class, subject) VALUES (?,?,?,?,?)")
	if err != nil {
		// http.Error(w, "Error in preparing SQL query", http.StatusInternalServerError)
		return nil, err
	}
	defer stmt.Close()

	addedTeachers := make([]model.Teacher, len(newTeachers))
	for i, newTeacher := range newTeachers {
		result, err := stmt.Exec(newTeacher.FirstName, newTeacher.LastName, newTeacher.Email, newTeacher.Class, newTeacher.Subject)
		if err != nil {
			// http.Error(w, "Error inserting data into database", http.StatusInternalServerError)
			return nil, err
		}
		lastID, err := result.LastInsertId()
		if err != nil {
			// http.Error(w, "Error getting last insert ID", http.StatusInternalServerError)
			return nil, err
		}
		newTeacher.ID = int(lastID)
		addedTeachers[i] = newTeacher
	}
	return newTeachers, nil
}

func UpdatedTeachersDbHandler(id int, updatedTeacher model.Teacher) (model.Teacher, error) {

	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Eror connecting to database", http.StatusInternalServerError)
		return model.Teacher{}, err
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
			return model.Teacher{}, err
		}
		// http.Error(w, "Unable to retrieve data", http.StatusInternalServerError)
		return model.Teacher{}, err
	}

	updatedTeacher.ID = existingTeacher.ID

	query = "UPDATE teachers SET first_name=?, last_name=?, email=?, class=?, subject=? WHERE id=?"

	_, err = db.Exec(query, updatedTeacher.FirstName, updatedTeacher.LastName, updatedTeacher.Email, updatedTeacher.Class, updatedTeacher.Subject, updatedTeacher.ID)
	if err != nil {
		// http.Error(w, "Error updating teacher", http.StatusInternalServerError)
		return model.Teacher{}, err
	}

	return updatedTeacher, nil
}

func PatchTeachersDbHandler(updatedTeachers []map[string]any) error {

	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Errro connecting to dabase", http.StatusInternalServerError)
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Print(err)
		// http.Error(w, "Error starting transaction", http.StatusInternalServerError)
		return err
	}

	for _, updateTeacher := range updatedTeachers {
		idFloat, ok := updateTeacher["id"].(float64)
		if !ok {
			// http.Error(w, "Invalid teacher ID in update", http.StatusBadRequest)
			return err
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
				return err
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
							return err
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
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		// http.Error(w, "Error committing transaction", http.StatusInternalServerError)
		return err
	}

	return nil
}

func PatchOneTeachersDbHandler(id int, updatedTeacher map[string]any) (model.Teacher, error) {

	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return model.Teacher{}, err
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
			return model.Teacher{}, err
		}
		// http.Error(w, "Unable to retrieve data", http.StatusInternalServerError)
		return model.Teacher{}, err
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
		return model.Teacher{}, err
	}
	return existingTeacher, nil
}

func DeleteOneTeachersDbHandler(id int) error {

	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return err
	}
	defer db.Close()

	query := "DELETE FROM teachers WHERE id=?"
	result, err := db.Exec(query, id)
	if err != nil {
		// http.Error(w, "Error deleting teacher", http.StatusInternalServerError)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// http.Error(w, "Error retrieving delete result", http.StatusInternalServerError)
		return err
	}
	if rowsAffected == 0 {
		// http.Error(w, "Teacher not found", http.StatusNotFound)
		return err
	}

	return nil
}

func DeleteTeachersDbHandler(ids []int) ([]int, error) {

	db, err := ConnectDB()
	if err != nil {
		// http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return nil, err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		// http.Error(w, "Error starting transaction", http.StatusInternalServerError)
		return nil, err
	}

	query := "DELETE FROM teachers WHERE id = ?"
	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		// http.Error(w, "Error preparing delete statement", http.StatusInternalServerError)
		return nil, err
	}
	defer stmt.Close()

	var deletedIDs []int

	for _, id := range ids {
		result, err := stmt.Exec(id)
		if err != nil {
			tx.Rollback()
			// http.Error(w, "Error deleting teacher", http.StatusInternalServerError)
			return nil, err
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			tx.Rollback()
			// http.Error(w, "Error retrieving deleted result", http.StatusInternalServerError)
			return nil, err
		}
		// if teacher was deleted then add the ID to deletedIDs slice
		if rowsAffected > 0 {
			deletedIDs = append(deletedIDs, id)
		}
		if rowsAffected < 1 {
			tx.Rollback()
			// http.Error(w, fmt.Sprintf("ID %d does not exist", id), http.StatusInternalServerError)
			return nil, err
		}
	}
	// commit
	err = tx.Commit()
	if err != nil {
		// http.Error(w, "Error committing transcaction", http.StatusInternalServerError)
		return nil, err
	}
	if len(deletedIDs) < 1 {
		// http.Error(w, "IDs do not exist", http.StatusBadRequest)
		return nil, err
	}
	return deletedIDs, nil

}
