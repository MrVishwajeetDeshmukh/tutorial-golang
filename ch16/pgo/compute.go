package main

import (
	"fmt"
	"sort"
	"time"
)

// Heavy task requiring computation
func Compute(data []int) int {
	result := 0
	for _, value := range data {
		if value%2 == 0 {
			result += value * 2
		} else {
			result -= value * 2
		}
	}
	return result
}

func main() {
	data := make([]int, 1000000000)
	for i := range data {
		data[i] = i % 100
	}

	sort.Ints(data)

	start := time.Now()
	result := Compute(data)
	elapsed := time.Since(start)

	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Execution time: %s\n", elapsed)
}
