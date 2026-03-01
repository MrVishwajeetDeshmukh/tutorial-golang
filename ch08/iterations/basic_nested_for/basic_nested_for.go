package main

import "fmt"

// Prints the multiplication table for a specific number (EX: 2-times table, 4-times table)
func displayMultiplicationTableAt(tableIndex int) {
	for i := 1; i <= 9; i++ {
		fmt.Printf("%d x %d = %d\n", tableIndex, i, tableIndex*i)
	}
}

// Prints all multiplication tables from 2 to 9
func displayMultiplicationTableAll() {
	for i := 2; i <= 9; i++ {
		displayMultiplicationTableAt(i)
		fmt.Print("\n")
	}
}

func main() {
	// Variable to store the user's input
	var tableStartIndex int

	// Program start message
	fmt.Println("=================")
	fmt.Println("Multiplication Table Program")
	fmt.Println("=================")

	// Ask the user which table they want to see
	fmt.Print("Which multiplication table do you want to print? ")
	fmt.Print("(0: Exit program): ")

	// Read user input
	fmt.Scanf("%d\n", &tableStartIndex)

	// If the user enters 0 → exit program
	if tableStartIndex == 0 {
		fmt.Println("Exiting the program.")
		return
	}

	// If the user enters 2–9 → print that specific table
	if tableStartIndex >= 2 && tableStartIndex <= 9 {
		displayMultiplicationTableAt(tableStartIndex)
	} else {
		// Otherwise → print all tables
		displayMultiplicationTableAll()
	}
}
