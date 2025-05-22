package main

import (
	"fmt"
)

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime int
	recipentName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %v", bm.recipentName, bm.birthdayTime)
}

type sendingReport struct {
	reportName   string
	numberOfsend int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf("Your %s report is  ready. You've sent %v", sr.reportName, sr.numberOfsend)
}

func main() {

	s := sendingReport{
		reportName:   "first",
		numberOfsend: 10,
	}
	sendMessage(s)

	b := birthdayMessage{
		birthdayTime: 1971,
		recipentName: "Jansen",
	}
	sendMessage(b)
}
