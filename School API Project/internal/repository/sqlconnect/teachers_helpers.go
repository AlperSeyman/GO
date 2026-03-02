package sqlconnect

import (
	"fmt"
	"net/http"
	"reflect"
	"restapi/internal/model"
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

func QueryFunc(r *http.Request, query string, args []any) (string, []any) {

	allowedColumns := getAllowedColumns()

	for param, values := range r.URL.Query() {
		if allowedColumns[param] {
			value := values[0]
			if value != "" {
				query += " AND " + param + " = ?"
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

func SortFunc(r *http.Request, query string) string {

	sortParams := r.URL.Query()["sortby"]
	if len(sortParams) == 0 {
		return query
	}

	var validSorts []string
	allowedColumns := getAllowedColumns()

	for _, param := range sortParams {
		parts := strings.Split(param, ":")
		if len(parts) != 2 {
			continue
		}
		field := parts[0]
		order := parts[1]

		if allowedColumns[field] && isValidSortOrder(order) {
			validSorts = append(validSorts, field+" "+strings.ToLower(order))
		}

		if isValidSortField(field) && isValidSortOrder(order) {
			validSorts = append(validSorts, field+" "+strings.ToLower(order))
		}
	}
	if len(validSorts) > 0 {
		query += " ORDER BY " + strings.Join(validSorts, ", ")
	}
	return query

}

func GenerateSelectQuery(model any) string {

	modelType := reflect.TypeOf(model)
	var columns string
	for i := 0; i < modelType.NumField(); i++ {
		dbTag := modelType.Field(i).Tag.Get("db")
		dbTag = strings.TrimSuffix(dbTag, ",omitempty")
		if dbTag != "" {
			if columns != "" {
				columns += ", "
			}
			columns += dbTag
		}
	}
	return fmt.Sprintf("SELECT %s FROM teachers", columns)

}

func GenerateInsertQuery(model any) string {

	modelType := reflect.TypeOf(model)

	var columns, placeholders string
	for i := 0; i < modelType.NumField(); i++ {
		dbTag := modelType.Field(i).Tag.Get("db")
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

func GenerateUpdateQuery(model any) string {

	modelType := reflect.TypeOf(model)
	var columns string
	for i := 0; i < modelType.NumField(); i++ {
		dbTag := modelType.Field(i).Tag.Get("db")
		dbTag = strings.TrimSuffix(dbTag, ",omitempty")
		if dbTag != "" && dbTag != "id" {
			if columns != "" {
				columns += ", "
			}
			columns += dbTag + " = ?"
		}
	}
	return fmt.Sprintf("UPDATE teachers SET %s", columns)
}

func GetStructValues(model any) []any {

	modelValue := reflect.ValueOf(model)
	modelType := modelValue.Type()
	values := []interface{}{}

	for i := 0; i < modelType.NumField(); i++ {
		dbTag := modelType.Field(i).Tag.Get("db")
		cleanTag := strings.TrimSuffix(dbTag, ",omitempty")
		if cleanTag != "" && cleanTag != "id" {
			values = append(values, modelValue.Field(i).Interface())
		}
	}
	return values
}

func GetStructPointers(modelPtr any) []any {

	modelValue := reflect.ValueOf(modelPtr).Elem()
	modelType := modelValue.Type()
	var pointers []any

	for i := 0; i < modelType.NumField(); i++ {
		dbTag := modelType.Field(i).Tag.Get("db")
		if dbTag != "" {
			pointers = append(pointers, modelValue.Field(i).Addr().Interface())
		}
	}
	return pointers
}
