package main

import (
	"fmt"
	"os"
)

func main() {
	if err := doSomething(); err != nil {
		fmt.Println("Error occurred:", err)
		os.Exit(1) // Set abnormal exit code 1
	}
	os.Exit(0) // Set normal exit code 0
}

func doSomething() error {
	// A dummy function that generates an error
	return fmt.Errorf("A problem occurred.")
}
