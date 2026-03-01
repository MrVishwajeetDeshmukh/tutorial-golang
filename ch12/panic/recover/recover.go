package main

import (
	"fmt"
	"os"
)

func appendText() {
	f, _ := os.OpenFile(
		"/tmp/appendFile", os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		os.FileMode(0644))

	defer func() {
		fmt.Println("defer occurred!")
		f.Close()
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic! recover: %v\n", r)
		}
	}()

	panic("Intentionally panicking")
}

func main() {
	appendText()
}
