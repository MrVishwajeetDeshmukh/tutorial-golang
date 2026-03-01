package main

import (
	"fmt"
	"os"
)

func appendText() {
	f, _ := os.OpenFile(
		"/tmp/appendFile", os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		os.FileMode(0644))

	// Ensure the File resource is returned when the function terminates
	defer func() {
		fmt.Println("defer occurred!")
		f.Close()
	}()

	panic("Intentionally panicking")
}

func main() {
	appendText()
}
