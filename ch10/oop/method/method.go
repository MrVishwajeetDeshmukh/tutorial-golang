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

const dividerBar = "================================================"

type Product struct {
	Item     *Coffee
	Quantity int
}

type Store struct {
	Money    int
	Products []*Product
}

// Constructor for Product object
func NewProduct(item *Coffee, quantity int) *Product {
	product := new(Product)
	product.Item = item
	product.Quantity = quantity
	return product
}

// Constructor for Store object
func NewStore(money int, products []*Product) *Store {
	store := new(Store)
	store.Money = money
	store.Products = products
	return store
}

// Sell a certain quantity of a specific product
// Sell 'quantity' of product with name 'productName'
// and add the revenue to the Store's Money field
func (s *Store) SellProduct(productName string, quantity int) {
	product := s.GetProduct(productName)
	product.Quantity -= quantity
	s.Money += product.Item.Price * quantity
}

// Get all product objects
func (s *Store) GetProducts() []*Product {
	return s.Products
}

// Get a specific product by name
func (s *Store) GetProduct(productName string) *Product {
	for _, product := range s.Products {
		if product.Item.Name == productName {
			return product
		}
	}
	return nil
}

// Check if a specific product has the desired quantity
// Returns true if enough quantity, false if insufficient
func (s *Store) CheckProductQuantity(
	productName string, quantity int) bool {
	product := s.GetProduct(productName)
	return product.Quantity >= quantity
}

// Function to get the user's choice of product and quantity
func HandleChoiceProduct(myStore *Store) (exit bool) {
	for {
		var choice string

		fmt.Println(dividerBar)
		fmt.Print("Please enter the name of the product to purchase (exit to quit): ")
		fmt.Scanln(&choice)

		if choice == "exit" {
			fmt.Println("Thank you for visiting!")
			fmt.Println(dividerBar)
			fmt.Printf("Final store balance: %d\n", myStore.Money)
			return true
		}

		product := myStore.GetProduct(choice)
		if product == nil {
			fmt.Printf("We don't have a product named %s in our store.\n", choice)
			continue
		} else if product.Quantity == 0 {
			fmt.Printf("%s is out of stock.\n", choice)
			continue
		}

		var quantity int
		fmt.Print("How many would you like to buy?: ")
		fmt.Scanln(&quantity)

		isExists := myStore.CheckProductQuantity(
			product.Item.Name,
			quantity)
		if !isExists {
			fmt.Printf("%s does not have enough stock.\n", choice)
			continue
		}

		myStore.SellProduct(product.Item.Name, quantity)
		fmt.Println("Thank you for shopping with us!")
		break
	}

	return false
}

func main() {
	// Define coffee types
	americano := NewCoffee("Americano", 3000, "Blended Coffee", Bitter)
	latte := NewCoffee("Café Latte", 3500, "Blended Coffee", Sweet)
	caffeMocha := NewCoffee("Caffè Mocha", 4000, "Dessert Coffee", Sweet)
	dripCoffee := NewCoffee("Drip Coffee", 7000, "Bean Coffee", FruitFlavor)
	dutchCoffee := NewCoffee("Dutch Coffee", 5000, "Dutch Coffee", Bitter)

	// Create Product objects with stock quantity
	productAmericano := NewProduct(americano, 5)
	productLatte := NewProduct(latte, 2)
	productCaffeMocha := NewProduct(caffeMocha, 3)
	productDripCoffee := NewProduct(dripCoffee, 4)
	productDutchCoffee := NewProduct(dutchCoffee, 6)

	// Create a Store object to sell products
	// Initial money is 0
	// Products for sale are provided as an array of Product objects
	myStore := NewStore(0, []*Product{
		productAmericano,
		productLatte,
		productCaffeMocha,
		productDripCoffee,
		productDutchCoffee})
	fmt.Printf("store: %v\n", myStore)

	for {
		fmt.Println(dividerBar)
		fmt.Println("Welcome to our store!")
		fmt.Printf("Current store balance: %d\n", myStore.Money)
		fmt.Println(dividerBar)
		fmt.Println("Here is the list of products available:")
		for i, product := range myStore.GetProducts() {
			fmt.Printf("%d. %s: %d, (Stock: %d)\n",
				i+1,
				product.Item.Name,
				product.Item.Price,
				product.Quantity)
		}

		isExit := HandleChoiceProduct(myStore)
		if isExit {
			return
		}
	}
}
