package main

import (
	"bytes"
	"database/sql"
	"log"
	"os/exec"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func openDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:mariadb@tcp(localhost:3306)/mariadb")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func runSubfinder(domain string) ([]string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	// Prepare the subfinder command
	cmd := exec.Command("subfinder", "-d", domain)

	// Capture the output
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		log.Printf("Error running subfinder: %s\n", stderr.String())
		return nil, err
	}

	// Split the output into lines
	lines := strings.Split(stdout.String(), "\n")

	// Filter out subdomains that start with 'www'
	var subdomains []string
	for _, line := range lines {
		if !strings.HasPrefix(line, "www.") {
			subdomains = append(subdomains, line)
		}
	}

	return subdomains, nil
}

func storeSubdomain(db *sql.DB, subdomain string) error {
	query := `
    INSERT INTO subdomains (name) 
    VALUES (?) 
    ON DUPLICATE KEY UPDATE 
    checked = CURRENT_TIMESTAMP;
    `

	_, err := db.Exec(query, subdomain)
	return err
}

func main() {
	db, err := openDatabase()
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	domain := "stutz-medien.dev"
	subdomains, err := runSubfinder(domain)
	if err != nil {
		log.Fatal(err)
	}

	for _, subdomain := range subdomains {
		if subdomain != "" {
			err := storeSubdomain(db, subdomain)
			if err != nil {
				log.Println("Error storing subdomain: ", err)
			} else {
				log.Println("Stored subdomain: ", subdomain)
			}
		}
	}
}
