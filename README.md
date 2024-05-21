# Domain Records Checker

This project is a simple Go application that checks domain records for the presence of MX, SPF, and DMARC settings. It reads domains from standard input and outputs the results in CSV format.

## Features

- Checks if the domain has MX records.
- Checks if the domain has SPF records and retrieves the SPF record.
- Checks if the domain has DMARC records and retrieves the DMARC record.

## Project Structure

- `main.go`: The Go source code for the domain records checker.

## Prerequisites

- Go installed on your system. You can download and install Go from the [official Go website](https://golang.org/dl/).

## Getting Started

1. **Clone the repository or create the project directory**:

    ```sh
    mkdir project-directory
    cd project-directory
    ```

2. **Create the Go file**:

    - `main.go`:

    ```go
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
    ```

3. **Build and run the application**:

    ```sh
    go build -o domain-checker
    ```

    This will create an executable file named `domain-checker`.

4. **Run the application**:

    You can run the application and provide domain names through standard input. For example:

    ```sh
    echo "example.com" | ./domain-checker
    ```

    Or you can provide multiple domains from a file:

    ```sh
    cat domains.txt | ./domain-checker
    ```

## Output

The application outputs the results in CSV format with the following columns:
- `domain`: The domain name.
- `hasMX`: Whether the domain has MX records (true/false).
- `hasSPF`: Whether the domain has SPF records (true/false).
- `spfRecord`: The SPF record (if any).
- `hasDMRAC`: Whether the domain has DMARC records (true/false).
- `dmarcRecord`: The DMARC record (if any).
