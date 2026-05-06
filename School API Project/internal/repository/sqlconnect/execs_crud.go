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

func GetExecsDbHandler(r *http.Request) ([]model.Exec, error) {

	db, err := ConnectDB()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error connecting to database.")
	}
	defer db.Close()

	var execs []model.Exec

	query := GenerateSelectQuery(model.Exec{}, "execs") + " WHERE 1 = 1"
	var args []any

	// filtering
	query, args = QueryFunc(r, model.Exec{}, query, args)

	// sorting
	query = SortFunc(r, model.Exec{}, query)
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Database query error")
	}

	for rows.Next() {
		var exec model.Exec
		structPointers := GetStructPointers(&exec)
		err = rows.Scan(structPointers...)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error scanning database results")
		}
		execs = append(execs, exec)
	}

	err = rows.Err()
	if err != nil {
		return nil, utils.ErrorHandler(err, "The database connection was lost while reading rows")
	}
	return execs, nil
}

func GetExecByIdHandler(id int) (model.Exec, error) {

	db, err := ConnectDB()
	if err != nil {
		return model.Exec{}, utils.ErrorHandler(err, "Error connecting to database...")
	}
	defer db.Close()

	var exec model.Exec

	query := GenerateSelectQuery(model.Exec{}, "execs") + " WHERE id = ?"

	structPointers := GetStructPointers(&exec)
	err = db.QueryRow(query, id).Scan(structPointers...)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.Exec{}, utils.ErrorHandler(err, "Exec not found")
		}
		return model.Exec{}, utils.ErrorHandler(err, "Database query error")
	}
	return exec, nil
}

func AddExecDbHandler(newExecs []model.Exec) ([]model.Exec, error) {

	db, err := ConnectDB()

	if err != nil {
		return nil, utils.ErrorHandler(err, "Error connecting to database..")
	}
	defer db.Close()

	query := GenerateInsertQuery(model.Exec{}, "execs")

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error in preparing SQL query")
	}
	defer stmt.Close()

	addedExecs := make([]model.Exec, len(newExecs))
	for i, newExec := range newExecs {
		structValues := GetStructValues(newExec)
		result, err := stmt.Exec(structValues...)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error inserting data into database")
		}
		lastID, err := result.LastInsertId()
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error getting last insert ID")
		}
		newExec.ID = int(lastID)
		addedExecs[i] = newExec
	}
	return addedExecs, nil
}

func UpdateExecsDbHandler(id int, updatedExec model.Exec) (model.Exec, error) {

	db, err := ConnectDB()
	if err != nil {
		return model.Exec{}, utils.ErrorHandler(err, "Error connecting to database.")
	}
	defer db.Close()

	var existingExec model.Exec

	query := GenerateSelectQuery(model.Exec{}, "execs") + " WHERE id = ?"

	structPointers := GetStructPointers(&existingExec)

	err = db.QueryRow(query, id).Scan(structPointers...)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.Exec{}, utils.ErrorHandler(err, "Exec not found")
		}
		return model.Exec{}, utils.ErrorHandler(err, "Unable to retrieve data")
	}

	updatedExec.ID = existingExec.ID

	query = GenerateUpdateQuery(model.Exec{}, "execs")

	execValues := GetStructValues(updatedExec)
	execValues = append(execValues, id)
	_, err = db.Exec(query, execValues...)
	if err != nil {
		return model.Exec{}, utils.ErrorHandler(err, "Error updating exec")
	}
	return updatedExec, nil
}

func PatchExecsDbHandler(updatedExecs []map[string]any) error {

	db, err := ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return utils.ErrorHandler(err, "Error starting transaction")
	}

	for _, updatedExec := range updatedExecs {
		idFloat, ok := updatedExec["id"].(float64)
		if !ok {
			return utils.ErrorHandler(err, "Invalid exec ID in update")
		}
		id := int(idFloat)

		var existingExec model.Exec

		query := GenerateSelectQuery(model.Exec{}, "execs") + " WHERE id = ?"

		structPointer := GetStructPointers(&existingExec)

		err = tx.QueryRow(query, id).Scan(structPointer...)
		if err != nil {
			tx.Rollback()
			if err == sql.ErrNoRows {
				return utils.ErrorHandler(err, "Exec not found")
			}
			return nil
		}

		execValue := reflect.ValueOf(&existingExec).Elem()
		execType := execValue.Type()

		for column, value := range updatedExec {
			if column == "id" {
				continue
			}
			for i := 0; i < execType.NumField(); i++ {
				field := execType.Field(i)
				cleanTag := strings.Split(field.Tag.Get("json"), ",")[0]
				if cleanTag == column {
					fieldToUpdate := execValue.Field(i) // old value
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
					break
				}
			}
		}

		query = GenerateUpdateQuery(model.Exec{}, "execs")
		execValues := GetStructValues(existingExec)
		execValues = append(execValues, id)
		_, err = tx.Exec(query, execValues...)
		if err != nil {
			tx.Rollback()
			return utils.ErrorHandler(err, "Error updating exec")
		}
	}
	err = tx.Commit()
	if err != nil {
		return utils.ErrorHandler(err, "Error committing transaction")
	}
	return nil
}

func PatchOneExecsDbHandler(id int, updatedExec map[string]any) (model.Exec, error) {

	db, err := ConnectDB()
	if err != nil {
		return model.Exec{}, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	var existingExec model.Exec
	query := GenerateSelectQuery(model.Exec{}, "execs") + " WHERE id = ?"

	structPointers := GetStructPointers(&existingExec)
	err = db.QueryRow(query, id).Scan(structPointers...)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Exec{}, utils.ErrorHandler(err, "Exec not found")
		}
		return model.Exec{}, utils.ErrorHandler(err, "Unable to retrieve data")
	}

	execValues := reflect.ValueOf(&existingExec).Elem()
	execType := execValues.Type()
	for column, value := range updatedExec {
		if column == "id" {
			continue
		}
		for i := 0; i < execType.NumField(); i++ {
			field := execType.Field(i)
			cleanTag := strings.Split(field.Tag.Get("json"), ",")[0]
			if cleanTag == column {
				if execValues.Field(i).CanSet() {
					execField := execValues.Field(i)
					execField.Set(reflect.ValueOf(value).Convert(execValues.Field(i).Type()))

				}
			}
		}
	}
	query = GenerateUpdateQuery(model.Exec{}, "execs")
	structValues := GetStructValues(existingExec)
	structValues = append(structValues, id)
	_, err = db.Exec(query, structValues...)
	if err != nil {
		return model.Exec{}, utils.ErrorHandler(err, "Error updating exec")
	}
	return existingExec, nil

}

func DeleteOneExecsDbHandler(id int) error {

	db, err := ConnectDB()
	if err != nil {
		return utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	query := "DELETE FROM execs WHERE id = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		return utils.ErrorHandler(err, "Error deleting exec")
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return utils.ErrorHandler(err, "Error retrieving delete result")
	}
	if rowAffected == 0 {
		return utils.ErrorHandler(err, "Exec not found")
	}
	return nil
}

func DeleteExecsDbHandler(ids []int) ([]int, error) {

	db, err := ConnectDB()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error connecting to databae")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error starting transaction")
	}

	query := "DELETE FROM execs WHERE id = ?"
	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return nil, utils.ErrorHandler(err, "Error preparing delete statement")
	}
	defer stmt.Close()

	var deletedIDs []int
	for _, id := range ids {
		result, err := stmt.Exec(id)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error deleting execs")
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error retrieving delete result")
		}
		if rowsAffected > 0 {
			deletedIDs = append(deletedIDs, id)
		} else {
			tx.Rollback()
			return nil, utils.ErrorHandler(err, fmt.Sprintf("ID %d does not exist", id))
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error committing transaction")
	}
	if len(deletedIDs) == 0 {
		return nil, utils.ErrorHandler(err, "No execs were deleted")
	}
	return deletedIDs, nil
}
