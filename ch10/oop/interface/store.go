package main

import "fmt"

const dividerBar = "================================================"

type Product struct {
	Item     ProductItem
	Quantity int
}

type Store struct {
	Money    int
	Products []*Product
}

// Constructor for Product object
func NewProduct(item ProductItem, quantity int) *Product {
	product := new(Product)
	product.Item = item
	product.Quantity = quantity
	return product
}

// Serve the Product to a customer
func (p *Product) Serve() error {
	// Make the product
	err := p.Item.Make()

	// If there is an error during the making process, propagate it to the parent function
	if err != nil {
		return err
	}

	// Package the product
	err = p.Item.Package()

	// If there is an error during the packaging process, propagate it to the parent function
	if err != nil {
		return err
	}

	// Pick and hand over the product
	err = p.Item.Pick()

	// If there is an error during picking, propagate it to the parent function
	if err != nil {
		return err
	}

	// If all processes succeed, return nil for error
	return nil
}

// Constructor for Store object
func NewStore(money int, products []*Product) *Store {
	store := new(Store)
	store.Money = money
	store.Products = products
	return store
}

// Sell a specific product in the desired quantity
// Sells the product with productName in quantity and
// adds the revenue to the Store's Money member variable
func (s *Store) SellProduct(productName string, quantity int) {
	product := s.GetProduct(productName)

	// Serve the product to the customer
	err := product.Serve()
	if err != nil {
		fmt.Printf("An error occurred while selling the product: %v\n",
			err)
		return
	}

	product.Quantity -= quantity
	s.Money += product.Item.Price() * quantity
}

// Return the list of all product objects
func (s *Store) GetProducts() []*Product {
	return s.Products
}

// Get a product object by its name
func (s *Store) GetProduct(productName string) *Product {
	for _, product := range s.Products {
		if product.Item.Name() == productName {
			return product
		}
	}
	return nil
}

// Check if a product has enough quantity available
// Enough quantity: true, Not enough: false
func (s *Store) CheckProductQuantity(
	productName string, quantity int) bool {
	product := s.GetProduct(productName)
	return product.Quantity >= quantity
}

// Get the product name and quantity from the user
func HandleChoiceProduct(myStore *Store) (exit bool) {
	for {
		var choice string

		fmt.Println(dividerBar)
		fmt.Print("Please enter the name of the product you want to buy (exit to leave): ")
		fmt.Scanln(&choice)

		if choice == "exit" {
			fmt.Println("Thank you for visiting!")
			fmt.Println(dividerBar)
			fmt.Printf("Final store balance: %d\n", myStore.Money)
			return true
		}

		product := myStore.GetProduct(choice)
		if product == nil {
			fmt.Printf("We don't have a product named %s in our store.\n",
				choice)
			continue
		} else if product.Quantity == 0 {
			fmt.Printf("%s is out of stock...\n", choice)
			continue
		}

		var quantity int
		fmt.Print("How many would you like to buy?: ")
		fmt.Scanln(&quantity)

		isExists := myStore.CheckProductQuantity(
			product.Item.Name(),
			quantity)
		if isExists == false {
			fmt.Printf("%s is out of stock...\n", choice)
			continue
		}

		myStore.SellProduct(product.Item.Name(), quantity)
		fmt.Println("Thank you for using our store!")
		break
	}

	return false
}

func main() {
	// Define coffee types
	americano := NewCoffee(
		"Americano", 3000, "Blending Coffee", Bitter, Waiting)
	latte := NewCoffee(
		"Cafe Latte", 3500, "Blending Coffee", Sweet, Waiting)
	caffeMocha := NewCoffee(
		"Cafe Mocha", 4000, "Dessert Coffee", Sweet, Waiting)
	dripCoffee := NewCoffee(
		"Drip Coffee", 7000, "Single Origin Coffee", FruitFlavor, Waiting)
	dutchCoffee := NewCoffee(
		"Dutch Coffee", 5000, "Dutch Coffee", Bitter, Waiting)

	// Define juice types
	orangeJuice := NewJuice(
		"Orange Juice", 3000, "Fruit Juice", Sweet, Waiting)
	aloeJuice := NewJuice(
		"Aloe Juice", 3000, "Health Juice", FruitFlavor, Waiting)

	// Define tea types
	milk := NewTea(
		"Milk", 3000, "Milk", Sweet, Waiting)
	hotChoco := NewTea(
		"Hot Chocolate", 3000, "Hot Chocolate", Sweet, Waiting)

	// Create Product objects with available quantities
	productAmericano := NewProduct(americano, 5)
	productLatte := NewProduct(latte, 2)
	productCaffeMocha := NewProduct(caffeMocha, 3)
	productDripCoffee := NewProduct(dripCoffee, 4)
	productDutchCoffee := NewProduct(dutchCoffee, 6)
	productOrangeJuice := NewProduct(orangeJuice, 8)
	productAloeJuice := NewProduct(aloeJuice, 3)
	productMilk := NewProduct(milk, 3)
	productHotChoco := NewProduct(hotChoco, 5)

	// Create the store object
	// Initial money: 0
	// Products to sell: the array of Product objects defined above
	myStore := NewStore(0, []*Product{
		productAmericano,
		productLatte,
		productCaffeMocha,
		productDripCoffee,
		productDutchCoffee,
		productOrangeJuice,
		productAloeJuice,
		productMilk,
		productHotChoco,
	})
	fmt.Printf("store: %v", myStore)

	for {
		fmt.Println(dividerBar)
		fmt.Println("Welcome to our store!")
		fmt.Printf("Current store balance: %d\n", myStore.Money)
		fmt.Println(dividerBar)
		fmt.Println("Here is the list of products in our store:")
		for i, product := range myStore.GetProducts() {
			fmt.Printf("%d. %s: %d, (Stock: %d)\n",
				i+1,
				product.Item.Name(),
				product.Item.Price(),
				product.Quantity)
		}

		isExit := HandleChoiceProduct(myStore)
		if isExit {
			return
		}
	}
}
