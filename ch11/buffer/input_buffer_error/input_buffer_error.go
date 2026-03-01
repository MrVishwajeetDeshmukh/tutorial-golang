package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Enter the first number: ")
	fmt.Scan(&a) // Newline character remains in the input buffer

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&b) // Input processing ends due to the newline character from previous input

	fmt.Printf("Entered values: %d, %d\n", a, b)
}
