package sqlconnect

import (
	"database/sql"
	"net/http"
	"reflect"
	"restapi/internal/model"
	"restapi/pkg/utils"
)

func GetStudentsDbHandler(r *http.Request) ([]model.Student, error) {

	db, err := ConnectDB()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error connecting to database.")
	}
	defer db.Close()

	var students []model.Student

	query := GenerateSelectQuery(model.Student{}, "students") + " WHERE 1 = 1"
	var args []any

	// filtering
	query, args = QueryFunc(r, model.Student{}, query, args)

	// sorting
	query = SortFunc(r, model.Student{}, query)
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Database query error")
	}
	for rows.Next() {
		var student model.Student
		structPointers := GetStructPointers(&student)
		err = rows.Scan(structPointers...)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error scanning database results")
		}
		students = append(students, student)
	}

	err = rows.Err()
	if err != nil {
		return nil, utils.ErrorHandler(err, "The database connection was lost while reading rows")
	}
	return students, nil
}

func GetStudentByIdHandler(id int) (model.Student, error) {

	db, err := ConnectDB()
	if err != nil {
		return model.Student{}, utils.ErrorHandler(err, "Error connecting to database...")
	}
	defer db.Close()

	var student model.Student

	query := GenerateSelectQuery(model.Student{}, "students") + " WHERE id = ?"

	structPointers := GetStructPointers(&student)
	err = db.QueryRow(query, id).Scan(structPointers...)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Student{}, utils.ErrorHandler(err, "Student not found")
		}
		return model.Student{}, utils.ErrorHandler(err, "Database query error")
	}
	return student, nil
}

func AddStudentDbHandler(newStudents []model.Student) ([]model.Student, error) {

	db, err := ConnectDB()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error connecting to database..")
	}
	defer db.Close()

	query := GenerateInsertQuery(model.Student{}, "students")
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error in preparing SQL query")
	}
	defer stmt.Close()

	addedStudents := make([]model.Student, len(newStudents))
	for i, newStudent := range newStudents {
		structValues := GetStructValues(newStudent)
		result, err := stmt.Exec(structValues...)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error inserting data into database")
		}
		lastID, err := result.LastInsertId()
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error getting last insert ID")
		}
		newStudent.ID = int(lastID)
		addedStudents[i] = newStudent
	}
	return addedStudents, nil
}

func UpdateStudentsDbHandler(id int, updatedStudent model.Student) (model.Student, error) {

	db, err := ConnectDB()
	if err != nil {
		return model.Student{}, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	var existingStudent model.Student

	query := GenerateSelectQuery(model.Student{}, "students") + " WHERE id = ?"

	structPointers := GetStructPointers(&existingStudent)

	err = db.QueryRow(query, id).Scan(structPointers...)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Student{}, utils.ErrorHandler(err, "Student not found")
		}
		return model.Student{}, utils.ErrorHandler(err, "Unable to retrieve data")
	}

	updatedStudent.ID = existingStudent.ID

	query = GenerateUpdateQuery(model.Student{}, "students")
	structValues := GetStructValues(updatedStudent)
	structValues = append(structValues, id)
	_, err = db.Exec(query, structValues...)
	if err != nil {
		return model.Student{}, utils.ErrorHandler(err, "Error updating student")
	}
	return updatedStudent, nil
}

func PatchStudentsDbHandler(updatedStudents []map[string]any) error {

	db, err := ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return utils.ErrorHandler(err, "Error starting transaction")
	}

	for _, updatedStudent := range updatedStudents {
		idFloat, ok := updatedStudent["id"].(float64)
		if !ok {
			return utils.ErrorHandler(err, "Invalid student ID in update")
		}
		id := int(idFloat)

		var existingStudent model.Student
		query := GenerateSelectQuery(model.Student{}, "students") + " WHERE id = ?"

		structPointer := GetStructPointers(&existingStudent)

		err = db.QueryRow(query, id).Scan(structPointer...)
		if err != nil {
			tx.Rollback()
			if err == sql.ErrNoRows {
				return utils.ErrorHandler(err, "Student not found")
			}
			return nil
		}

		studentValue := reflect.ValueOf(&existingStudent).Elem()
		studentType := studentValue.Type()

		for column, value := range updatedStudent {
			if column == "id" {
				continue
			}
			for i := 0; i < studentType.NumField(); i++ {
				field := studentType.Field(i)
				fieldTag := field.Tag.Get("json")
				if fieldTag == column+",omitempty" {
					fielValue := studentValue.Field(i) // old value
					if fielValue.CanSet() {
						val := reflect.ValueOf(value) // new value
						if val.Type().ConvertibleTo(fielValue.Type()) {
							fielValue.Set(val.Convert(fielValue.Type()))
						} else {
							tx.Rollback()
							return utils.ErrorHandler(err, "Error updating data")
						}
					}
					break
				}
			}
		}

		query = GenerateUpdateQuery(model.Student{}, "students")

		structValues := GetStructValues(existingStudent)
		structValues = append(structValues, id)
		_, err = tx.Exec(query, structValues...)
		if err != nil {
			tx.Rollback()
			return utils.ErrorHandler(err, "Error updating student")
		}
	}
	err = tx.Commit()
	if err != nil {
		return utils.ErrorHandler(err, "Error committing transaction")
	}

	return nil

}
