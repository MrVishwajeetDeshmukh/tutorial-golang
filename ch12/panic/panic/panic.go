package main

import (
	"os"
)

func openFile(filePath string) {
	_, err := os.OpenFile(filePath, os.O_RDONLY, os.FileMode(0644))
	if err != nil {
		panic(err) // Panic if the file cannot be opened.
	}
}

func main() {
	filePath := "noFile"
	openFile(filePath) // 1. Attempt to open a non-existent file to cause a panic.
}
