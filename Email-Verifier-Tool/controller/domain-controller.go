package controller

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func CheckDomain(domain string) {

	// hasMX indicates whether the domain has MX (Mail Exchange) records.
	// If true, the domain can receive emails (e.g., contact@example.com is valid).
	var hasMX bool

	// hasSPF indicates whether the domain has an SPF (Sender Policy Framework) record.
	// SPF is used to specify which servers are allowed to send emails on behalf of the domain.
	var hasSPF bool

	// spfRecord contains the actual SPF TXT record from the domain's DNS.
	// It includes rules like: which IPs or domains can send email for this domain.
	var spfRecord string

	// hasDMARC indicates whether the domain has a DMARC (Domain-based Message Authentication) record.
	// DMARC protects against spoofing by enforcing policies on SPF/DKIM failures.
	var hasDMARC bool

	// dmarcRecord contains the actual DMARC TXT record from DNS.
	// It defines what should happen to emails that fail SPF/DKIM (e.g., reject, quarantine).
	var dmarcRecord string

	mxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Fatal(err)
	}

	if len(mxRecord) > 0 {
		hasMX = true
	}

	txtRecord, err := net.LookupTXT(domain)
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range txtRecord {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("Domanin: %v\n hasMX: %v\n hasSPF: %v\n sfRecord: %v\n hasDMARC: %v\n dmarcRecord: %v\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
