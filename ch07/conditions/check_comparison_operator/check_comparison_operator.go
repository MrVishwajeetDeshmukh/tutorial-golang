package main

import "fmt"

func main() {
	// Define a string variable containing the word "사과" (Korean for "apple")
	korWord := "사과"

	// Compare the variable with the literal string "사과"
	// The comparison returns true because both strings are identical.
	fmt.Printf("korWord == \"%v\" = %v\n", korWord, korWord == "사과")
}
