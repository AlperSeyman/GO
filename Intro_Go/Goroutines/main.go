package main

import (
	"fmt"
)

func logMessage(chEmails, chSms chan string) {

	for {
		select {
		case email, ok := <-chEmails:
			if !ok {
				return
			}
			logEmail(email)
		case sms, ok := <-chSms:
			if !ok {
				return
			}
			logSms(sms)
		}
	}
}

func logSms(sms string) {
	fmt.Println("SMS:", sms)
}

func logEmail(email string) {
	fmt.Println("E-Mail:", email)
}

func main() {

	chEmail := make(chan string)
	chSMS := make(chan string)

	go logMessage(chEmail, chSMS)

	chEmail <- "Email"
	chSMS <- "SMS"

}
