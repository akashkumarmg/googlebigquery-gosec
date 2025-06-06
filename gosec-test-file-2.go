package main

import (
	"fmt"
	"net/http"
)

func main() {
	// G101: Hardcoded credentials
	username := "admin"
	password := "SuperSecret123" // <-- Vulnerable

	req, err := http.NewRequest("GET", "https://example.com", nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(username, password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
}
