package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var nameSlice []string
	var name string

	for i := 0; name != "exit"; i++ {
		fmt.Printf("Enter the name of person %d (type 'exit' to stop): ", i+1)
		fmt.Scanf("%s\n", &name)
		if name == "exit" {
			break
		}
		nameSlice = append(nameSlice, name)
	}

	randomNumber := rand.Intn(len(nameSlice))

	fmt.Print("\n")
	fmt.Printf("%s will pay for today's lunch!\n", nameSlice[randomNumber])
}
