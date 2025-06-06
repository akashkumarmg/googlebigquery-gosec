package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// G204: Subprocess launched with variable
	userInput := "ls -la"
	cmd := exec.Command("sh", "-c", userInput)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(output))
}
##
