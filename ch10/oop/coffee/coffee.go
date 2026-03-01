package main

import "fmt"

// Taste type to represent flavor of coffee
type Taste string

const (
	Sweet       Taste = "SWEET"
	Bitter      Taste = "BITTER"
	FruitFlavor Taste = "FRUIT"
	Heavy       Taste = "HEAVY"
)

// Coffee struct represents a coffee object
type Coffee struct {
	Name     string
	Price    int
	Category string
	Taste    Taste
}

// NewCoffee creates a new Coffee instance and returns a pointer
func NewCoffee(
	name string,
	price int,
	category string,
	taste Taste) *Coffee {
	coffee := new(Coffee)
	coffee.Name = name
	coffee.Price = price
	coffee.Category = category
	coffee.Taste = taste
	return coffee
}

func main() {
	// Create a new coffee instance
	coffee := NewCoffee("Americano", 3000, "Blended Coffee", Bitter)

	fmt.Println("================================================")
	fmt.Printf("coffee object: %v\n", coffee)
	fmt.Println("================================================")
	fmt.Printf("name: %s\n", coffee.Name)
	fmt.Printf("price: %d\n", coffee.Price)
	fmt.Printf("category: %s\n", coffee.Category)
	fmt.Printf("taste: %s\n", coffee.Taste)
	fmt.Println("================================================")
}
