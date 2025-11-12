package main

import "fmt"

func main() {
	// Assignment operator (=) is used to store a value in a variable
	var x int   // declare a variable named x of type int
	x = 5       // assign 5 to x
	fmt.Println(x) // prints: 5

	// Short variable declaration (:=) lets you declare and assign in one step
	y := 10     // Go automatically infers the type (int in this case)
	fmt.Println(y) // prints: 10
}
