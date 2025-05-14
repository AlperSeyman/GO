package helper

import "strings"

func UserValidation(first_name string, last_name string, email string, user_tickets uint, remaining_tickets uint) (bool, bool, bool) {
	isValidName := len(first_name) >= 2 && len(last_name) >= 2 // return boolen result, true or false

	isValidEmail := strings.Contains(email, "@") //  return boolen result, true or false

	isValidTicketNumber := user_tickets > 0 && user_tickets <= uint(remaining_tickets)

	return isValidName, isValidEmail, isValidTicketNumber
}
