package main

import "fmt"

type Rectangle struct {
	x, y int
}

func NewRectangle(x int, y int) *Rectangle {
	rectangle := new(Rectangle)
	rectangle.x = x
	rectangle.y = y
	return rectangle
}

// This part has been changed
// Using *Rectangle to receive a pointer to the struct
func (r *Rectangle) Draw(x int, y int) {
	fmt.Printf("Before change x, y: %d, %d\n", r.x, r.y)
	r.x = x
	r.y = y
	fmt.Printf("After change x, y: %d, %d\n", r.x, r.y)
}

func main() {
	rectangle := NewRectangle(5, 5)
	rectangle.Draw(10, 20)
	fmt.Printf("Original Rectangle x, y: %d, %d\n",
		rectangle.x, rectangle.y)
}
