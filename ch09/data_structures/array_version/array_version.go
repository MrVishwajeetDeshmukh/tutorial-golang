package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var nameArray [5]string

	for i := 0; i < 5; i++ {
		fmt.Printf("Please enter the name of person %d: ", i+1)
		fmt.Scanf("%s\n", &nameArray[i])
	}

	randomNumber := rand.Intn(len(nameArray))

	fmt.Print("\n")
	fmt.Printf("%s will pay for today’s lunch!\n", nameArray[randomNumber])
}
