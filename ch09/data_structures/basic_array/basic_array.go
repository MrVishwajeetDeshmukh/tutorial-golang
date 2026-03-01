package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var name1, name2, name3, name4, name5 string

	for i := 0; i < 5; i++ {
		var name string
		fmt.Printf("Please enter the name of person %d: ", i+1)
		fmt.Scanf("%s\n", &name)
		switch i {
		case 0:
			name1 = name
		case 1:
			name2 = name
		case 2:
			name3 = name
		case 3:
			name4 = name
		case 4:
			name5 = name
		}
	}

	randomNumber := rand.Intn(5)

	var targetPeople string
	switch randomNumber {
	case 0:
		targetPeople = name1
	case 1:
		targetPeople = name2
	case 2:
		targetPeople = name3
	case 3:
		targetPeople = name4
	case 4:
		targetPeople = name5
	}

	fmt.Print("\n")
	fmt.Printf("%s will pay for today’s lunch!\n", targetPeople)
}
