package main

import "fmt"

// Rectangle struct
type Rectangle struct {
	x, y int
}

// Constructor function for Rectangle
func NewRectangle() *Rectangle {
	return new(Rectangle)
}

// Draw method for Rectangle
func (r Rectangle) Draw() {
	fmt.Println("Drawing a rectangle!")
}

func main() {
	rectangle := NewRectangle()
	rectangle.Draw()
}
