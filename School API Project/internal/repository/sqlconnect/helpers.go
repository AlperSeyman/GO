package sqlconnect

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

func getAllowedColumns(model any) map[string]bool {

	allowed := make(map[string]bool)
	modelType := reflect.TypeOf(model)

	for i := 0; i < modelType.NumField(); i++ {
		dbTag := modelType.Field(i).Tag.Get("db")
		if dbTag != "" && dbTag != "-" {
			columnName := strings.Split(dbTag, ",")[0]
			allowed[columnName] = true
		}
	}
	return allowed
}

func GetAllowedFields(model any) map[string]struct{} {

	allowedFields := make(map[string]struct{})
	val := reflect.TypeOf(model)

	for i := 0; i < val.NumField(); i++ {
		fieldName := val.Field(i).Name
		if fieldName == "ID" {
			continue
		}
		tag := val.Field(i).Tag.Get("json")
		cleanTag := strings.Split(tag, ",")[0]
		if cleanTag != "" && cleanTag != "-" {
			allowedFields[cleanTag] = struct{}{}
		}
	}
	return allowedFields

}

func QueryFunc(r *http.Request, model any, query string, args []any) (string, []any) {

	allowedColumns := getAllowedColumns(model)

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
	return cleanOrder == "asc" || cleanOrder == "desc"
}

func isValidSortField(model any, field string) bool {

	allowedColumns := getAllowedColumns(model)
	cleanField := strings.ToLower(strings.TrimSpace(field))
	return allowedColumns[cleanField]

}

func SortFunc(r *http.Request, model any, query string) string {

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

		if isValidSortField(model, field) && isValidSortOrder(order) {
			validSorts = append(validSorts, field+" "+strings.ToLower(order))
		}
	}
	if len(validSorts) > 0 {
		query += " ORDER BY " + strings.Join(validSorts, ", ")
	}
	return query

}

func GenerateSelectQuery(model any, tableName string) string {

	modelType := reflect.TypeOf(model)
	var columns string
	for i := 0; i < modelType.NumField(); i++ {
		dbTag := modelType.Field(i).Tag.Get("db")
		cleanTag := strings.Split(dbTag, ",")[0]
		if cleanTag != "" {
			if columns != "" {
				columns += ", "
			}
			columns += cleanTag
		}
	}
	return fmt.Sprintf("SELECT %s FROM %s", columns, tableName)

}

func GenerateSelectQueryForTeacher(model any, tableName1 string, tableName2 string) string {

	modelType := reflect.TypeOf(model)
	var columns string
	for i := 0; i < modelType.NumField(); i++ {
		dbTag := modelType.Field(i).Tag.Get("db")
		cleanTag := strings.Split(dbTag, ",")[0]
		if cleanTag != "" {
			if columns != "" {
				columns += ", "
			}
			columns += cleanTag
		}
	}
	return fmt.Sprintf("SELECT %s FROM %s WHERE class = (SELECT class FROM %s WHERE id = ?)", columns, tableName1, tableName2)

}

func GenerateInsertQuery(model any, tableName string) string {

	modelType := reflect.TypeOf(model)
	var columns, placeholders string
	for i := 0; i < modelType.NumField(); i++ {
		dbTag := modelType.Field(i).Tag.Get("db")
		cleanTag := strings.Split(dbTag, ",")[0]
		if cleanTag != "" && cleanTag != "id" {
			if columns != "" {
				columns += ", "
				placeholders += ", "
			}
			columns += cleanTag
			placeholders += "?"
		}
	}
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, columns, placeholders)
}

func GenerateUpdateQuery(model any, tableName string) string {

	modelType := reflect.TypeOf(model)
	var columns string
	for i := 0; i < modelType.NumField(); i++ {
		dbTag := modelType.Field(i).Tag.Get("db")
		cleanTag := strings.Split(dbTag, ",")[0]
		if cleanTag != "" && cleanTag != "id" {
			if columns != "" {
				columns += ", "
			}
			columns += cleanTag + " = ?"
		}
	}
	return fmt.Sprintf("UPDATE %s SET %s WHERE id = ?", tableName, columns)
}

func GetStructValues(model any) []any {

	modelValue := reflect.ValueOf(model)
	modelType := modelValue.Type()
	values := []any{}

	for i := 0; i < modelType.NumField(); i++ {
		dbTag := modelType.Field(i).Tag.Get("db")
		cleanTag := strings.Split(dbTag, ",")[0]
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
		cleanTag := strings.Split(dbTag, ",")[0]
		if cleanTag != "" {
			pointers = append(pointers, modelValue.Field(i).Addr().Interface())
		}
	}
	return pointers
}
