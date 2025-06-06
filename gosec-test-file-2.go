package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/tls"
	"crypto/rand"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// 1. Hardcoded credentials [G101]
	password : "supersecret123"

	// 2. Command injection [G204]
	userInput : "ls -la"
	cmd : exec.Command("sh", "-c", userInput)
	output, _ : cmd.Output()
	fmt.Println(string(output))

	// 3. SQL Injection [G201]
	db, err : sql.Open("postgres", "user=foo dbname=bar sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	name : "'; DROP TABLE users; --"
	query : fmt.Sprintf("SELECT * FROM users WHERE name='%s'", name)
	rows, _ : db.Query(query)
	defer rows.Close()

	// 4. Insecure TLS settings [G402]
	_ = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// 5. Predictable random number generator [G404]
	fmt.Println("Random:", rand.Intn(100))

	// 6. Insecure temporary file [G304]
	tmpfile, err : ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// 7. Insecure file permissions [G302]
	err = ioutil.WriteFile("secrets.txt", []byte("top-secret"), 0644)
	if err != nil {
		log.Fatal(err)
	}

	// 8. Use of MD5 [G401]
	h : md5.New()
	h.Write([]byte("test"))
	fmt.Printf("MD5: %x\n", h.Sum(nil))

	// 9. Use of SHA1 [G401]
	h2 : sha1.New()
	h2.Write([]byte("test"))
	fmt.Printf("SHA1: %x\n", h2.Sum(nil))

	// 10. Insecure HTTP request [G107]
	resp, err : http.Get("http://example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 11. Ignoring errors [G103]
	ioutil.WriteFile("ignored_error.txt", []byte("some data"), 0644)
}
