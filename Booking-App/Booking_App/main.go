package main

import (
	"fmt"
	"strings"
)

// Package Level
var conference_name = "Go Lang Conference"

const conference_tickets int = 50

var remaining_tickets uint = 50
var bookings = []string{}

func main() {

	// call function 'greetUser'
	greetUser()

	// Arrays and Slices
	// Array --->  var bookings [50]string or  var bookings = [50]string{}
	// Slices --->  var bookings = []string{} or bookings := []string{}

	// Infinte Loop
	for remaining_tickets > 0 && len(bookings) < 50 { // for true {} ---> Infinite Loop

		// call function 'userInput'
		first_name, last_name, email, user_tickets := userInput()
		//  call function 'userValidation'
		isValidName, isValidEmail, isValidTicketNumber := userValidation(first_name, last_name, email, user_tickets)

		// If...Else Statement & Boolean
		if isValidName && isValidEmail && isValidTicketNumber {

			// call function 'booking'
			bookTickets(user_tickets, first_name, last_name, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are %v\n", firstNames)

			// If...Else Statement & Boolean
			if remaining_tickets <= 0 {
				// end program
				fmt.Println("Our Conference is booked out. See you next conference....")
				break

			}

		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain '@' sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tictkest you entered is invalid")
			}
		}
	}

}

// Function with parameters
func greetUser() {
	fmt.Printf("Welcome to %s booking application.\n", conference_name)
	fmt.Printf("We have total of %v tickets and %v are still avaliable.\n", conference_tickets, remaining_tickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, value := range bookings { //To only show value or index, we can omit the other output using an underscore (_)
		var full_names = strings.Fields(value) //Splits the string with white space as separator.
		firstNames = append(firstNames, full_names[0])
	}
	return firstNames
}

func userValidation(first_name string, last_name string, email string, user_tickets uint) (bool, bool, bool) {
	isValidName := len(first_name) >= 2 && len(last_name) >= 2 // return boolen result, true or false

	isValidEmail := strings.Contains(email, "@") //  return boolen result, true or false

	isValidTicketNumber := user_tickets > 0 && user_tickets <= uint(remaining_tickets)

	return isValidName, isValidEmail, isValidTicketNumber
}

func userInput() (string, string, string, uint) {

	var first_name string
	var last_name string
	var email string
	var user_tickets uint
	// Getting User Input
	fmt.Println("Enter your first name : ")
	fmt.Scan(&first_name)

	fmt.Println("Enter your last name : ")
	fmt.Scan(&last_name)

	fmt.Println("Enter your email address : ")
	fmt.Scan(&email)

	fmt.Println("Enter number of ticktets : ")
	fmt.Scan(&user_tickets)

	return first_name, last_name, email, user_tickets
}

func bookTickets(user_tickets uint, first_name string, last_name string, email string) {
	remaining_tickets = remaining_tickets - uint(user_tickets)
	bookings = append(bookings, first_name+" "+last_name)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", first_name, last_name, user_tickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remaining_tickets, conference_name)

}
