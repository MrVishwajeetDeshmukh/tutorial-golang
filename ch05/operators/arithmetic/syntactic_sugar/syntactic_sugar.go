package main

import "fmt"

func main() {
	operand := 5

	// Normal way of adding two numbers
	sum := 10

	// Add operand (5) to sum (10) and store the result back into sum
	sum = sum + operand
	fmt.Println(sum) // prints 15

	// Shortcut way (syntactic sugar) of doing the same thing
	sumSyntacticSugar := 10

	// Instead of writing sumSyntacticSugar = sumSyntacticSugar + operand
	// We can just use += to make it shorter
	sumSyntacticSugar += operand
	fmt.Println(sumSyntacticSugar) // prints 15

	// Subtract and assign
	difference := 10
	difference -= 5 // same as difference = difference - 5
	fmt.Println(difference) // prints 5

	// Multiply and assign
	product := 5
	product *= 4 // same as product = product * 4
	fmt.Println(product) // prints 20

	// Divide and assign
	quotient := 20
	quotient /= 5 // same as quotient = quotient / 5
	fmt.Println(quotient) // prints 4

	// Find remainder and assign
	remainder := 10
	remainder %= 3 // same as remainder = remainder % 3
	fmt.Println(remainder) // prints 1
}
