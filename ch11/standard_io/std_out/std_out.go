package main

import "fmt"

func main() {
	name := "Alice"
	age := 30

	// Basic output
	fmt.Print("Name: ")
	fmt.Print(name)
	fmt.Print("\n")

	// Automatic newline
	fmt.Println("Age:", age)

	// Formatted output
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}
