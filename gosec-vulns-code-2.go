package main

import (
	"crypto/md5"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Vulnerability 1: Hardcoded credentials
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Vulnerability 2: SQL Injection
	query := "SELECT * FROM users WHERE name = '" + "Robert'); DROP TABLE users; --" + "'"
	_, err = db.Exec(query)
	if err != nil {
		log.Println(err)
	}

	// Vulnerability 3: Command Injection
	cmd := exec.Command("sh", "-c", "echo "+strings.TrimSpace("Hello; echo World"))
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(output))

	// Vulnerability 4: Weak cryptographic primitive (MD5)
	hash := md5.Sum([]byte("password"))
	fmt.Println(hash)

	// Vulnerability 5: Insecure use of SHA1
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hash := sha1.Sum([]byte("secret data"))
		fmt.Fprintf(w, "%x", hash)
	})
	http.ListenAndServe(":8080", nil)

	// Additional vulnerability: File inclusion vulnerability
	data, err := ioutil.ReadFile("/tmp/" + "sensitive_data.txt")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(data))
}
