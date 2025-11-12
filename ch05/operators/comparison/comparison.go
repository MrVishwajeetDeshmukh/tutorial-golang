package main

import "fmt"

func main() {
	// ==  checks if two values are equal
	fmt.Printf("%v\t=\ttrue == false\n", true == false) // false

	// >  checks if the left value is greater than the right value
	fmt.Printf("%v\t=\t2 > 3\n", 2 > 3)               // false
	fmt.Printf("%v\t=\t\"go\" > \"c\"\n", "go" > "c") // true (compares alphabetically by Unicode value)

	// >=  checks if the left value is greater than or equal to the right value
	fmt.Printf("%v\t=\t3.5 >= 3\n", 3.5 >= 3)    // true
	fmt.Printf("%v\t=\t'Z' >= 'A'\n", 'Z' >= 'A') // true (compares ASCII/Unicode values)

	// <  checks if the left value is less than the right value
	fmt.Printf("%v\t=\t-1 < -1.1\n", -1 < -1.1) // false

	// <=  checks if the left value is less than or equal to the right value
	fmt.Printf("%v\t=\t3.5 <= 3.5\n", 3.5 <= 3.5) // true
}
