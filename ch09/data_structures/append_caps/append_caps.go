package main

import "fmt"

func main() {
	// Declare a nil slice
	var sliceValue []int

	fmt.Println("================================================")
	fmt.Printf("Before adding elements, sliceValue's len: %d, cap: %d\n",
		len(sliceValue), cap(sliceValue))
	fmt.Println("================================================")

	var itemNumber int
	fmt.Print("How many elements should be added to the array?: ")
	fmt.Scanf("%d\n", &itemNumber)

	for i := 0; i < itemNumber; i++ {
		// Add elements to the slice using the append function
		sliceValue = append(sliceValue, i)

		fmt.Printf("Added element %d, len: %d, cap: %d\n",
			i, len(sliceValue), cap(sliceValue))
	}

	fmt.Println("================================================")
	fmt.Printf("After adding elements, sliceValue's len: %d, cap: %d\n",
		len(sliceValue), cap(sliceValue))
	fmt.Println("================================================")
}
