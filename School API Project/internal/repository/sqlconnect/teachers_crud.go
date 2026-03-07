package sqlconnect

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"restapi/internal/model"
	"strings"

	"restapi/pkg/utils"
)

func GetTeachersDbHandler(r *http.Request) ([]model.Teacher, error) {

	db, err := ConnectDB()
	if err != nil {
		//http.Error(w, "Error connecting ta database", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Error connecting to database..")
	}
	defer db.Close()

	var teachers []model.Teacher

	query := GenerateSelectQuery(model.Teacher{}, "teachers") + " WHERE 1 = 1"
	var args []any

	// filtering
	query, args = QueryFunc(r, model.Teacher{}, query, args)

	// sorting
	query = SortFunc(r, model.Teacher{}, query)

	rows, err := db.Query(query, args...)
	if err != nil {
		// http.Error(w, "Database query error", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Database query error")
	}
	// teacherList := make([]model.Teacher, 0)
	for rows.Next() {
		var teacher model.Teacher
		structPointers := GetStructPointers(&teacher)
		err = rows.Scan(structPointers...)
		if err != nil {
			// http.Error(w, "Error scanning database results", http.StatusInternalServerError)
			return nil, utils.ErrorHandler(err, "Error scanning database results")
		}
		teachers = append(teachers, teacher)
	}

	err = rows.Err()
	if err != nil {
		return nil, utils.ErrorHandler(err, "The database connection was lost while reading rows")
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

	query := GenerateSelectQuery(teacher, "teachers") + " WHERE  id = ?"

	structPointers := GetStructPointers(&teacher)
	err = db.QueryRow(query, id).Scan(structPointers...)
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

	query := GenerateInsertQuery(model.Teacher{}, "teachers")

	stmt, err := db.Prepare(query)
	if err != nil {
		// http.Error(w, "Error in preparing SQL query", http.StatusInternalServerError)
		return nil, utils.ErrorHandler(err, "Error in preparing SQL query")
	}
	defer stmt.Close()

	addedTeachers := make([]model.Teacher, len(newTeachers))
	for i, newTeacher := range newTeachers {
		structValues := GetStructValues(newTeacher)
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

	query := GenerateSelectQuery(model.Teacher{}, "teachers") + " WHERE id = ?"

	structPointers := GetStructPointers(&existingTeacher)

	err = db.QueryRow(query, id).Scan(structPointers...)
	if err != nil {
		if err == sql.ErrNoRows {
			// http.Error(w, "Teacher not found", http.StatusNotFound)
			return model.Teacher{}, utils.ErrorHandler(err, "Teacher not found")
		}
		// http.Error(w, "Unable to retrieve data", http.StatusInternalServerError)
		return model.Teacher{}, utils.ErrorHandler(err, "Unable to retrieve data")
	}

	updatedTeacher.ID = existingTeacher.ID

	query = GenerateUpdateQuery(model.Teacher{}, "teachers")

	structValues := GetStructValues(updatedTeacher)
	structValues = append(structValues, id)
	_, err = db.Exec(query, structValues...)
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
		return utils.ErrorHandler(err, "Errro connecting to database")
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
		query := GenerateSelectQuery(model.Teacher{}, "teachers") + " WHERE id = ?"

		structPointers := GetStructPointers(&existingTeacher)

		err = tx.QueryRow(query, id).Scan(structPointers...)
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
				cleanTag := strings.Split(field.Tag.Get("json"), ",")[0]
				if cleanTag == column {
					fieldToUpdate := teacherValue.Field(i) // old value
					if fieldToUpdate.CanSet() {
						val := reflect.ValueOf(value) // new value
						if val.Type().ConvertibleTo(fieldToUpdate.Type()) {
							fieldToUpdate.Set(val.Convert(fieldToUpdate.Type()))
						} else {
							tx.Rollback()
							log.Printf("Cannot convert %v to %v", val.Type(), fieldToUpdate.Type())
							return utils.ErrorHandler(err, "Error updating data")
						}
					}
				}
			}
		}

		query = GenerateUpdateQuery(model.Teacher{}, "teachers")

		structValues := GetStructValues(existingTeacher)
		structValues = append(structValues, id)
		_, err = tx.Exec(query, structValues...)
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
	query := GenerateSelectQuery(model.Teacher{}, "teachers") + " WHERE id = ?"

	structPointers := GetStructPointers(&existingTeacher)

	err = db.QueryRow(query, id).Scan(structPointers...)

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
			cleanTag := strings.Split(field.Tag.Get("json"), ",")[0]
			if cleanTag == column {
				if teacherValue.Field(i).CanSet() {
					teacherField := teacherValue.Field(i)
					teacherField.Set(reflect.ValueOf(value).Convert(teacherValue.Field(i).Type()))
				}
			}
		}
	}

	query = GenerateUpdateQuery(model.Teacher{}, "teachers")
	structValues := GetStructValues(existingTeacher)
	structValues = append(structValues, id)
	_, err = db.Exec(query, structValues...)
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

func GetStudentsByTeacher(id int) ([]model.Student, error) {

	db, err := ConnectDB()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	var students []model.Student

	query := GenerateSelectQueryForTeacher(model.Student{}, "students", "teachers")
	row, err := db.Query(query, id)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Database query error")
	}
	defer row.Close()

	for row.Next() {
		var student model.Student
		structPointers := GetStructPointers(&student)
		err = row.Scan(structPointers...)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error scanning database results")
		}
		students = append(students, student)
	}
	err = row.Err()
	if err != nil {
		return nil, utils.ErrorHandler(err, "The database connection was lost while reading rows")
	}
	return students, nil
}

func GetStudentCountByTeacher(id int) (int, error) {

	db, err := ConnectDB()
	if err != nil {
		return 0, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	var studentCount int
	query := "SELECT COUNT(*) FROM students WHERE class = (SELECT class FROM teachers WHERE id = ?)"
	err = db.QueryRow(query, id).Scan(&studentCount)
	if err != nil {
		return 0, utils.ErrorHandler(err, "Error retrieving data")
	}

	return studentCount, nil

}
