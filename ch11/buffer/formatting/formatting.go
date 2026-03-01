package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Enter name and age (e.g., John 30): ")
	fmt.Scanf("%s %d\n", &name, &age) // Use format string
	fmt.Printf("Entered name: %s, Age: %d\n", name, age)
}
