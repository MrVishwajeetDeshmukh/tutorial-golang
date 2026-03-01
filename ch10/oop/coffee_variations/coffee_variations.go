package main

import "fmt"

type Taste string

const (
	Sweet       Taste = "SWEET"
	Bitter      Taste = "BITTER"
	FruitFlavor Taste = "FRUIT"
	Heavy       Taste = "HEAVY"
)

type Coffee struct {
	Name     string
	Price    int
	Category string
	Taste    Taste
}

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
	americano := NewCoffee("Americano", 3000, "Blended Coffee", Bitter)
	latte := NewCoffee("Cafe Latte", 3500, "Blended Coffee", Sweet)
	caffeMocha := NewCoffee("Cafe Mocha", 4000, "Dessert Coffee", Sweet)
	dripCoffee := NewCoffee("Drip Coffee", 7000, "Single Origin Coffee", FruitFlavor)
	dutchCoffee := NewCoffee("Dutch Coffee", 5000, "Dutch Coffee", Bitter)

	fmt.Printf("americano: %v\n", americano)
	fmt.Printf("latte: %v\n", latte)
	fmt.Printf("caffeMocha: %v\n", caffeMocha)
	fmt.Printf("dripCoffee: %v\n", dripCoffee)
	fmt.Printf("dutchCoffee: %v\n", dutchCoffee)
}
