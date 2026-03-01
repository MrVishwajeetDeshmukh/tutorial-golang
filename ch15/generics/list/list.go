package main

import "fmt"

// Define a generic list type
type List[T any] struct {
	items []T
}

// Add an item to the list
func (l *List[T]) Add(item T) {
	l.items = append(l.items, item)
}

// Print list items
func (l *List[T]) Print() {
	for _, item := range l.items {
		fmt.Println(item)
	}
}

func main() {
	// Integer list
	intList := List[int]{}
	intList.Add(1)
	intList.Add(2)
	intList.Print() // Output: 1 \n 2

	// String list
	strList := List[string]{}
	strList.Add("hello")
	strList.Add("world")
	strList.Print() // Output: hello \n world
}
