package main

import "fmt"

func main() {
	// Declare two string variables.
	text1 := "hello"
	text2 := "world"

	// In Go, the '+' operator is used to concatenate (join) strings.
	// Here we combine text1 and text2 with a space in between.
	message := text1 + " " + text2

	// Print the combined string.
	fmt.Println(message)
}
