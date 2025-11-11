package main

import "fmt"

func main() {
	// In Go, arithmetic operators work with two values (called binary operators)
	// Example: left operand  operator  right operand
	format := 1 + 2
	fmt.Println(format) // 3

	// + adds the two numbers together
	sum := 5 + 6
	fmt.Println(sum) // 11

	// - subtracts the right number from the left one
	difference := 5 - 6
	fmt.Println(difference) // -1

	// * multiplies the two numbers
	product := 5 * 6
	fmt.Println(product) // 30

	// / divides the left number by the right one (integer division)
	quotient := 6 / 2
	fmt.Println(quotient) // 3

	// % gives the remainder after dividing the two numbers
	remainder := 10 % 3
	fmt.Println(remainder) // 1
}
