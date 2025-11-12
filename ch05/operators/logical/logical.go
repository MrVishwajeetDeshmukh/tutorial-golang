package main

import "fmt"

func main() {
	// Logical AND (&&) operator
	// Returns true only if BOTH sides are true
	fmt.Println("AND operator")
	fmt.Println("========")
	fmt.Printf("%v\t=\ttrue && true\n", true && true)   // true  (both true)
	fmt.Printf("%v\t=\ttrue && false\n", true && false) // false (right side is false)
	fmt.Printf("%v\t=\tfalse && true\n", false && true) // false (left side is false)
	fmt.Printf("%v\t=\tfalse && false\n", false && false) // false (both false)
	fmt.Println("========")

	// Logical OR (||) operator
	// Returns true if AT LEAST ONE side is true
	fmt.Println("OR operator")
	fmt.Println("========")
	fmt.Printf("%v\t=\ttrue || true\n", true || true)   // true  (both true)
	fmt.Printf("%v\t=\ttrue || false\n", true || false) // true  (left side is true)
	fmt.Printf("%v\t=\tfalse || true\n", false || true) // true  (right side is true)
	fmt.Printf("%v\t=\tfalse || false\n", false || false) // false (both false)
	fmt.Println("========")

	// Logical NOT (!) operator
	// Reverses the boolean value
	fmt.Println("NOT operator")
	fmt.Println("========")
	fmt.Printf("%v\t=\t!true\n", !true)   // false (negation of true)
	fmt.Printf("%v\t=\t!false\n", !false) // true  (negation of false)
}
