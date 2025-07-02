package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Hardcoded credentials
	dbUser := "root"
	dbPass := "password123"
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/testdb", dbUser, dbPass)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/run", func(w http.ResponseWriter, r *http.Request) {
		// Command injection via unsanitized input
		cmd := exec.Command("sh", "-c", r.URL.Query().Get("cmd"))
		out, err := cmd.CombinedOutput()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write(out)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// SQL Injection vulnerability
		query := "SELECT * FROM users WHERE username = '" + username + "' AND password = '" + password + "'"
		rows, err := db.Query(query)
		if err != nil {
			http.Error(w, "DB error", 500)
			return
		}
		defer rows.Close()

		if rows.Next() {
			fmt.Fprintf(w, "Welcome %s!", username)
		} else {
			http.Error(w, "Unauthorized", 401)
		}
	})

	http.HandleFunc("/hash", func(w http.ResponseWriter, r *http.Request) {
		// Weak hashing algorithm
		data := r.URL.Query().Get("data")
		hash := md5.Sum([]byte(data))
		fmt.Fprintf(w, "%x", hash)
	})

	// Insecure random number generation
	rand.Seed(time.Now().UnixNano())
	token := rand.Intn(1000000)
	fmt.Println("Generated token:", token)

	// Insecure file creation with broad permissions
	err = os.WriteFile("/tmp/insecure_file.txt", []byte("secret data"), 0777)
	if err != nil {
		log.Fatal(err)
	}

	// Insecure HTTP server (no TLS)
	http.ListenAndServe(":8080", nil)
}
/##
