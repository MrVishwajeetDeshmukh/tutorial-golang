package main

// Interface for Coffee, Juice, and Tea objects
// Defines the methods Make(), Package(), and Pick()
type ProductItem interface {
	Make() error
	Package() error
	Pick() error

	// Getter methods to access member fields
	Name() string
	Price() int
	Category() string
	Taste() Taste
	State() State
}
