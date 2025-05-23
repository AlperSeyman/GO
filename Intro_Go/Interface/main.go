package main

import "fmt"

func getExpenReport(e expense) (string, float64) {

	email, ok := e.(email)
	if ok {
		return email.toAddress, email.cost()
	}
	sms, ok := e.(sms)
	if ok {
		return sms.toPhoneNumber, sms.cost()
	}
	return "", 0.0

}

type expense interface {
	cost() float64
}

type printter interface {
	print()
}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return 0.05 * float64(len(e.body))
	}
	return 0.01 * float64(len(e.body))
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return .1 * float64(len(s.body))
	}
	return 0.03 * float64(len(s.body))
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

func main() {

	e := email{
		isSubscribed: true,
		body:         "I want my money back",
		toAddress:    "aaaaa@bbbbb.com",
	}

	sms := sms{
		isSubscribed:  false,
		body:          "Confirmation Code",
		toPhoneNumber: "12113123123",
	}

	fmt.Println(getExpenReport(e))
	fmt.Println("******************")
	fmt.Println(getExpenReport(sms))

}
