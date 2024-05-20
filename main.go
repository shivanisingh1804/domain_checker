package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain,hasMX,hasSPF,spfRecord,hasDMRAC,dmarcRecord\n")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("ERROR: could not read from the input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMRAC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMRAC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v,%v,%v,%v,%v,%v\n", domain, hasMX, hasSPF, spfRecord, hasDMRAC, dmarcRecord)
}

// Certainly! Let's break down the code step by step:

// 1. **Imports**:
//    - The code imports necessary packages:
//      - `bufio`: For reading input from standard input.
//      - `fmt`: For formatted I/O.
//      - `log`: For logging errors.
//      - `net`: For network-related operations.

// 2. **`main()` Function**:
//    - The `main()` function is the entry point of the program.
//    - It initializes a scanner to read input from the standard input (keyboard).
//    - It prints the header for the output CSV: `"domain,hasMX,hasSPF,spfRecord,hasDMRAC,dmarcRecord"`.

// 3. **`checkDomain(domain string)` Function**:
//    - This function takes a domain name as input and checks various DNS records related to it.
//    - It initializes boolean variables `hasMX`, `hasSPF`, and `hasDMRAC`.
//    - It also initializes strings `spfRecord` and `dmarcRecord`.

// 4. **MX Records**:
//    - It looks up MX (Mail Exchange) records for the given domain using `net.LookupMX(domain)`.
//    - If there are MX records, it sets `hasMX` to `true`.

// 5. **TXT Records (SPF)**:
//    - It looks up TXT records for the given domain using `net.LookupTXT(domain)`.
//    - If any record starts with `"v=spf1"`, it sets `hasSPF` to `true` and assigns the record to `spfRecord`.

// 6. **TXT Records (DMARC)**:
//    - It looks up TXT records for the `_dmarc.<domain>` using `net.LookupTXT("_dmarc." + domain)`.
//    - If any record starts with `"v=DMARC1"`, it sets `hasDMRAC` to `true` and assigns the record to `dmarcRecord`.

// 7. **Output**:
//    - Finally, it prints the results in CSV format: `"%v,%v,%v,%v,%v,%v\n"`.
//      - `%v` represents the values of the variables.
//      - The order is: `domain, hasMX, hasSPF, spfRecord, hasDMRAC, dmarcRecord`.

// This program checks DNS records for a given domain and reports whether it has MX records, SPF records, and DMARC records. It's useful for analyzing domain configurations related to email delivery and security. 🌐🔍
