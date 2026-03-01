package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
			if j == 1 {
				// Exit the loop on line 7
				break

				// Exit the loop on line 6
				break
			}
		}
	}
}
