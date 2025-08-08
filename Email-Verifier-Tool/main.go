package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/AlperSeyman/email-verifier-tool/controller"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord")
	fmt.Print("Enter mail: ")

	if scanner.Scan() {
		domain_name := scanner.Text()
		controller.CheckDomain(domain_name)
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

}
