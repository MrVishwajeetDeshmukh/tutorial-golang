package main

import "fmt"

// showLogicalOperator demonstrates how logical operators
// (AND, OR, NOT) work with boolean values in Go.
func showLogicalOperator() {
	// ----- AND (&&) -----
	// Returns true only if both sides are true.
	fmt.Print("AND operators ====\n")
	fmt.Printf("%t && %t = %t\n", true, true, true && true)     // both true → true
	fmt.Printf("%t && %t = %t\n", true, false, true && false)   // one false → false
	fmt.Printf("%t && %t = %t\n", false, true, false && true)   // one false → false
	fmt.Printf("%t && %t = %t\n", false, false, false && false) // both false → false
	fmt.Print("====================\n\n")

	// ----- OR (||) -----
	// Returns true if at least one side is true.
	fmt.Print("OR operators ====\n")
	fmt.Printf("%t || %t = %t\n", true, true, true || true)     // both true → true
	fmt.Printf("%t || %t = %t\n", true, false, true || false)   // one true → true
	fmt.Printf("%t || %t = %t\n", false, true, false || true)   // one true → true
	fmt.Printf("%t || %t = %t\n", false, false, false || false) // both false → false
	fmt.Print("====================\n\n")

	// ----- NOT (!) -----
	// Flips the boolean value — true becomes false, false becomes true.
	fmt.Print("NOT operators ====\n")
	fmt.Printf("!%t = %t\n", true, !true)   // not true → false
	fmt.Printf("!%t = %t\n", false, !false) // not false → true
	fmt.Print("====================\n\n")
}

func main() {
	// Call the function to display all logical operator examples
	showLogicalOperator()
}
