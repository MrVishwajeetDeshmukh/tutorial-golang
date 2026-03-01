package main

import "fmt"

func main() {
	var name string
	var age int

	// Get input from the user
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter your age: ")
	fmt.Scanf("%d\n", &age)

	// Print the entered data
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}
