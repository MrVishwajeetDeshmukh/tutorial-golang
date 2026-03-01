package main

import "fmt"

// Shape interface defines a Draw method
type Shape interface {
	Draw()
}

// Circle struct
type Circle struct{}

// Draw method for Circle
func (c Circle) Draw() {
	fmt.Println("Drawing a Circle")
}

// Square struct
type Square struct{}

// Draw method for Square
func (s Square) Draw() {
	fmt.Println("Drawing a Square")
}

func main() {
	var shape Shape

	shape = Circle{}
	shape.Draw() // Output: Drawing a Circle

	shape = Square{}
	shape.Draw() // Output: Drawing a Square
}
