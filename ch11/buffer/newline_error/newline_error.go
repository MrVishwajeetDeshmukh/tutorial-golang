package main

import "fmt"

func main() {
	var name string
	var age int
	var message string

	fmt.Scanf("%s %d\n", &name, &age)
	fmt.Scanln(&message) // Terminates without input due to remaining newline character
}
