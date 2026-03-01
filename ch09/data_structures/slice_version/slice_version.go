package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nameSlice := make([]string, 5)

	for i := 0; i < 5; i++ {
		// "Enter the name of person %d: "
		fmt.Printf("Enter the name of person %d: ", i+1)
		fmt.Scanf("%s\n", &nameSlice[i])
	}

	randomNumber := rand.Intn(len(nameSlice))

	fmt.Print("\n")
	// "This time, %s will pay for lunch!"
	fmt.Printf("This time, %s will pay for lunch!\n", nameSlice[randomNumber])
}
